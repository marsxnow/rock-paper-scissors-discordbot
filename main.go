package main

import (
	"fmt"       // for printing
	"math/rand" // for random numbers
	"sort"
	"strings" // for string manipulation
	"time"    // for time manipulation

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
			return "Hold on to your horses young buck. (wait 5 seconds) ", ""
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
			return "OMEGA LOL.", ""
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
			} else if botChoice == "OMEGA LOL." {
				dmChannel, err := s.UserChannelCreate(m.Author.ID)
				if err != nil {
					fmt.Println("error creating DM channel,", err)
					return
				}
				s.ChannelMessageSend(dmChannel.ID, "You stink at this game.")
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
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dg, err := discordgo.New("Bot MTE5MDg4NjM2NTQ1ODYwNDA1NQ.G7Bg7z.RHwDB152Vm_uPLRJt6e2y5S9FVV-jxiTQwMIF4")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageHandler)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
