package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/yerlanov/simplebank/api"
	db "github.com/yerlanov/simplebank/db/sqlc"
	"github.com/yerlanov/simplebank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load configs: %s", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
