package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"time"
	"strings"
	"math/rand"
)

func main() {
	discord, err := discordgo.New("Bot " + "MTY4MzA0OTgxOTY1Nzk5NDI0.C_jhaA.0khhXhj3QqeZ7OXQiIAbFdo1kHw")
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
	if m.Author.ID != "168304981965799424" {
		words := strings.Split(m.Content, " ")
		if len(words) == 1 {
			return
		}
		wordPlace := rand.Intn(len(words))
		if len(words[wordPlace]) > 8 {
			if rand.Intn(1) == 1 {
				words[wordPlace] = words[wordPlace][0:len(words[wordPlace])-5] + "pylly"
			} else {
				words[wordPlace] = "pylly" + words[wordPlace][len(words[wordPlace])-5:]
			}

		} else {
			words[wordPlace] = "pylly"
		}

		replacedWords := strings.Join(words, " ")
		if rand.Intn(10) == 1 {
			s.ChannelMessageSend("215600297802727426", m.Author.Username + ": " +replacedWords)
		}

	}
}
