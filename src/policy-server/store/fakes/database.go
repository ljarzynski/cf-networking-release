// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"database/sql"
	"policy-server/store"
	"sync"

	"code.cloudfoundry.org/cf-networking-helpers/db"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	BeginxStub        func() (db.Transaction, error)
	beginxMutex       sync.RWMutex
	beginxArgsForCall []struct{}
	beginxReturns     struct {
		result1 db.Transaction
		result2 error
	}
	beginxReturnsOnCall map[int]struct {
		result1 db.Transaction
		result2 error
	}
	ExecStub        func(query string, args ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		query string
		args  []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	NamedExecStub        func(query string, arg interface{}) (sql.Result, error)
	namedExecMutex       sync.RWMutex
	namedExecArgsForCall []struct {
		query string
		arg   interface{}
	}
	namedExecReturns struct {
		result1 sql.Result
		result2 error
	}
	namedExecReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	GetStub        func(dest interface{}, query string, args ...interface{}) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		dest  interface{}
		query string
		args  []interface{}
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	SelectStub        func(dest interface{}, query string, args ...interface{}) error
	selectMutex       sync.RWMutex
	selectArgsForCall []struct {
		dest  interface{}
		query string
		args  []interface{}
	}
	selectReturns struct {
		result1 error
	}
	selectReturnsOnCall map[int]struct {
		result1 error
	}
	QueryRowStub        func(query string, args ...interface{}) *sql.Row
	queryRowMutex       sync.RWMutex
	queryRowArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryRowReturns struct {
		result1 *sql.Row
	}
	queryRowReturnsOnCall map[int]struct {
		result1 *sql.Row
	}
	QueryStub        func(query string, args ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 *sql.Rows
		result2 error
	}
	DriverNameStub        func() string
	driverNameMutex       sync.RWMutex
	driverNameArgsForCall []struct{}
	driverNameReturns     struct {
		result1 string
	}
	driverNameReturnsOnCall map[int]struct {
		result1 string
	}
	RawConnectionStub        func() *sqlx.DB
	rawConnectionMutex       sync.RWMutex
	rawConnectionArgsForCall []struct{}
	rawConnectionReturns     struct {
		result1 *sqlx.DB
	}
	rawConnectionReturnsOnCall map[int]struct {
		result1 *sqlx.DB
	}
	RebindStub        func(string) string
	rebindMutex       sync.RWMutex
	rebindArgsForCall []struct {
		arg1 string
	}
	rebindReturns struct {
		result1 string
	}
	rebindReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Db) Beginx() (db.Transaction, error) {
	fake.beginxMutex.Lock()
	ret, specificReturn := fake.beginxReturnsOnCall[len(fake.beginxArgsForCall)]
	fake.beginxArgsForCall = append(fake.beginxArgsForCall, struct{}{})
	fake.recordInvocation("Beginx", []interface{}{})
	fake.beginxMutex.Unlock()
	if fake.BeginxStub != nil {
		return fake.BeginxStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.beginxReturns.result1, fake.beginxReturns.result2
}

func (fake *Db) BeginxCallCount() int {
	fake.beginxMutex.RLock()
	defer fake.beginxMutex.RUnlock()
	return len(fake.beginxArgsForCall)
}

func (fake *Db) BeginxReturns(result1 db.Transaction, result2 error) {
	fake.BeginxStub = nil
	fake.beginxReturns = struct {
		result1 db.Transaction
		result2 error
	}{result1, result2}
}

func (fake *Db) BeginxReturnsOnCall(i int, result1 db.Transaction, result2 error) {
	fake.BeginxStub = nil
	if fake.beginxReturnsOnCall == nil {
		fake.beginxReturnsOnCall = make(map[int]struct {
			result1 db.Transaction
			result2 error
		})
	}
	fake.beginxReturnsOnCall[i] = struct {
		result1 db.Transaction
		result2 error
	}{result1, result2}
}

func (fake *Db) Exec(query string, args ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Exec", []interface{}{query, args})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(query, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.execReturns.result1, fake.execReturns.result2
}

func (fake *Db) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *Db) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].query, fake.execArgsForCall[i].args
}

