package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MOZGIII/discord-bot/internal/bot"

	"github.com/bwmarrin/discordgo"
)

func crash(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		crash("No token provided. Please set DISCORD_TOKEN evironment variable.")
	}

	// Create a new Discord session using the provided bot token.
	// Bot tokens are prefixed with "Bot ".
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		crash("Error creating Discord session: %s", err)
	}

	bot.AddHandlers(dg)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		crash("Error opening Discord session: %s", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	if err := dg.Close(); err != nil {
		crash("Error closing Discord session: %s", err)
	}
}
