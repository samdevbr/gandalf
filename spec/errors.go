package spec

import (
	"errors"
	"fmt"
)

var ErrSpecWithoutScopes = errors.New("specification without scopes")
var ErrUnsupportedPhpVersion = errors.New("unsupported Php version")
var ErrSpecWithoutPhpSettings = errors.New("specification without Php settings")

var ErrScopeWithoutName = func(index int) error {
	return errors.New(fmt.Sprintf("scope does not have a name: index %d", index))
}

var ErrScopeWithoutPath = func(index int) error {
	return errors.New(fmt.Sprintf("scope does not have a path: index %d", index))
}

var ErrScopeWithoutIncludes = func(index int) error {
	return errors.New(fmt.Sprintf("scope does not have includes: index %d", index))
}

var ErrScopeWithoutExcludes = func(index int) error {
	return errors.New(fmt.Sprintf("scope does not have excludes: index %d", index))
}