func (fake *Db) ExecReturns(result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *Db) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *Db) NamedExec(query string, arg interface{}) (sql.Result, error) {
	fake.namedExecMutex.Lock()
	ret, specificReturn := fake.namedExecReturnsOnCall[len(fake.namedExecArgsForCall)]
	fake.namedExecArgsForCall = append(fake.namedExecArgsForCall, struct {
		query string
		arg   interface{}
	}{query, arg})
	fake.recordInvocation("NamedExec", []interface{}{query, arg})
	fake.namedExecMutex.Unlock()
	if fake.NamedExecStub != nil {
		return fake.NamedExecStub(query, arg)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.namedExecReturns.result1, fake.namedExecReturns.result2
}

func (fake *Db) NamedExecCallCount() int {
	fake.namedExecMutex.RLock()
	defer fake.namedExecMutex.RUnlock()
	return len(fake.namedExecArgsForCall)
}

func (fake *Db) NamedExecArgsForCall(i int) (string, interface{}) {
	fake.namedExecMutex.RLock()
	defer fake.namedExecMutex.RUnlock()
	return fake.namedExecArgsForCall[i].query, fake.namedExecArgsForCall[i].arg
}

func (fake *Db) NamedExecReturns(result1 sql.Result, result2 error) {
	fake.NamedExecStub = nil
	fake.namedExecReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *Db) NamedExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.NamedExecStub = nil
	if fake.namedExecReturnsOnCall == nil {
		fake.namedExecReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.namedExecReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *Db) Get(dest interface{}, query string, args ...interface{}) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		dest  interface{}
		query string
		args  []interface{}
	}{dest, query, args})
	fake.recordInvocation("Get", []interface{}{dest, query, args})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(dest, query, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *Db) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *Db) GetArgsForCall(i int) (interface{}, string, []interface{}) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].dest, fake.getArgsForCall[i].query, fake.getArgsForCall[i].args
}

