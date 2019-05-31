package cdb

import (
	"log"
	"time"
)

type Logger struct {
	LogLevel  Level  `toml:"loglevel"`
	LogPrefix string `toml:"logprefix"`
	LogDir    string `toml:"logdir"`
	LogRecord chan interface{}
}

const FORMAT = "level:%s, time:%s : %s "
const DEFAULTPREFIX = "Cdb"

func NewLogger(path string) *Logger {
	return &Logger{
		LogLevel:  DEBUG,
		LogPrefix: DEFAULTPREFIX,
	}
}

func (logger *Logger) LogDebug(arg ...interface{}) {
	log.Printf(FORMAT, arg...)
}

func (logger *Logger) LogError(arg ...interface{}) {
	log.Printf(FORMAT, arg...)
}

func (logger *Logger) RecordLogData(data interface{}) {
	ln := ParseLevel(logger.LogLevel)
	if logger.LogLevel > WARN {
		logger.LogError(ln, time.Now().Format("2006-01-02 15:04:05"), data)
		return
	}
	logger.LogDebug(ln, time.Now().Format("2006-01-02 15:04:05"), data)
}

func (logger *Logger) WriteLog(data interface{}) {
	logger.LogRecord <- data
}

// init background
func (logger *Logger) loggerRecordHandle() {
	for {
		select {
		case rec, ok := <-logger.LogRecord:
			if !ok {
				logger.RecordLogData("record close")
				return
			}
			logger.RecordLogData(rec)
		}
	}
}
