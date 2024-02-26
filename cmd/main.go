package main

import (
	"animals-game/internal/repository"
	"animals-game/internal/tg-bot"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {

	time.Sleep(time.Second * 1)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	var host = "game_postgres"
	var port = "5432"
	var user = "gameDbUser"
	var password = os.Getenv("DB_PASSWORD")
	var dbname = "gameDb"
	var sslmode = "disable"

	var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sqlx.Open("postgres", dbInfo)

	repo := repository.NewRepository(db)

	TgBot := tg_bot.NewTgBot(os.Getenv("TELEGRAM_APITOKEN"), repo)

	err = TgBot.Run()

	if err != nil {
		log.Fatal(err)
	}

}
