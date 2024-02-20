package main

import (
	tg_bot "animals-game/internal/tg-bot"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {

	time.Sleep(time.Second * 2)
	/*
		// testing work of db
		tgId := "123"
		username := "egores"
		password_hash := "aboba"
		animal_id := 1
		data := `INSERT INTO users(username, telegram_id, password_hash, animal_id) VALUES($1, $2, $3, $4);`
		fmt.Println("All's good 1")
		//Выполняем наш SQL запрос
		if _, err = db.Exec(data, username, tgId, password_hash, animal_id); err != nil {
			log.Fatal(err)
		}*/

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

	if err != nil {
		log.Fatal(err)
	}

	TgBot := tg_bot.NewTgBot(os.Getenv("TELEGRAM_APITOKEN"))
	TgBot.RepoInit(db)

	err = TgBot.Run()

	if err != nil {
		log.Fatal(err)
	}

}
