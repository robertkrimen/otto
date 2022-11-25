package otto_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/require"
)

// testResult is a test driver.Result.
type testResult struct{}

func (r *testResult) LastInsertId() (int64, error) {
	return 0, fmt.Errorf("not supported")
}

func (r *testResult) RowsAffected() (int64, error) {
	return 1, nil
}

// testStmt is a test driver.Stmt.
type testStmt struct{}

// Close implements driver.Stmt.
func (s *testStmt) Close() error {
	return nil
}

// NumInput implements driver.Stmt.
func (s *testStmt) NumInput() int {
	return -1
}

// Exec implements driver.Stmt.
func (s *testStmt) Exec(args []driver.Value) (driver.Result, error) {
	return &testResult{}, nil
}

// Query implements driver.Stmt.
func (s *testStmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, fmt.Errorf("not supported")
}

// testConn is a test driver.Conn.
type testConn struct{}

// Prepare implements driver.Conn.
func (c *testConn) Prepare(query string) (driver.Stmt, error) {
	return &testStmt{}, nil
}

// Close implements driver.Conn.
func (c *testConn) Close() error {
	return nil
}

// Begin implements driver.Conn.
func (c *testConn) Begin() (driver.Tx, error) {
	return nil, fmt.Errorf("not supported")
}

// testDriver is test driver.Driver.
type testDriver struct{}

// Open implements driver.Driver.
func (db *testDriver) Open(name string) (driver.Conn, error) {
	return &testConn{}, nil
}

func TestIssue390(t *testing.T) {
	sql.Register("testDriver", &testDriver{})
	db, err := sql.Open("testDriver", "test.db")
	require.NoError(t, err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS log (message)")
	require.NoError(t, err)

	vm := otto.New()
	vm.Set("db", db)
	val, err := vm.Run(`
		db.Exec("CREATE TABLE log (message)")
		var results = db.Exec("INSERT INTO log(message) VALUES(?)", "test123");
		var res = results[0];
		var err = results[1];
		if (typeof err !== 'undefined') {
			result = err
		} else {
			results = res.RowsAffected()
			var rows = results[0];
			var err = results[1];
			if (typeof err !== 'undefined') {
				result = err
			} else {
				result = rows;
			}
		}
		result`,
	)
	require.NoError(t, err)
	rows, err := val.ToInteger()
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)
}
