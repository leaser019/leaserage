package install

import (
	"bytes"
	"os"
)

func RemoveFiles(plan Plan, dryRun bool) error {
	for _, op := range plan.Ops {
		if dryRun {
			continue
		}
		content, err := os.ReadFile(op.Target)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		if !bytes.Equal(content, []byte(op.Body)) {
			continue
		}
		if err := os.Remove(op.Target); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
