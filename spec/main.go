package spec

import (
	"github.com/BurntSushi/toml"
)

type Specification struct {
	PHP    *Php
	Scopes []Scope `toml:"Scope"`
}

func (spec *Specification) hasPhpSettings() bool {
	return spec.PHP != nil
}

func (spec *Specification) hasScopes() bool {
	return len(spec.Scopes) > 0
}

func (spec *Specification) validateScopes() error {
	for i, scope := range spec.Scopes {
		err := scope.validate(i)

		if err != nil {
			return err
		}
	}

	return nil
}

func (spec *Specification) validate() error {
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

func New(specificationFile string) (*Specification, error) {
	var spec Specification

	_, err := toml.DecodeFile(specificationFile, &spec)

	if err != nil {
		return nil, err
	}

	return &spec, spec.validate()
}
