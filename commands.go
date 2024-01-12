package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Define command creation function
func CreateCommand(dg *discordgo.Session, appID string, guildID string) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "dream",
			Description: "Trigger the dream command",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "kitten",
			Description: "Trigger the kitten command",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "song",
			Description: "get the song of the day",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		//add more for later
	}
	for _, command := range commands {
		_, err := dg.ApplicationCommandCreate(appID, guildID, command)
		if err != nil {
			log.Fatalf("Unable to create command: %v", err)
		}
	}
}

// Define command handler function
func DreamCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "dream" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "When the ants come at me, It'll take 10,000, 100,000 of them to take me down. So that's how miniscule you are to my size, right. My stature of intelligence, character and body and um, Reverence in this world. No man, because I'm gonna do movies, stand up comedy, everything all the shit, music, whatever the fuck e commerce, You don't understand that you're talking to like a Michealangelo of my time, right. Like I'm a genius, Albert Einstein level, History book maker. You're gonna be forgotten like the dust in the sand when you're in the fucking sahara, and there's a hundred million thousand billion sand particles, you're gonna be one of those, and I'm gonna be a staute erected in gold.",
				},
			})
		}
	}
}

func KittenCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "kitten" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "🫵",
				},
			})
		}
	}
}

// func SongOfTheDayCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if i.Type == discordgo.InteractionApplicationCommand {
// 		data := i.ApplicationCommandData()
// 		if data.Name == "song" {
// 			songOfTheDayLink := "https://music.youtube.com/watch?v=2PVUZ5ZX79Q&si=E5-rQ8Y98YgZRXqI"

// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Content: songOfTheDayLink,
// 				},
// 			})
// 		}
// 	}
// }
