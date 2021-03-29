package logger

import (
	"os"

	"github.com/studypyth/hw12_13_14_15_calendar/internal/config"

	"github.com/rs/zerolog"
)

type ApiLogger struct {
	zerolog.Logger
}

func New(cfg config.LoggerConf) *ApiLogger {
	var log ApiLogger
	switch {
	case cfg.File:
		f, err := os.Create(cfg.FilePath)
		if err != nil {
			panic(err)
		}
		log = ApiLogger{zerolog.New(f).Level(cfg.Level).With().Timestamp().Logger()}
	default:
		log = ApiLogger{zerolog.New(os.Stderr).Level(cfg.Level).With().Timestamp().Logger()}
	}
	zerolog.TimeFieldFormat = cfg.Timeformat
	return &log
}

func (l *ApiLogger) InfoMsg(msg string) {
	l.Info().Msg(msg)
}

func (l *ApiLogger) ErrorMsg(msg string) {
	l.Error().Msg(msg)
}
