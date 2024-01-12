package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var songList = []string{
	"https://music.youtube.com/watch?v=2PVUZ5ZX79Q&si=E5-rQ8Y98YgZRXqI",
	"https://music.youtube.com/watch?v=sQMpD7bLYtc&si=CK91-vdn58lk0vnt",
	"https://music.youtube.com/watch?v=N1cJFgSuvfY&si=zMBBugA970gLsNM1",
	"https://music.youtube.com/watch?v=IFA92iTE95w&si=zBhw_qnHHhnhPEY9",
	"https://music.youtube.com/watch?v=peUjZ3n18Uk&si=9iGNCC9WzObbGx_z",
	"https://music.youtube.com/watch?v=2uRU2iqcqCU&si=bZQtn55ClP52k1qu",
	"https://music.youtube.com/watch?v=ea5w1VqzVIY&si=3cGx5Z7fKPD4O276",
}

func SongOfTheDayCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "song" {
			// User Current date to calculate the index of the song
			t := time.Now()
			dateSeed := t.Year()*10000 + int(t.Month())*100 + t.Day()
			songIndex := dateSeed % len(songList)
			songOfTheDay := songList[songIndex]
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: songOfTheDay,
				},
			})
		}
	}
}
