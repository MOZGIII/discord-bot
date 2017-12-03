package bot

import (
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
	if !strings.HasPrefix(raw, p.Prefix) {
		return nil, nil
	}
	raw = raw[len(p.Prefix):]

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
}
