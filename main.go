package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"time"
)


func main() {
	discord, err := discordgo.New("Bot " + "")
	if err != nil {
		fmt.Println("error creating session", err)
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection, ", err)
		return
	}

	fmt.Println("bot running")
	<-make(chan struct{})
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
	if m.ChannelID == "168311106950004737" {
		s.ChannelTyping("168311106950004737")
		s.MessageReactionAdd(m.ChannelID, m.ID, ":babbu:230233791665405953")
	}
}