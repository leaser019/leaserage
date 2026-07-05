package providers

import "sort"

func All() []Provider {
	return []Provider{
		{
			ID: "opencode", DisplayName: "OpenCode", Root: ".config/opencode", SkillDir: ".config/opencode/skills",
			Files: []ProviderFile{
				{TemplatePath: "providers/opencode/opencode.jsonc", TargetPath: ".config/opencode/opencode.jsonc", MergeMode: MergeReplaceManaged},
				{TemplatePath: "agents/AGENTS.md", TargetPath: ".config/opencode/AGENTS.md", MergeMode: MergeCreateOnly},
			},
		},
		{
			ID: "kilo", DisplayName: "Kilo CLI", Root: ".config/kilo", SkillDir: ".config/kilo/skills",
			Files: []ProviderFile{
				{TemplatePath: "providers/kilo/kilo.jsonc", TargetPath: ".config/kilo/kilo.jsonc", MergeMode: MergeReplaceManaged},
				{TemplatePath: "agents/AGENTS.md", TargetPath: ".config/kilo/AGENTS.md", MergeMode: MergeCreateOnly},
			},
		},
		{
			ID: "codex", DisplayName: "Codex", Root: ".codex", SkillDir: ".codex/skills",
			Files: []ProviderFile{
				{TemplatePath: "providers/codex/config.toml", TargetPath: ".codex/config.toml", MergeMode: MergeReplaceManaged},
				{TemplatePath: "agents/AGENTS.md", TargetPath: ".codex/AGENTS.md", MergeMode: MergeCreateOnly},
			},
		},
		{
			ID: "claude-code", DisplayName: "Claude Code", Root: ".claude", SkillDir: ".claude/skills",
			Files: []ProviderFile{
				{TemplatePath: "providers/claude-code/.mcp.json", TargetPath: ".claude/.mcp.json", MergeMode: MergeReplaceManaged},
				{TemplatePath: "providers/claude-code/CLAUDE.md", TargetPath: ".claude/CLAUDE.md", MergeMode: MergeCreateOnly},
			},
		},
		{
			ID: "cursor", DisplayName: "Cursor", Root: ".cursor", SkillDir: ".cursor/skills",
			Files: []ProviderFile{
				{TemplatePath: "providers/cursor/mcp.json", TargetPath: ".cursor/mcp.json", MergeMode: MergeReplaceManaged},
				{TemplatePath: "providers/cursor/leaserage.mdc", TargetPath: ".cursor/rules/leaserage.mdc", MergeMode: MergeCreateOnly},
			},
		},
		{
			ID: "github-copilot", DisplayName: "GitHub Copilot", Root: ".github-copilot", SkillDir: ".github-copilot/skills",
			Files: []ProviderFile{{TemplatePath: "providers/github-copilot/copilot-instructions.md", TargetPath: ".github-copilot/copilot-instructions.md", MergeMode: MergeCreateOnly}},
		},
	}
}

func IDs() []string {
	all := All()
	ids := make([]string, 0, len(all))
	for _, provider := range all {
		ids = append(ids, provider.ID)
	}
	sort.Strings(ids)
	return ids
}

func Select(ids []string) ([]Provider, []string) {
	all := All()
	byID := map[string]Provider{}
	for _, provider := range all {
		byID[provider.ID] = provider
	}

	selected := make([]Provider, 0, len(ids))
	unknown := []string{}
	for _, id := range ids {
		provider, ok := byID[id]
		if !ok {
			unknown = append(unknown, id)
			continue
		}
		selected = append(selected, provider)
	}
	return selected, unknown
}
