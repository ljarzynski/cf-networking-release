package scaling_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib/models"
	"lib/policy_client"
	"lib/testsupport"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"code.cloudfoundry.org/lager/lagertest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf-experimental/rainmaker"
)

const Timeout_Check = 20 * time.Minute

// 3 * observed agent total poll time with 200 containers per cell
const Policy_Update_Wait = 180 * time.Second

var ports []int

var _ = Describe("connectivity between containers on the overlay network", func() {
	Describe("networking policy", func() {
		var (
			appProxy       string
			appRegistry    string
			appsTest       []string
			appInstances   int
			applications   int
			proxyInstances int
			sampleSize     int
		)

		BeforeEach(func() {
			appInstances = pushConfig.AppInstances
			applications = pushConfig.Applications
			proxyInstances = pushConfig.ProxyInstances
			sampleSize = pushConfig.SampleSize
			appProxy = pushConfig.Prefix + "proxy"
			appRegistry = pushConfig.Prefix + "registry"
			for i := 0; i < applications; i++ {
				appsTest = append(appsTest, fmt.Sprintf(pushConfig.Prefix+"tick-%d", i))
			}

			ports = []int{8080}
			for i := 0; i < pushConfig.ExtraListenPorts; i++ {
				ports = append(ports, 7000+i)
			}
		})

		It("allows the user to configure policies", func(done Done) {
			By("checking that all test app instances have registered themselves")
			checkRegistry(appRegistry, 10*time.Second, 500*time.Millisecond, len(appsTest)*appInstances)

			appIPs := getAppIPs(appRegistry)
			sample := sampleIPs(appIPs, sampleSize)

			By(fmt.Sprintf("checking that the connection fails sampling %d out of %d IPs", len(sample), len(appIPs)))
			runWithTimeout("check connection failures", Timeout_Check, func() {
				assertConnectionFails(appProxy, sample, ports, proxyInstances)
			})

			By("creating policies")
			doAllPolicies("create", appProxy, appsTest, ports)

			By(fmt.Sprintf("waiting %s for policies to be updated on cells", Policy_Update_Wait))
			time.Sleep(Policy_Update_Wait)

			sample = sampleIPs(appIPs, sampleSize)
			By(fmt.Sprintf("checking that the connection succeeds sampling %d out of %d IPs", len(sample), len(appIPs)))
			runWithTimeout("check connection success", Timeout_Check, func() {
				assertConnectionSucceeds(appProxy, sample, ports, proxyInstances)
			})

			dumpStats(appProxy, config.AppsDomain)

			By("deleting policies")
			doAllPolicies("delete", appProxy, appsTest, ports)

			By(fmt.Sprintf("waiting %s for policies to be updated on cells", Policy_Update_Wait))
			time.Sleep(Policy_Update_Wait)

			sample = sampleIPs(appIPs, sampleSize)
			By(fmt.Sprintf("checking that the connection succeeds sampling %d out of %d IPs", len(sample), len(appIPs)))
			runWithTimeout("check connection failures, again", Timeout_Check, func() {
				assertConnectionFails(appProxy, sample, ports, proxyInstances)
			})

			close(done)
		}, 30*60 /* <-- overall spec timeout in seconds */)
	})
})

func getToken() string {
	By("getting token")
	cmd := exec.Command("cf", "oauth-token")
	session, err := gexec.Start(cmd, nil, nil)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session.Wait(Timeout_Short)).Should(gexec.Exit(0))
	rawOutput := string(session.Out.Contents())
	return strings.TrimSpace(strings.TrimPrefix(rawOutput, "bearer "))
}

func getGuids(sourceAppName string, dstAppNames []string) (string, []string) {
	dstGuids := []string{}
	sourceGuid := ""
	token := getToken()
	appsClient := rainmaker.NewApplicationsService(rainmaker.Config{Host: "http://" + config.ApiEndpoint})

	appsList, err := appsClient.List(token)
	Expect(err).NotTo(HaveOccurred())

	for {
		for _, app := range appsList.Applications {
			if app.Name == sourceAppName {
				sourceGuid = app.GUID
				continue
			}
			for _, tickAppName := range dstAppNames {
				if app.Name == tickAppName {
					dstGuids = append(dstGuids, app.GUID)
					break
				}
			}
		}
		if appsList.HasNextPage() {
			appsList, err = appsList.Next(token)
			Expect(err).NotTo(HaveOccurred())
		} else {
			break
		}
	}

	Expect(sourceGuid).NotTo(BeEmpty())
	Expect(dstGuids).To(HaveLen(len(dstAppNames)))

	return sourceGuid, dstGuids
}

func doAllPolicies(action string, source string, dstList []string, dstPorts []int) {
	policyClient := policy_client.NewExternal(lagertest.NewTestLogger("test"), &http.Client{}, "http://"+config.ApiEndpoint)
	sourceGuid, dstGuids := getGuids(source, dstList)
	policies := []models.Policy{}
	for _, dstGuid := range dstGuids {
		for _, port := range dstPorts {
			policies = append(policies, models.Policy{
				Source: models.Source{
					ID: sourceGuid,
				},
				Destination: models.Destination{
					ID:       dstGuid,
					Port:     port,
					Protocol: "tcp",
				},
			})
		}
	}
	token := getToken()
	if action == "create" {
		Expect(policyClient.AddPolicies(token, policies)).To(Succeed())
	} else if action == "delete" {
		Expect(policyClient.DeletePolicies(token, policies)).To(Succeed())
	}
}

