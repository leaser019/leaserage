package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vomkhang/leaserage/internal/providers"
	"github.com/vomkhang/leaserage/internal/templates"
)

func BuildPlan(home string, provider providers.Provider, opts PlanOptions) (Plan, error) {
	if opts.MCPMode == "" {
		opts.MCPMode = MCPDefault
	}
	if opts.Hook == "" {
		opts.Hook = HookNone
	}

	plan := Plan{ProviderID: provider.ID}
	for _, file := range provider.Files {
		if opts.MCPMode == MCPNone && file.Kind == providers.FileMCP {
			if file.NoMCPTemplatePath == "" {
				continue
			}
			file.TemplatePath = file.NoMCPTemplatePath
		}
		body, err := templates.Read(file.TemplatePath)
		if err != nil {
			return plan, err
		}
		plan.Ops = append(plan.Ops, Operation{
			Kind:   OpWriteFile,
			Target: filepath.Join(home, filepath.FromSlash(file.TargetPath)),
			Body:   body,
			Mode:   file.MergeMode,
			Perm:   0o644,
		})
	}

	skillFiles, err := templates.SkillFiles()
	if err != nil {
		return plan, err
	}
	for _, skillFile := range skillFiles {
		body, err := templates.Read(skillFile.Path)
		if err != nil {
			return plan, err
		}
		relativeSkillPath, err := filepath.Rel("skills", filepath.FromSlash(skillFile.Path))
		if err != nil {
			return plan, err
		}
		plan.Ops = append(plan.Ops, Operation{
			Kind:   OpCopySkill,
			Target: filepath.Join(home, filepath.FromSlash(provider.SkillDir), relativeSkillPath),
			Body:   body,
			Mode:   providers.MergeReplaceManaged,
			Perm:   normalizedPerm(skillFile.Perm),
		})
	}

	if opts.Hook == HookRTK {
		canRunRTK := opts.RTKInstalled
		if !opts.RTKInstalled && (opts.RuntimeOS == "linux" || opts.RuntimeOS == "darwin") {
			plan.Ops = append(plan.Ops, Operation{
				Kind:    OpRunCommand,
				Command: []string{"sh", "-c", "curl -fsSL https://raw.githubusercontent.com/rtk-ai/rtk/master/install.sh | sh"},
			})
			canRunRTK = true
		}
		if command := rtkInitCommand(provider.ID); canRunRTK && len(command) > 0 {
			plan.Ops = append(plan.Ops, Operation{Kind: OpRunCommand, Command: command})
		}
	}
	return plan, nil
}

func Apply(plan Plan, opts Options) error {
	for _, op := range plan.Ops {
		switch op.Kind {
		case OpWriteFile, OpCopySkill:
			if op.Mode == providers.MergeCreateOnly {
				if _, err := os.Stat(op.Target); err == nil {
					continue
				} else if !os.IsNotExist(err) {
					return err
				}
			}
			if err := WriteManagedFile(op.Target, op.Body, op.Perm, opts.DryRun); err != nil {
				return err
			}
		case OpRunCommand:
			if opts.DryRun {
				continue
			}
			if len(op.Command) == 0 {
				return fmt.Errorf("empty command operation")
			}
			cmd := exec.Command(op.Command[0], op.Command[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			if opts.Home != "" {
				cmd.Env = append(os.Environ(), "PATH="+filepath.Join(opts.Home, ".local", "bin")+string(os.PathListSeparator)+os.Getenv("PATH"))
			}
			if err := cmd.Run(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported operation: %s", op.Kind)
		}
	}
	return nil
}

func normalizedPerm(perm os.FileMode) os.FileMode {
	if perm&0o111 != 0 {
		return 0o755
	}
	return 0o644
}

func rtkInitCommand(providerID string) []string {
	switch providerID {
	case "claude-code":
		return []string{"rtk", "init", "--global"}
	case "cursor":
		return []string{"rtk", "init", "--global", "--agent", "cursor"}
	case "opencode":
		return []string{"rtk", "init", "--global", "--opencode"}
	case "github-copilot":
		return []string{"rtk", "init", "--global", "--copilot"}
	case "codex":
		return []string{"rtk", "init", "--global", "--codex"}
	default:
		return nil
	}
}
