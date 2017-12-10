package bot

import (
	"log"
	"os"
	"strings"
)

var comparse *commandParser

type commandParser struct {
	Prefix string
}

type parsedCommand struct {
	Command string
	Args    string
}

func (p *commandParser) ParseCommand(raw string) (*parsedCommand, error) {
	return parseCommand(raw, p.Prefix)
}

func parseCommand(raw, prefix string) (*parsedCommand, error) {
	if !strings.HasPrefix(raw, prefix) {
		return nil, nil
	}
	raw = raw[len(prefix):]

	split := strings.SplitN(raw, " ", 2)
	if len(split) < 1 {
		return nil, nil
	}

	command := split[0]
	args := ""
	if len(split) > 1 {
		args = split[1]
	}

	return &parsedCommand{
		Command: command,
		Args:    args,
	}, nil
}

func init() {
	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		prefix = "-"
	}

	comparse = &commandParser{
		Prefix: prefix,
	}
	log.Printf("Prefix set to: %q", comparse.Prefix)
}