func runWithTimeout(operation string, timeout time.Duration, work func()) {
	done := make(chan bool)
	go func() {
		fmt.Printf("starting %s\n", operation)
		work()
		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("completed %s\n", operation)
		return
	case <-time.After(timeout):
		Fail("timeout on " + operation)
	}
}

func dumpStats(host, domain string) {
	resp, err := httpGetBytes(fmt.Sprintf("http://%s.%s/stats", host, domain))
	Expect(err).NotTo(HaveOccurred())

	fmt.Printf("STATS: %s\n", string(resp.Body))
	netStatsFile := os.Getenv("NETWORK_STATS_FILE")
	if netStatsFile != "" {
		Expect(ioutil.WriteFile(netStatsFile, resp.Body, 0600)).To(Succeed())
	}
}

type RegistryInstancesResponse struct {
	Instances []struct {
		ServiceName string `json:"service_name"`
		Endpoint    struct {
			Value string `json:"value"`
		} `json:"endpoint"`
	} `json:"instances"`
}

func getInstancesFromA8(registry string) (*RegistryInstancesResponse, error) {
	resp, err := httpGetBytes(fmt.Sprintf("http://%s.%s/api/v1/instances", registry, config.AppsDomain))
	if err != nil {
		return nil, err
	}

	Expect(resp.StatusCode).To(Equal(http.StatusOK))

	var instancesResponse RegistryInstancesResponse
	err = json.Unmarshal(resp.Body, &instancesResponse)
	if err != nil {
		return nil, err
	}
	return &instancesResponse, nil
}

func checkRegistry(registry string, timeout, pollingInterval time.Duration, totalInstances int) {
	registeredApps := func() (int, error) {
		instancesResponse, err := getInstancesFromA8(registry)
		if err != nil {
			return 0, err
		}
		return len(instancesResponse.Instances), nil
	}

	Eventually(registeredApps, timeout, pollingInterval).Should(Equal(totalInstances))
}

func getAppIPs(registry string) []string {
	instancesResponse, err := getInstancesFromA8(registry)
	Expect(err).NotTo(HaveOccurred())

	ips := []string{}
	for _, instance := range instancesResponse.Instances {
		ip, _, err := net.SplitHostPort(instance.Endpoint.Value)
		Expect(err).NotTo(HaveOccurred())
		ips = append(ips, ip)
	}
	return ips
}

func sampleIPs(population []string, sampleSize int) []string {
	populationSize := len(population)
	if len(population) < sampleSize || sampleSize == 0 {
		sampleSize = populationSize
	}
	var sample = []string{}
	for i := 0; i < sampleSize; i++ {
		j := rand.Intn(populationSize + 1)
		sample = append(sample, population[j])
	}
	return sample
}

func assertConnectionSucceeds(sourceApp string, destApps []string, ports []int, nProxies int) {
	parallelRunner := &testsupport.ParallelRunner{
		NumWorkers: 50 * nProxies,
	}
	parallelRunner.RunOnSliceStrings(destApps, func(appIP string) {
		for _, port := range ports {
			assertSingleConnection(appIP, port, sourceApp, true)
		}
	})
}

func assertConnectionFails(sourceApp string, destApps []string, ports []int, nProxies int) {
	parallelRunner := &testsupport.ParallelRunner{
		NumWorkers: 50 * nProxies,
	}
	parallelRunner.RunOnSliceStrings(destApps, func(appIP string) {
		for _, port := range ports {
			assertSingleConnection(appIP, port, sourceApp, false)
		}
	})
}

func assertSingleConnection(destIP string, port int, sourceAppName string, shouldSucceed bool) {
	if shouldSucceed {
		By(fmt.Sprintf("eventually proxy should reach %s at port %d", destIP, port))
		assertResponseContains(destIP, port, sourceAppName, "application_name")
	} else {
		By(fmt.Sprintf("eventually proxy should NOT reach %s at port %d", destIP, port))
		assertResponseContains(destIP, port, sourceAppName, "request failed")
	}
}

func assertResponseContains(destIP string, port int, sourceAppName string, desiredResponse string) {
	proxyTest := func() (string, error) {
		resp, err := httpGetBytes(fmt.Sprintf("http://%s.%s/proxy/%s:%d", sourceAppName, config.AppsDomain, destIP, port))
		if err != nil {
			return "", err
		}
		return string(resp.Body), nil
	}
	Eventually(proxyTest, 10*time.Second, 500*time.Millisecond).Should(ContainSubstring(desiredResponse))
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
		Dial: (&net.Dialer{
			Timeout:   4 * time.Second,
			KeepAlive: 0,
		}).Dial,
	},
}

type httpResp struct {
	StatusCode int
	Body       []byte
}

func httpGetBytes(url string) (httpResp, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return httpResp{}, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResp{}, err
	}

	return httpResp{resp.StatusCode, respBytes}, nil
}
