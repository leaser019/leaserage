package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vomkhang/leaserage/internal/providers"
)

type CheckKind string

const (
	CheckConfig CheckKind = "config"
	CheckBinary CheckKind = "binary"
)

type CheckResult struct {
	Name     string
	Kind     CheckKind
	OK       bool
	Info     string
	Required bool
}

func Doctor(home string, provider providers.Provider, opts PlanOptions) []CheckResult {
	if opts.MCPMode == "" {
		opts.MCPMode = MCPDefault
	}
	if opts.Hook == "" {
		opts.Hook = HookNone
	}

	results := []CheckResult{}
	for _, file := range provider.Files {
		if opts.MCPMode == MCPNone && file.Kind == providers.FileMCP {
			if file.NoMCPTemplatePath == "" {
				continue
			}
		}
		target := filepath.Join(home, filepath.FromSlash(file.TargetPath))
		_, err := os.Stat(target)
		info := ""
		if err != nil {
			info = err.Error()
		}
		results = append(results, CheckResult{Name: target, Kind: CheckConfig, OK: err == nil, Info: info, Required: true})
	}
	if opts.MCPMode == MCPDefault {
		for _, binary := range []string{"uvx", "codegraph", "npx"} {
			results = append(results, binaryCheck(binary, true))
		}
	}
	if opts.Hook == HookRTK {
		results = append(results, binaryCheck("rtk", true))
	}
	return results
}

func binaryCheck(binary string, required bool) CheckResult {
	_, err := exec.LookPath(binary)
	info := ""
	if err != nil {
		info = fmt.Sprint(err)
	}
	return CheckResult{Name: binary, Kind: CheckBinary, OK: err == nil, Info: info, Required: required}
}
