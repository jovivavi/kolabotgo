package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"time"
)

var token string
var botId string
var channelId string

func main() {
	flag.StringVar(&token, "token", "", "Discord bot token")
	flag.StringVar(&botId, "botId", "", "Bot user id")
	flag.StringVar(&channelId, "channelId", "", "Channel id for message sending")
	flag.Parse()

	discord, err := discordgo.New("Bot " + token)
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
	if m.Author.ID == botId {
		return
	}
	words := strings.Split(m.Content, " ")
	if len(words) == 1 {
		return
	}
	wordPlace := rand.Intn(len(words))

	if len(words[wordPlace]) > 8 {
		if strings.HasPrefix(words[wordPlace], "<") || strings.HasPrefix(words[wordPlace], "http") {
			words[wordPlace] = "pylly"
		} else {
			if rand.Intn(2) == 1 {
				words[wordPlace] = words[wordPlace][0:len(words[wordPlace])-5] + "pylly"
			} else {
				words[wordPlace] = "pylly" + words[wordPlace][len(words[wordPlace])-5:]
			}
		}

	} else {
		words[wordPlace] = "pylly"
	}

	replacedWords := strings.Join(words, " ")
	if rand.Intn(10) == 1 {
		channel, _ := s.State.Channel(m.ChannelID)
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    m.Author.Username,
				IconURL: m.Author.AvatarURL("128"),
			},
			Color: s.State.UserColor(m.Author.ID, m.ChannelID),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "In " + channel.Name + ":",
					Value:  replacedWords,
					Inline: true,
				},
			},
		}

		s.ChannelMessageSendEmbed(channelId, embed)
	}

}
