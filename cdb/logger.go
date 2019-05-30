package cdb

import "log"

type Logger struct {
	LogLevel  Level  `toml:"loglevel"`
	LogPrefix string `toml:"logprefix"`
	LogDir    string `toml:"logdir"`
}

const DEFAULTPREFIX = "Cdb"

func NewLogger(path string) *Logger {
	return &Logger{
		LogLevel:  DEBUG,
		LogPrefix: DEFAULTPREFIX,
	}
}

func (logger *Logger) LogDebug(arg ...interface{}) {
	log.Println(arg...)
}
