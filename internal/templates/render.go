package templates

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

type SkillFile struct {
	Path string
	Perm fs.FileMode
}

func Read(path string) (string, error) {
	data, err := fs.ReadFile(FS, "files/"+strings.TrimPrefix(path, "/"))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SkillFiles() ([]SkillFile, error) {
	files := []SkillFile{}
	err := fs.WalkDir(FS, "files/skills", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		files = append(files, SkillFile{
			Path: strings.TrimPrefix(filepath.ToSlash(path), "files/"),
			Perm: info.Mode().Perm(),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})
	return files, nil
}
