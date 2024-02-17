package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func main() {

	// sudo chmod -R 777 ./postgres_data

	time.Sleep(time.Second * 2)

	var host = "game_postgres"
	var port = "5432"
	var user = "gameDbUser"
	var password = "bebra"
	var dbname = "gameDb"
	var sslmode = "disable"

	var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sqlx.Open("postgres", dbInfo)

}
