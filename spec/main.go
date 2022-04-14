package spec

import (
	"github.com/BurntSushi/toml"
)

type specification struct {
	PHP    *php
	Scopes []scope `toml:"scope"`
}

func (spec *specification) hasPhpSettings() bool {
	return spec.PHP != nil
}

func (spec *specification) hasFramework() bool {
	return spec.PHP.Framework != nil
}

func (spec *specification) hasScopes() bool {
	return len(spec.Scopes) > 0
}

func (spec *specification) validateScopes() error {
	for i, scope := range spec.Scopes {
		err := scope.validate(i, spec.PHP)

		if err != nil {
			return err
		}
	}

	return nil
}

func (spec *specification) validate() error {
	if !spec.hasPhpSettings() {
		return ErrSpecWithoutPhpSettings
	}

	if !spec.hasScopes() {
		return ErrSpecWithoutScopes
	}

	if !spec.PHP.hasValidVersion() {
		return ErrUnsupportedPhpVersion
	}

	scopesErr := spec.validateScopes()

	if scopesErr != nil {
		return scopesErr
	}

	return nil
}

func New(specificationFile string) (*specification, error) {
	var spec specification

	_, err := toml.DecodeFile(specificationFile, &spec)

	if err != nil {
		return nil, err
	}

	return &spec, spec.validate()
}
