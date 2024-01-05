package main

import (
	"fmt"       // for printing
	"math/rand" // for random numbers
	"sort"      // for sorting
	"strings"   // for string manipulation
	"time"      // for time manipulation

	// for logging

	"github.com/bwmarrin/discordgo" // for discord
)

var choices = map[string]int{
	"rock":     0,
	"paper":    1,
	"scissors": 2,
}

var scores = make(map[string]*Score)

var lastPlayed = make(map[string]time.Time)

var losingStreaks = make(map[string]int)

type UserScore struct {
	ID    string
	Score *Score
}

type Score struct {
	Wins   int
	Losses int
	Ties   int
}

func showLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	var userScores []UserScore
	for id, score := range scores {
		userScores = append(userScores, UserScore{id, score})
	}
	sort.Slice(userScores, func(i, j int) bool {
		return userScores[i].Score.Wins > userScores[j].Score.Wins
	})
	leaderboard := "Leaderboard:\n"
	for i, userScore := range userScores {
		if i >= 10 { //Display the top 10 scores
			break
		}
		leaderboard += fmt.Sprintf("%d. <@%s>: %d wins, %d losses, %d ties\n", i+1, userScore.ID, userScore.Score.Wins, userScore.Score.Losses, userScore.Score.Ties)
	}
	s.ChannelMessageSend(m.ChannelID, leaderboard)
}

func playRPS(userChoice string, userID string) (string, string) {
	now := time.Now()
	if lastPlayed, ok := lastPlayed[userID]; ok {
		if now.Sub(lastPlayed) < time.Second*1 {
			return "Hold on to your horses young buck. (wait a seconds) ", ""
		}
	}
	lastPlayed[userID] = now
	botChoice := rand.Intn(3)
	result := (choices[userChoice] - botChoice + 3) % 3
	botChoiceString := ""
	for k, v := range choices {
		if v == botChoice {
			botChoiceString = k
			break
		}
	}
	if _, ok := scores[userID]; !ok {
		scores[userID] = &Score{}
	}

	if result == 1 {
		losingStreaks[userID] = 0
	} else if result == 2 {
		losingStreaks[userID]++
		if losingStreaks[userID] >= 2 {
			return "OMEGA LOL YOU LOST TWICE.", botChoiceString
		}
	} else {
		losingStreaks[userID] = 0
	}

	switch result {
	case 0:
		scores[userID].Ties++
		return "It's a tie!", botChoiceString
	case 1:
		scores[userID].Wins++
		return "You win!", botChoiceString
	case 2:
		scores[userID].Losses++
		return "You lose!", botChoiceString
	default:
		return "Error man idk bussy bussy", botChoiceString
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!rps ") {
		userChoice := strings.TrimPrefix(m.Content, "!rps ")
		if _, ok := choices[userChoice]; ok {
			result, botChoice := playRPS(userChoice, m.Author.ID)
			if botChoice == "" {
				s.ChannelMessageSend(m.ChannelID, result)
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s Bot chose %s.", result, botChoice))
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Invalid choice, please choose rock, paper, or scissors.")
		}

	} else if strings.HasPrefix(m.Content, "!stats") {
		score, ok := scores[m.Author.ID]
		if !ok {
			s.ChannelMessageSend(m.ChannelID, "You haven't played yet!")
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Wins: %d\nLosses: %d\nTies: %d", score.Wins, score.Losses, score.Ties))
		}
	}
	if strings.HasPrefix(m.Content, "!leaderboard") {
		showLeaderboard(s, m)
	}
	if strings.HasPrefix(m.Content, "!dream") {
		s.ChannelMessageSend(m.ChannelID, "When the ants come at me, It'll take 10,000, 100,000 of them to take me down. So that's how miniscule you are to my size, right. My stature of intelligence, character and body and um, Reverence in this world. No man, because I'm gonna do movies, stand up comedy, everything all the shit, music, whatever the fuck e commerce, You don't understand that you're talking to like a Michealangelo of my time, right. Like I'm a genius, Albert Einstein level, History book maker. You're gonna be forgotten like the dust in the sand when you're in the fucking sahara, and there's a hundred million thousand billion sand particles, you're gonna be one of those, and I'm gonna be a staute erected in gold.")
	}
}

func main() {
	config := LoadConfig("config.json")

	dg := CreateBot(config.Token)

	appID := "1190886365458604055"
	guildID := "1192264999071125545"

	dg.AddHandler(messageHandler)
	CreateCommand(dg, appID, guildID)

	// dg.AddHandler(DreamCommandHandler(dg))
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "When the ants come at me, It'll take 10,000, 100,000 of them to take me down. So that's how miniscule you are to my size, right. My stature of intelligence, character and body and um, Reverence in this world. No man, because I'm gonna do movies, stand up comedy, everything all the shit, music, whatever the fuck e commerce, You don't understand that you're talking to like a Michealangelo of my time, right. Like I'm a genius, Albert Einstein level, History book maker. You're gonna be forgotten like the dust in the sand when you're in the fucking sahara, and there's a hundred million thousand billion sand particles, you're gonna be one of those, and I'm gonna be a staute erected in gold.",
				},
			})
		}
	})

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
