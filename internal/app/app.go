package app

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vomkhang/leaserage/internal/cli"
	"github.com/vomkhang/leaserage/internal/install"
	"github.com/vomkhang/leaserage/internal/providers"
)

const Version = "0.1.0-dev"

type App struct {
	Stdout io.Writer
	Stderr io.Writer
}

func New(stdout io.Writer, stderr io.Writer) *App {
	return &App{Stdout: stdout, Stderr: stderr}
}

func (a *App) Run(args []string) int {
	cmd, err := cli.Parse(args)
	if err != nil {
		fmt.Fprintf(a.Stderr, "%v\n", err)
		return 2
	}

	switch cmd.Name {
	case "help":
		fmt.Fprint(a.Stdout, HelpText())
		return 0
	case "version":
		fmt.Fprintf(a.Stdout, "leaserage %s\n", Version)
		return 0
	case "providers":
		for _, id := range providers.IDs() {
			fmt.Fprintln(a.Stdout, id)
		}
		return 0
	case "install":
		return a.install(cmd)
	case "doctor":
		return a.doctor(cmd)
	case "uninstall":
		return a.uninstall(cmd)
	default:
		fmt.Fprintf(a.Stderr, "unknown command: %s\n\n%s", cmd.Name, HelpText())
		return 2
	}
}

func (a *App) install(cmd cli.Command) int {
	home, selected, ok := a.resolveProviderCommand("install", cmd)
	if !ok {
		return 2
	}

	for _, provider := range selected {
		plan, err := install.BuildPlan(home, provider)
		if err != nil {
			fmt.Fprintf(a.Stderr, "plan %s: %v\n", provider.ID, err)
			return 1
		}
		if err := install.Apply(plan, install.Options{Home: home, DryRun: cmd.DryRun}); err != nil {
			fmt.Fprintf(a.Stderr, "install %s: %v\n", provider.ID, err)
			return 1
		}
		status := "installed"
		if cmd.DryRun {
			status = "planned"
		}
		fmt.Fprintf(a.Stdout, "%s %s (%d operations)\n", status, provider.ID, len(plan.Ops))
	}
	return 0
}

func (a *App) doctor(cmd cli.Command) int {
	home, selected, ok := a.resolveProviderCommand("doctor", cmd)
	if !ok {
		return 2
	}

	allOK := true
	for _, provider := range selected {
		for _, result := range install.Doctor(home, provider) {
			status := "ok"
			if !result.OK {
				status = "missing"
				allOK = false
			}
			if result.Info == "" {
				fmt.Fprintf(a.Stdout, "%s provider=%s check=%s kind=%s\n", status, provider.ID, result.Name, result.Kind)
			} else {
				fmt.Fprintf(a.Stdout, "%s provider=%s check=%s kind=%s detail=%s\n", status, provider.ID, result.Name, result.Kind, result.Info)
			}
		}
	}
	if !allOK {
		return 1
	}
	return 0
}

func (a *App) uninstall(cmd cli.Command) int {
	home, selected, ok := a.resolveProviderCommand("uninstall", cmd)
	if !ok {
		return 2
	}

	for _, provider := range selected {
		plan, err := install.BuildPlan(home, provider)
		if err != nil {
			fmt.Fprintf(a.Stderr, "plan %s: %v\n", provider.ID, err)
			return 1
		}
		if err := install.RemoveFiles(plan, cmd.DryRun); err != nil {
			fmt.Fprintf(a.Stderr, "uninstall %s: %v\n", provider.ID, err)
			return 1
		}
		status := "removed"
		if cmd.DryRun {
			status = "planned removal"
		}
		fmt.Fprintf(a.Stdout, "%s %s (%d operations)\n", status, provider.ID, len(plan.Ops))
	}
	return 0
}

func (a *App) resolveProviderCommand(name string, cmd cli.Command) (string, []providers.Provider, bool) {
	if len(cmd.Provider) == 0 {
		fmt.Fprintf(a.Stderr, "%s requires --provider\n", name)
		return "", nil, false
	}

	home := cmd.Home
	if home == "" {
		resolved, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(a.Stderr, "resolve home: %v\n", err)
			return "", nil, false
		}
		home = resolved
	}

	selected, unknown := providers.Select(cmd.Provider)
	if len(unknown) > 0 {
		fmt.Fprintf(a.Stderr, "unknown provider(s): %s\n", strings.Join(unknown, ","))
		return "", nil, false
	}
	return home, selected, true
}

func HelpText() string {
	return `leaserage installs personal agent workflow configs.

Usage:
  leaserage install --provider opencode,kilo [--dry-run] [--home /tmp/home]
  leaserage doctor --provider opencode,kilo [--home /tmp/home]
  leaserage uninstall --provider opencode,kilo [--dry-run] [--home /tmp/home]
  leaserage providers
  leaserage version
`
}
