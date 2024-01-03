package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func CreateCommand(dg *discordgo.Session, appID string, guildID string) {
	_, err := dg.ApplicationCommandCreate(appID, guildID, &discordgo.ApplicationCommand{
		Name:        "dream",
		Description: "Trigger the dream command",
	})
	if err != nil {
		log.Fatalf("Unable to create command: %v", err)
	}
}

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