func (fake *Db) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *Db) GetReturnsOnCall(i int, result1 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Db) Select(dest interface{}, query string, args ...interface{}) error {
	fake.selectMutex.Lock()
	ret, specificReturn := fake.selectReturnsOnCall[len(fake.selectArgsForCall)]
	fake.selectArgsForCall = append(fake.selectArgsForCall, struct {
		dest  interface{}
		query string
		args  []interface{}
	}{dest, query, args})
	fake.recordInvocation("Select", []interface{}{dest, query, args})
	fake.selectMutex.Unlock()
	if fake.SelectStub != nil {
		return fake.SelectStub(dest, query, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.selectReturns.result1
}

func (fake *Db) SelectCallCount() int {
	fake.selectMutex.RLock()
	defer fake.selectMutex.RUnlock()
	return len(fake.selectArgsForCall)
}

func (fake *Db) SelectArgsForCall(i int) (interface{}, string, []interface{}) {
	fake.selectMutex.RLock()
	defer fake.selectMutex.RUnlock()
	return fake.selectArgsForCall[i].dest, fake.selectArgsForCall[i].query, fake.selectArgsForCall[i].args
}

func (fake *Db) SelectReturns(result1 error) {
	fake.SelectStub = nil
	fake.selectReturns = struct {
		result1 error
	}{result1}
}

func (fake *Db) SelectReturnsOnCall(i int, result1 error) {
	fake.SelectStub = nil
	if fake.selectReturnsOnCall == nil {
		fake.selectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.selectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Db) QueryRow(query string, args ...interface{}) *sql.Row {
	fake.queryRowMutex.Lock()
	ret, specificReturn := fake.queryRowReturnsOnCall[len(fake.queryRowArgsForCall)]
	fake.queryRowArgsForCall = append(fake.queryRowArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("QueryRow", []interface{}{query, args})
	fake.queryRowMutex.Unlock()
	if fake.QueryRowStub != nil {
		return fake.QueryRowStub(query, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.queryRowReturns.result1
}

func (fake *Db) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *Db) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return fake.queryRowArgsForCall[i].query, fake.queryRowArgsForCall[i].args
}

func (fake *Db) QueryRowReturns(result1 *sql.Row) {
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *Db) QueryRowReturnsOnCall(i int, result1 *sql.Row) {
	fake.QueryRowStub = nil
	if fake.queryRowReturnsOnCall == nil {
		fake.queryRowReturnsOnCall = make(map[int]struct {
			result1 *sql.Row
		})
	}
	fake.queryRowReturnsOnCall[i] = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *Db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Query", []interface{}{query, args})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(query, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryReturns.result1, fake.queryReturns.result2
}

func (fake *Db) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *Db) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].query, fake.queryArgsForCall[i].args
}

func (fake *Db) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *Db) QueryReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 *sql.Rows
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *Db) DriverName() string {
	fake.driverNameMutex.Lock()
	ret, specificReturn := fake.driverNameReturnsOnCall[len(fake.driverNameArgsForCall)]
	fake.driverNameArgsForCall = append(fake.driverNameArgsForCall, struct{}{})
	fake.recordInvocation("DriverName", []interface{}{})
	fake.driverNameMutex.Unlock()
	if fake.DriverNameStub != nil {
		return fake.DriverNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.driverNameReturns.result1
}

func (fake *Db) DriverNameCallCount() int {
	fake.driverNameMutex.RLock()
	defer fake.driverNameMutex.RUnlock()
	return len(fake.driverNameArgsForCall)
}

func (fake *Db) DriverNameReturns(result1 string) {
	fake.DriverNameStub = nil
	fake.driverNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *Db) DriverNameReturnsOnCall(i int, result1 string) {
	fake.DriverNameStub = nil
	if fake.driverNameReturnsOnCall == nil {
		fake.driverNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.driverNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Db) RawConnection() *sqlx.DB {
	fake.rawConnectionMutex.Lock()
	ret, specificReturn := fake.rawConnectionReturnsOnCall[len(fake.rawConnectionArgsForCall)]
	fake.rawConnectionArgsForCall = append(fake.rawConnectionArgsForCall, struct{}{})
	fake.recordInvocation("RawConnection", []interface{}{})
	fake.rawConnectionMutex.Unlock()
	if fake.RawConnectionStub != nil {
		return fake.RawConnectionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rawConnectionReturns.result1
}

func (fake *Db) RawConnectionCallCount() int {
	fake.rawConnectionMutex.RLock()
	defer fake.rawConnectionMutex.RUnlock()
	return len(fake.rawConnectionArgsForCall)
}

func (fake *Db) RawConnectionReturns(result1 *sqlx.DB) {
	fake.RawConnectionStub = nil
	fake.rawConnectionReturns = struct {
		result1 *sqlx.DB
	}{result1}
}

func (fake *Db) RawConnectionReturnsOnCall(i int, result1 *sqlx.DB) {
	fake.RawConnectionStub = nil
	if fake.rawConnectionReturnsOnCall == nil {
		fake.rawConnectionReturnsOnCall = make(map[int]struct {
			result1 *sqlx.DB
		})
	}
	fake.rawConnectionReturnsOnCall[i] = struct {
		result1 *sqlx.DB
	}{result1}
}

func (fake *Db) Rebind(arg1 string) string {
	fake.rebindMutex.Lock()
	ret, specificReturn := fake.rebindReturnsOnCall[len(fake.rebindArgsForCall)]
	fake.rebindArgsForCall = append(fake.rebindArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Rebind", []interface{}{arg1})
	fake.rebindMutex.Unlock()
	if fake.RebindStub != nil {
		return fake.RebindStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rebindReturns.result1
}

func (fake *Db) RebindCallCount() int {
	fake.rebindMutex.RLock()
	defer fake.rebindMutex.RUnlock()
	return len(fake.rebindArgsForCall)
}

func (fake *Db) RebindArgsForCall(i int) string {
	fake.rebindMutex.RLock()
	defer fake.rebindMutex.RUnlock()
	return fake.rebindArgsForCall[i].arg1
}

func (fake *Db) RebindReturns(result1 string) {
	fake.RebindStub = nil
	fake.rebindReturns = struct {
		result1 string
	}{result1}
}

func (fake *Db) RebindReturnsOnCall(i int, result1 string) {
	fake.RebindStub = nil
	if fake.rebindReturnsOnCall == nil {
		fake.rebindReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.rebindReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Db) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginxMutex.RLock()
	defer fake.beginxMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.namedExecMutex.RLock()
	defer fake.namedExecMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.selectMutex.RLock()
	defer fake.selectMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.driverNameMutex.RLock()
	defer fake.driverNameMutex.RUnlock()
	fake.rawConnectionMutex.RLock()
	defer fake.rawConnectionMutex.RUnlock()
	fake.rebindMutex.RLock()
	defer fake.rebindMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Db) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ store.Database = new(Db)
