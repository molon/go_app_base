package connector

type Options struct {
	// basic options
	LogLevel string `flag:"log-level"`
}

func NewOptions() *Options {
	return &Options{
		LogLevel: "info",
	}
}
