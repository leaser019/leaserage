package templates

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

func Read(path string) (string, error) {
	data, err := fs.ReadFile(FS, "files/"+strings.TrimPrefix(path, "/"))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SkillFiles() ([]string, error) {
	entries, err := fs.ReadDir(FS, "files/skills")
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			files = append(files, filepath.ToSlash(filepath.Join("skills", entry.Name(), "SKILL.md")))
		}
	}
	sort.Strings(files)
	return files, nil
}
