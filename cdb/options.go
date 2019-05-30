package cdb

type Options struct {
	Logger *Logger
}

func NewOptions() *Options {
	return &Options{
		Logger: NewLogger(),
	}
}
