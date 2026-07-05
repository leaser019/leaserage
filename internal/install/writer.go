package install

import (
	"os"
	"path/filepath"
	"time"
)

func WriteManagedFile(path string, body string, dryRun bool) error {
	if dryRun {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	if _, err := os.Stat(path); err == nil {
		backup := path + ".leaserage-backup-" + time.Now().UTC().Format("20060102T150405Z")
		if err := copyFile(path, backup); err != nil {
			return err
		}
	}
	return os.WriteFile(path, []byte(body), 0o644)
}

func copyFile(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}
