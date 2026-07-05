package providers

type Provider struct {
	ID          string
	DisplayName string
	Root        string
	SkillDir    string
	Files       []ProviderFile
}

type ProviderFile struct {
	TemplatePath string
	TargetPath   string
	MergeMode    MergeMode
}

type MergeMode string

const (
	MergeReplaceManaged MergeMode = "replace-managed"
	MergeCreateOnly     MergeMode = "create-only"
)
