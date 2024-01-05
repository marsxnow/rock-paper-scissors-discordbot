package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Define command creation function
func CreateCommand(dg *discordgo.Session, appID string, guildID string) {
	_, err := dg.ApplicationCommandCreate(appID, guildID, &discordgo.ApplicationCommand{
		Name:        "dream",
		Description: "Trigger the dream command",
	})
	if err != nil {
		log.Fatalf("Unable to create command: %v", err)
	}
	_, pog := dg.ApplicationCommandCreate(appID, guildID, &discordgo.ApplicationCommand{
		Name:        "1v1",
		Description: "1v1 someone in rock paper scissors",
	})
	if pog != nil {
		log.Fatalf("Unable to create command: %v", pog)
	}
}

// Define command handler function
func DreamCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "When the ants come at me, It'll take 10,000, 100,000 of them to take me down. So that's how miniscule you are to my size, right. My stature of intelligence, character and body and um, Reverence in this world. No man, because I'm gonna do movies, stand up comedy, everything all the shit, music, whatever the fuck e commerce, You don't understand that you're talking to like a Michealangelo of my time, right. Like I'm a genius, Albert Einstein level, History book maker. You're gonna be forgotten like the dust in the sand when you're in the fucking sahara, and there's a hundred million thousand billion sand particles, you're gonna be one of those, and I'm gonna be a staute erected in gold.",
			},
		})
	}
}

func StartGameCommandHandler(s *discordgo.Session) func(*discordgo.InteractionCreate) {
	return func(i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			data := i.Data.(*discordgo.ApplicationCommandInteractionData)
			if data.Name == "startgame" {
				// Extract user IDs from command arguments
				player1ID := data.Options[0].UserValue(s).ID
				player1Name := data.Options[0].UserValue(s).Username
				player2ID := data.Options[1].UserValue(s).ID
				player2Name := data.Options[1].UserValue(s).Username

				// Create a new game
				game := NewGame(player1ID, player1Name, player2ID, player2Name)

				// Start the game
				PlayGame(s, i.Message, game)
			}
		}
	}
}
