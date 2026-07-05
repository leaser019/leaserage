package providers

type Provider struct {
	ID          string
	DisplayName string
	Root        string
	SkillDir    string
	Files       []ProviderFile
}

type ProviderFile struct {
	TemplatePath      string
	NoMCPTemplatePath string
	TargetPath        string
	MergeMode         MergeMode
	Kind              FileKind
}

type MergeMode string

const (
	MergeReplaceManaged MergeMode = "replace-managed"
	MergeCreateOnly     MergeMode = "create-only"
)

type FileKind string

const (
	FileGeneral FileKind = "general"
	FileMCP     FileKind = "mcp"
)
