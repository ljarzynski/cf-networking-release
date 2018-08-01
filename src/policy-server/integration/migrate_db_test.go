package integration_test

import (
	"fmt"
	"math"
	"time"

	"policy-server/config"
	"policy-server/db"
	"policy-server/integration/helpers"
	"policy-server/store/migrations"
	"test-helpers"

	dbHelper "code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport/ports"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

const TimeoutShort = 20 * time.Second

var _ = Describe("Migrate DB Binary", func() {
	var (
		dbConf dbHelper.Config
		conf   config.Config
	)

	BeforeEach(func() {
		dbConf = testsupport.GetDBConfig()
		dbConf.DatabaseName = fmt.Sprintf("migrate_test_node_%d", ports.PickAPort())

		conf, _ = helpers.DefaultTestConfig(dbConf, "unused", "fixtures")
		conf.Database = dbConf
	})

	Context("when the db is available", func() {
		BeforeEach(func() {
			testhelpers.CreateDatabase(dbConf)
		})

		AfterEach(func() {
			testhelpers.RemoveDatabase(dbConf)
		})

		It("runs the migrations and seeds the groups table", func() {
			session := helpers.RunMigrationsPreStartBinary(migrateDbPath, conf)
			Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(0))
			conn := db.NewConnectionPool(
				dbConf,
				1,
				1,
				"test-db",
				"test-job-prefix",
				lagertest.NewTestLogger("test"),
			)

			defer conn.Close()

			numMigrations := len(migrations.V1ModifiedMigrationsToPerform) +
				len(migrations.V2ModifiedMigrationsToPerform) +
				len(migrations.V3ModifiedMigrationsToPerform) +
				len(migrations.MigrationsToPerform)

			var migrationCount int
			conn.QueryRow("SELECT COUNT(*) FROM gorp_migrations").Scan(&migrationCount)
			Expect(migrationCount).To(Equal(numMigrations))

			var groupCount int
			conn.QueryRow("SELECT COUNT(*) FROM groups").Scan(&groupCount)
			Expect(groupCount).To(Equal(int(math.Exp2(float64(conf.TagLength*8))) - 1))
		})

		Context("when the migrations have already run", func() {
			It("runs successfully", func() {
				session := helpers.RunMigrationsPreStartBinary(migrateDbPath, conf)
				Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(0))
				session = helpers.RunMigrationsPreStartBinary(migrateDbPath, conf)
				Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(0))
			})
		})
	})

	Context("when the db is not available", func() {
		It("exits non zero", func() {
			session := helpers.RunMigrationsPreStartBinary(migrateDbPath, conf)
			Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(1))
		})
	})
})
