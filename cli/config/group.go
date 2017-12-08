package config

import (
	"fmt"
)

// Group represents a custom group
type Group struct {
	Sequence []GroupSequence `yaml:"sequence"`
}

type GroupSequence interface{}

func (g GroupSequence) IsValid() bool {
	switch t := g.(type) {
	case string:
		return true
	case []string:
		return true
	default:
		return false
	}
}

func (g GroupSequence) JobNames() ([]string, error) {
	switch t := g.(type) {
	case string:
		return []string{t}, nil
	case []string:
		return t, nil
	default:
		return nil, fmt.Errorf("invalid type %v", t)
	}
}
