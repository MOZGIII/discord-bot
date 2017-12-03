package main

import (
	"io"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/rylio/ytdl"
)

var options = encodeOptions()

func encodeOptions() *dca.EncodeOptions {
	options := *dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = "lowdelay"
	return &options
}

func stream(voiceConnection *discordgo.VoiceConnection, videoURL string) error {
	videoInfo, err := ytdl.GetVideoInfo(videoURL)
	if err != nil {
		return err
	}

	format := videoInfo.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
	downloadURL, err := videoInfo.GetDownloadURL(format)
	if err != nil {
		return err
	}

	encodingSession, err := dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		return err
	}
	defer encodingSession.Cleanup()

	done := make(chan error)
	dca.NewStream(encodingSession, voiceConnection, done)
	err = <-done
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
