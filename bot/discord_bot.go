package bot

import (
	"fmt"
	"godiscopwn/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var botCallCount int

func StartDiscordBot(token string) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("\nBot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!botcall" {
		botCallCount++
		fmt.Printf("Bot call count: %d\n", botCallCount)
		// Broadcast the bot call count and message to WebSocket clients
		server.BroadcastBotCallCount(botCallCount, m.Content)
	}
}
