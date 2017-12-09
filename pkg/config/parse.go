package config

import (
	"github.com/BurntSushi/toml"
)

func Parse(filename string) (*toml.MetaData, *File, error) {
	f := &File{}
	md, err := toml.DecodeFile(filename, f)
	if err != nil {
		return nil, nil, err
	}
	return &md, f, nil
}
