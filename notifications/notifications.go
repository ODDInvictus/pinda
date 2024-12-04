package notifications

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/clarkmcc/go-reroutine"
	"github.com/oddinvictus/pinda/db"
	"github.com/redis/go-redis/v9"
)

type NewActivityJob struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Date string `json:"date"`
	Type string `json:"type"`
}

func Init(rdb *redis.Client, client *db.PrismaClient) {
	stop := make(chan struct{})

	reroutine.Go(stop, func() {
		go newActivity(rdb, client)
	})
	
}

func info(msg ...interface{}) {
	log.Println(append([]interface{}{"[notifications]"}, msg...)...)
}

func newActivity(rdb *redis.Client, client *db.PrismaClient) {
	ctx := context.Background()
	pubsub := rdb.Subscribe(ctx, "new-activity")

	discord := Discord{}

	info("listening for \"new-activity\"")

	for {
		msg, err := pubsub.ReceiveMessage(ctx)

		if err != nil {
			panic(err)
		}

		job := NewActivityJob{}
		json.Unmarshal([]byte(msg.Payload), &job)

		info("new activity job for id:", job.Data)

		aid, err := strconv.Atoi(job.Data)

		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		activity, err := client.Activity.
			FindFirst(
				db.Activity.ID.Equals(aid),
			).With(
				db.Activity.Location.Fetch(),
			).
			Exec(ctx)

		if err != nil {
			panic(err)
		}

		discord.SendNewActivity(activity)
	}
}