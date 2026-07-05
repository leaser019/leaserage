package install

import "github.com/vomkhang/leaserage/internal/providers"

type OperationKind string

const (
	OpWriteFile OperationKind = "write-file"
	OpCopySkill OperationKind = "copy-skill"
)

type Operation struct {
	Kind   OperationKind
	Target string
	Body   string
	Mode   providers.MergeMode
}

type Plan struct {
	ProviderID string
	Ops        []Operation
}

type Options struct {
	Home   string
	DryRun bool
}
