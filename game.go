package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Player struct {
	ID   string
	Name string
}

type Game struct {
	Player1 *Player
	Player2 *Player
}

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

func NewGame(player1ID, player1Name, player2ID, player2Name string) *Game {
	return &Game{
		Player1: &Player{
			ID:   player1ID,
			Name: player1Name,
		},
		Player2: &Player{
			ID:   player2ID,
			Name: player2Name,
		},
	}
}

func (m Move) String() string {
	return [...]string{"Rock", "Paper", "Scissors"}[m]
}

func PlayGame(s *discordgo.Session, m *discordgo.MessageCreate, game *Game) {
	// Send message to channel that game is starting
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Starting game between %s and %s", game.Player1.Name, game.Player2.Name))

	// Wait for a message for either player
	player1Move := waitForMove(s, game.Player1.ID)
	player2Move := waitForMove(s, game.Player2.ID)

	// Determine winner
	winner := determineWinner(game, player1Move, player2Move)

	// Announce Winner
	if winner == game.Player1.ID {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s won!", game.Player1.Name))
	} else if winner == game.Player2.ID {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s won!", game.Player2.Name))
	} else {
		s.ChannelMessageSend(m.ChannelID, "It's a tie!")
	}
}

func waitForMove(s *discordgo.Session, playerID string) Move {
	moveCh := make(chan Move)

	// Add a message create event handler
	removeHandler := s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Check if the message is from the player
		if m.Author.ID == playerID {
			switch strings.ToLower(m.Content) {
			case "rock":
				moveCh <- Rock
			case "paper":
				moveCh <- Paper
			case "scissors":
				moveCh <- Scissors
			}
		}
	})

	// Wait for a move from the player
	move := <-moveCh

	// Remove the message create event handler
	removeHandler()

	return move
}

func determineWinner(game *Game, player1Move Move, player2Move Move) string {
	if player1Move == player2Move {
		return ""
	} else if (player1Move == Rock && player2Move == Scissors) || (player1Move == Scissors && player2Move == Paper) || (player1Move == Paper && player2Move == Rock) {
		return game.Player1.ID
	} else {
		return game.Player2.ID
	}
}
