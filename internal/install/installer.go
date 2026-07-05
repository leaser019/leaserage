package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vomkhang/leaserage/internal/providers"
	"github.com/vomkhang/leaserage/internal/templates"
)

func BuildPlan(home string, provider providers.Provider) (Plan, error) {
	plan := Plan{ProviderID: provider.ID}
	for _, file := range provider.Files {
		body, err := templates.Read(file.TemplatePath)
		if err != nil {
			return plan, err
		}
		plan.Ops = append(plan.Ops, Operation{
			Kind:   OpWriteFile,
			Target: filepath.Join(home, filepath.FromSlash(file.TargetPath)),
			Body:   body,
			Mode:   file.MergeMode,
		})
	}

	skillFiles, err := templates.SkillFiles()
	if err != nil {
		return plan, err
	}
	for _, skillFile := range skillFiles {
		body, err := templates.Read(skillFile)
		if err != nil {
			return plan, err
		}
		name := filepath.Base(filepath.Dir(skillFile))
		plan.Ops = append(plan.Ops, Operation{
			Kind:   OpCopySkill,
			Target: filepath.Join(home, filepath.FromSlash(provider.SkillDir), name, "SKILL.md"),
			Body:   body,
			Mode:   providers.MergeReplaceManaged,
		})
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
			if err := WriteManagedFile(op.Target, op.Body, opts.DryRun); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported operation: %s", op.Kind)
		}
	}
	return nil
}
