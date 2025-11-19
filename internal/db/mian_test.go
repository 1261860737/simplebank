// 设置连接 和 查询对象  testqueries对象为全局对象
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
)

// 特定Golang包中所有的单元测试入口 如db
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("cannot conn to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}