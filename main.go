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
	if strings.HasPrefix(m.Content, "!dan") {
		s.ChannelMessageSend(m.ChannelID, "Cock and Balls")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1193496949832298496/kiss-anime.gif?ex=65aceda6&is=659a78a6&hm=11bd484498f33f82f84c61a1a06cec7395f1897c881f40cd871b81e7859bbf79&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!john") {
		s.ChannelMessageSend(m.ChannelID, "Cock and Balls")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://media.discordapp.net/attachments/756680254705238063/875992997035139072/image0.gif?ex=65ab6b8d&is=6598f68d&hm=10a5b3627e273fdff47e0931b03975cafd396c9f6b81170779997566a0f6d6c7&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!aaron") {
		s.ChannelMessageSend(m.ChannelID, "Cock and Balls")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1193498365829328957/maidenless-bitchless.gif?ex=65aceef8&is=659a79f8&hm=4b21e130a162d66e2f98bb1dd040787e2df1fa598f54f426624d73a73118044f&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!jaden") {
		s.ChannelMessageSend(m.ChannelID, "Cock and Balls")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1194437211953844234/yappin.gif?ex=65b05956&is=659de456&hm=4e6d076c416381c06868ba2b4cebbeda98ac65fadae7ef80cd1b7cc0d8a39d1a&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!rax") {
		s.ChannelMessageSend(m.ChannelID, "goat")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1193496023226646539/goat.gif?ex=65acecca&is=659a77ca&hm=5be62955ea4ddd5a8a34324752a4a224a4f3019debfa4e21b38afae60b749d56&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!chad") {
		s.ChannelMessageSend(m.ChannelID, "boba fondler")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1194436148408356944/fondle.gif?ex=65b05859&is=659de359&hm=ebcea77672eaedfae33a951356095c254365020510200cc7ae64c3488f7f7644&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!jocelyn") {
		s.ChannelMessageSend(m.ChannelID, "goat #2")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1193499882636447816/richarlison-tottenham.gif?ex=65acf062&is=659a7b62&hm=329f65ba726087b5cf31ead49935fe4551296b53f7b70fc0bc8d53f175c123ba&",
				},
			},
		})
	}
	if strings.HasPrefix(m.Content, "!tiffany") {
		s.ChannelMessageSend(m.ChannelID, "Cock and Balls")
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/821968732209283073/1194439177366933624/amonguss.gif?ex=65b05b2b&is=659de62b&hm=e582f267d05a12917782494ab87d051ff4ac77545039ce0de0a9569326370b4c&",
				},
			},
		})
	}

}

func main() {
	config := LoadConfig("config.json")

	dg := CreateBot(config.Token)

	appID := "1190886365458604055"
	guildID := "925145676407537724"

	dg.AddHandler(messageHandler)
	CreateCommand(dg, appID, guildID)
	dg.AddHandler(SongOfTheDayCommandHandler)
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages from the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Check if the message starts with "!rps"
		if strings.HasPrefix(m.Content, "!1v1") {
			// Extract the mentioned user
			if len(m.Mentions) == 0 {
				s.ChannelMessageSend(m.ChannelID, "You need to mention a user to play with!")
				return
			}
			mentionedUser := m.Mentions[0]

			// Create a new game
			game := NewGame(m.Author.ID, m.Author.Username, mentionedUser.ID, mentionedUser.Username)

			// Start the game
			PlayGame(s, m, game)
		}
	})

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "dream":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "When the ants come at me, It'll take 10,000, 100,000 of them to take me down. So that's how miniscule you are to my size, right. My stature of intelligence, character and body and um, Reverence in this world. No man, because I'm gonna do movies, stand up comedy, everything all the shit, music, whatever the fuck e commerce, You don't understand that you're talking to like a Michealangelo of my time, right. Like I'm a genius, Albert Einstein level, History book maker. You're gonna be forgotten like the dust in the sand when you're in the fucking sahara, and there's a hundred million thousand billion sand particles, you're gonna be one of those, and I'm gonna be a staute erected in gold.",
					},
				})
			case "kitten":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "🫵",
					},
				})
			}
		}
	})

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
