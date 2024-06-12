package common

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() {
	log.Logger = zerolog.New(&lumberjack.Logger{
		Filename:   "logs/app.json",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     30, // days
	}).With().Timestamp().Logger()
}
