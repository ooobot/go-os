package os

/*
	OS source uses the Micro config-srv
*/

import (
	"github.com/ooobot/go-os/config"
)

func NewSource(opts ...config.SourceOption) config.Source {
	return config.NewSource(opts...)
}
