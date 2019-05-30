package cdb

import "log"

type Logger struct {
	LogLevel  Level
	LogPrefix string
}

const DEFAULTPREFIX = "Cdb"

func NewLogger() *Logger {
	return &Logger{
		LogLevel:  DEBUG,
		LogPrefix: DEFAULTPREFIX,
	}
}

func (logger *Logger) LogDebug(arg ...interface{}) {
	log.Println(arg...)
}
