package cli

import (
	"errors"
	"strings"
)

type Command struct {
	Name     string
	Provider []string
	DryRun   bool
	Home     string
	Force    bool
	Verbose  bool
}

func Parse(args []string) (Command, error) {
	if len(args) == 0 {
		return Command{Name: "help"}, nil
	}

	cmd := Command{Name: args[0]}
	if cmd.Name == "--help" || cmd.Name == "-h" {
		cmd.Name = "help"
	}

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--provider":
			if i+1 >= len(args) {
				return cmd, errors.New("--provider requires a value")
			}
			cmd.Provider = splitCSV(args[i+1])
			i++
		case "--dry-run":
			cmd.DryRun = true
		case "--home":
			if i+1 >= len(args) {
				return cmd, errors.New("--home requires a value")
			}
			cmd.Home = args[i+1]
			i++
		case "--force":
			cmd.Force = true
		case "--verbose":
			cmd.Verbose = true
		case "--help", "-h":
			cmd.Name = "help"
		default:
			return cmd, errors.New("unknown flag: " + args[i])
		}
	}

	return cmd, nil
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}
