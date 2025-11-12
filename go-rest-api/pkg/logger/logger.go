package logger

import (
    "log"
)

type Logger struct{}

func New() *Logger {
    return &Logger{}
}

func (l *Logger) Info(v ...interface{}) {
    log.Println("[INFO]", v)
}

func (l *Logger) Error(v ...interface{}) {
    log.Println("[ERROR]", v)
}
