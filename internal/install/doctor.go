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
	Name string
	Kind CheckKind
	OK   bool
	Info string
}

func Doctor(home string, provider providers.Provider) []CheckResult {
	results := []CheckResult{}
	for _, file := range provider.Files {
		target := filepath.Join(home, filepath.FromSlash(file.TargetPath))
		_, err := os.Stat(target)
		info := ""
		if err != nil {
			info = err.Error()
		}
		results = append(results, CheckResult{Name: target, Kind: CheckConfig, OK: err == nil, Info: info})
	}
	for _, binary := range []string{"uvx", "codegraph", "npx"} {
		_, err := exec.LookPath(binary)
		info := ""
		if err != nil {
			info = fmt.Sprint(err)
		}
		results = append(results, CheckResult{Name: binary, Kind: CheckBinary, OK: err == nil, Info: info})
	}
	return results
}
