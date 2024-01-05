package main

import (
	"github.com/bwmarrin/discordgo"
)

type player struct {
	id   string
	Name string
}

type game struct {
	Player1 Player
	Player2 Player
}

func NewGame(player1ID string, player1Name string, player2ID string, player2Name string) *Game {
	return &Game{
		Player1: Player{
			id:   player1ID,
			Name: player1Name,
		},
		Player2: Player{
			id:   player2ID,
			Name: player2Name,
		},
	}
}

func PlayGame(s *discordgo.Session, m *discordgo.MessageCreate, game *Game) {

}
