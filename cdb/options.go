package cdb

import (
	"github.com/BurntSushi/toml"
)

type Options struct {
	Logger Logger
}

func NewOptions(path string) (*Options, error) {
	op := Options{}
	_, err := toml.DecodeFile(path, &op)
	if err != nil {
		return nil, err
	}

	return &op, nil
}
