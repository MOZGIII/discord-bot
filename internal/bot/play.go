package bot

import (
	"time"

	discord "github.com/bwmarrin/discordgo"
)

// playSound plays the current buffer to the provided channel.
func playSound(s *discord.Session, guildID, channelID, videoURL string) (err error) {
	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	withSpeaking := func() error { return stream(vc, videoURL) }
	withSleepFn := func() error { return wrapWithSpeaking(withSpeaking, vc) }
	f := func() error { return wrapWithSleep(withSleepFn, 250*time.Millisecond, 250*time.Millisecond, false) }

	// Call the payload.
	err1 := f()

	// Disconnect from the provided voice channel.
	if err := vc.Disconnect(); err != nil && err1 == nil {
		return err
	}
	return err1
}

func wrapWithSleep(f func() error, before, after time.Duration, skipSleepOnError bool) error {
	time.Sleep(before)

	err := f()
	if err != nil && skipSleepOnError {
		return err
	}

	time.Sleep(after)
	return err
}

func wrapWithSpeaking(f func() error, vc *discord.VoiceConnection) error {
	// Start speaking.
	if err := vc.Speaking(true); err != nil {
		return err
	}

	// Play the sound.
	err1 := f()

	// Stop speaking.
	if err := vc.Speaking(false); err != nil && err1 == nil {
		return err
	}

	return err1
}
