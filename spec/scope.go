package spec

import (
	"github.com/karrick/godirwalk"
	"path/filepath"
	"regexp"
	"strings"
)

type Scope struct {
	Name     string
	Path     string
	Includes []string
	Excludes []string
}

func (s *Scope) hasName() bool {
	return len(s.Name) > 0
}

func (s *Scope) hasPath() bool {
	return len(s.Path) > 0
}

func (s *Scope) hasIncludes() bool {
	return len(s.Includes) > 0 || len(s.Excludes) > 0
}

func (s *Scope) hasExcludes() bool {
	return len(s.Excludes) > 0 || len(s.Includes) > 0
}

func (s *Scope) validate(index int) error {
	if !s.hasName() {
		return ErrScopeWithoutName(index)
	}

	if !s.hasPath() {
		return ErrScopeWithoutPath(index)
	}

	if !s.hasIncludes() {
		return ErrScopeWithoutIncludes(index)
	}

	if !s.hasExcludes() {
		return ErrScopeWithoutExcludes(index)
	}

	return nil
}

func (s *Scope) isValidPath(path string) bool {
	if path == s.Path {
		return true
	}

	path = strings.ToLower(path)

	for _, include := range s.Includes {
		if matches, _ := regexp.MatchString(include, path); matches {
			return true
		}
	}

	for _, exclude := range s.Excludes {
		if matches, _ := regexp.MatchString(exclude, path); matches {
			return false
		}
	}

	return len(s.Includes) == 0
}

func (s *Scope) GetFiles() ([]string, error) {
	var files []string

	options := &godirwalk.Options{
		Callback: func(path string, entry *godirwalk.Dirent) error {
			if !s.isValidPath(path) {
				return godirwalk.SkipThis
			}

			if !entry.IsDir() && filepath.Ext(path) == ".php" {
				files = append(files, path)
			}

			return nil
		},
	}

	err := godirwalk.Walk(s.Path, options)

	if err != nil {
		return nil, err
	}

	return files, nil
}
