package connector

import (
	"github.com/molon/go_app_base/internal/lg"
)

type Options struct {
	// basic options
	LogLevel  string `flag:"log-level"`
	LogPrefix string `flag:"log-prefix"`
	Logger    Logger
	logLevel  lg.LogLevel // private, not really an option
}

func NewOptions() *Options {
	return &Options{
		LogPrefix: "[connector] ",
		LogLevel:  "info",
	}
}
