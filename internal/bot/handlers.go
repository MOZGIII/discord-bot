package bot

import (
	"fmt"

	"github.com/MOZGIII/discord-bot/internal/youtube"

	discord "github.com/bwmarrin/discordgo"
)

// AddHandlers registers our handlers for the provided session.
func AddHandlers(session *discord.Session) {
	session.AddHandler(handleReady)
	session.AddHandler(handleMessageCreate)
	session.AddHandler(handleGuildCreate)
}

func handleReady(s *discord.Session, event *discord.Ready) {
	reportErrorInAny(s.UpdateStatus(0, "Hello world"))
}

// Called for every incoming message.
func handleMessageCreate(s *discord.Session, m *discord.MessageCreate) {
	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, err := comparse.ParseCommand(m.Content)
	if err != nil {
		reportError(err)
		return
	}

	// Do nothing if command in not known to us.
	if command == nil {
		return
	}

	if command.Command == "play" {
		// Determine video URL.
		if command.Args == "" {
			// No argument given.
			reportError(fmt.Errorf("play error: no video specified"))
			return
		}
		input := command.Args

		videoID, err := youtube.DefaultClient.Resolve(input)
		if err != nil {
			// Error determining video id.
			reportError(err)
			return
		}
		videoURL := youtube.VideoURL(videoID)

		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			reportError(fmt.Errorf("play error: could not find channel: %s", err))
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			reportError(fmt.Errorf("play error: could not find guild: %s", err))
			return
		}

		// Look for the message sender in that guild's current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				err = playSound(s, g.ID, vs.ChannelID, videoURL)
				if err != nil {
					reportError(fmt.Errorf("play error: error playing sound: %s", err))
				}
				return
			}
		}
		reportError(fmt.Errorf("play error: the command author not found in any voice channel within the guild"))
	}
}

// Called everytime bot joins a server (a guild).
func handleGuildCreate(s *discord.Session, event *discord.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Bot is ready!")
			return
		}
	}
}
