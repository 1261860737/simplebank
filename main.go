package main

import (
	"log"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/chen/simplebank/api"
	"github.com/chen/simplebank/internal/db"
)

// serverAddress 即 localhost 端口号为8080
const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot conn to db:", err)
	}

	store := db.NewStore(conn)     // 连接数据库
	server := api.NewServer(store) //连接api服务器

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
