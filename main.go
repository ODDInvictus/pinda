package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/oddinvictus/pinda/db"
	"github.com/oddinvictus/pinda/image"
	"github.com/oddinvictus/pinda/notifications"
	"github.com/redis/go-redis/v9"
)

func main() {
	/**
		DOTENV
	*/
	err := godotenv.Load(filepath.Join(os.Getenv("APP_PATH"), ".env"))

	if err != nil {
		panic(err)
	}

	/**
		DATABASE
	*/
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	/**
		REDIS
	*/

	redisURL := os.Getenv("REDIS_URL")

	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
		Password: "",
		DB: 0,
	})

	image.Init(rdb)
	notifications.Init(rdb, client)

	/**
		HTTP Routing
	*/
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Invictus Bier Systeem API")
	})

	log.Println("Listening on http://127.0.0.1:3000")
	err = http.ListenAndServe(":3000", router)

	if errors.Is(err, http.ErrServerClosed) {
		log.Panicln("server closed")
	} else if err != nil {
		log.Panicln("error starting server", err)
	}
}