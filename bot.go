package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateBot(token string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return nil
	}

	return dg
}
