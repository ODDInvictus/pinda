package notifications

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gtuk/discordwebhook"
	"github.com/oddinvictus/pinda/db"
)

type Discord struct {}

func (d *Discord) SendNewActivity(activity *db.ActivityModel) {
	username := "Miel Monteur"
	url := os.Getenv("DISCORD_WEBHOOK")
	ibsUrl := os.Getenv("IBS_URL")

	if url == "" {
		panic("env DISCORD_WEBHOOK unset")
	}

	if ibsUrl == "" {
		panic("env IBS_URL unset")
	}

	title := "Nieuwe activiteit: " + activity.Name
	color := "5577610"

	date := "Datum"
	loc, _ := time.LoadLocation("Europe/Amsterdam")
	time := activity.StartTime.In(loc).Format("02 January 2006 15:04")

	locName := "Locatie"
	location, _ := activity.Location()
	locValue := location.Name

	infoName := "Meer informatie en aanmelden"
	infoValue := ibsUrl + "/activiteit/" + strconv.Itoa(activity.ID)

	embeds := make([]discordwebhook.Embed, 0) 
	embeds = append(embeds, discordwebhook.Embed{
		Title: &title,
		Color: &color,
		Description: &activity.Description,
		Fields: &[]discordwebhook.Field{
			{
				Name: &date,
				Value: &time,
			},
			{
				Name: &locName,
				Value: &locValue,
			},
			{
				Name: &infoName,
				Value: &infoValue,
			},
		},
	})

	message := discordwebhook.Message{
		Username: &username,
		Embeds: &embeds,
	}

	err := discordwebhook.SendMessage(url, message)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}