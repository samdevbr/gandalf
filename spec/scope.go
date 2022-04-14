package spec

type scope struct {
	Name     string
	Path     string
	Includes []string
	Excludes []string
}

func (s *scope) hasName() bool {
	return len(s.Name) > 0
}

func (s *scope) hasPath() bool {
	return len(s.Path) > 0
}

func (s *scope) hasIncludes(phpSettings *php) bool {
	return len(s.Includes) > 0 || phpSettings.Framework != nil
}

func (s *scope) hasExcludes(phpSettings *php) bool {
	return len(s.Excludes) > 0 || phpSettings.Framework != nil
}

func (s *scope) validate(index int, phpSettings *php) error {
	if !s.hasName() {
		return ErrScopeWithoutName(index)
	}

	if !s.hasPath() {
		return ErrScopeWithoutPath(index)
	}

	if !s.hasIncludes(phpSettings) {
		return ErrScopeWithoutIncludes(index)
	}

	if !s.hasExcludes(phpSettings) {
		return ErrScopeWithoutExcludes(index)
	}

	return nil
}
