package logging

// io.MultiWriter auf loggile, os.stdout, discord channel

import (
	"fmt"
	"log"
)

const (
	Fatal = 1
	Warn  = 2
	Info  = 3
	Debug = 4
)

// log Level
var LogLevel = Debug

var LevelToSymbol = map[int]string{
	Fatal: "F",
	Warn:  "W",
	Info:  "I",
	Debug: "D",
}

func SetLogLevel(level int) {
	LogLevel = level
}

func Logf(level int, format string, args ...interface{}) {
	if level <= LogLevel {
		symbol := LevelToSymbol[level]
		if symbol == "" {
			symbol = "?"
		}
		format = fmt.Sprintf("[%s] %s", symbol, format)
		if level > 1 {
			log.Printf(format, args...)
		} else {
			log.Panicf(format, args...)
		}
	}
}
