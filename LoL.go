package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
)

func LolCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "lol" {
			username := data.Options[0].StringValue()

			// Instantiate default collector
			c := colly.NewCollector()

			// Define variables to store wins and losses
			var wins, losses int

			// On every element with the class "win-lose", get the text
			c.OnResponse(func(r *colly.Response) {
				fmt.Println("Visited", r.Request.URL)
				fmt.Println(string(r.Body))
			})

			c.OnHTML("body", func(e *colly.HTMLElement) {
				// Get the entire body text
				bodyText := e.Text

				// Count the number of wins and losses
				wins = strings.Count(bodyText, `"result":"WIN"`)
				losses = strings.Count(bodyText, `"result":"LOSE"`)
			})

			c.Visit(fmt.Sprintf("https://na.op.gg/summoner/userName=%s", username))
			c.Wait()
			fmt.Printf("%s has %d wins and %d losses\n", username, wins, losses)
			// Respond with the user's win/loss stats
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("%s has %d wins and %d losses", username, wins, losses),
				},
			})
		}
	}
}
