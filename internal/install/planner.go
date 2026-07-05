package install

import (
	"os"

	"github.com/vomkhang/leaserage/internal/providers"
)

type OperationKind string

const (
	OpWriteFile  OperationKind = "write-file"
	OpCopySkill  OperationKind = "copy-skill"
	OpRunCommand OperationKind = "run-command"
)

type Operation struct {
	Kind    OperationKind
	Target  string
	Body    string
	Mode    providers.MergeMode
	Perm    os.FileMode
	Command []string
}

type Plan struct {
	ProviderID string
	Ops        []Operation
}

type Options struct {
	Home   string
	DryRun bool
}

type MCPMode string

const (
	MCPDefault MCPMode = "default"
	MCPNone    MCPMode = "none"
)

type HookMode string

const (
	HookNone HookMode = "none"
	HookRTK  HookMode = "rtk"
)

type PlanOptions struct {
	MCPMode      MCPMode
	Hook         HookMode
	RuntimeOS    string
	RTKInstalled bool
}
