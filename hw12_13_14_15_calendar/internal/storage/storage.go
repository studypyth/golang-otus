package storage

import (
	"errors"
	"strings"

	"github.com/studypyth/hw12_13_14_15_calendar/internal/app"
	"github.com/studypyth/hw12_13_14_15_calendar/internal/config"
	memory "github.com/studypyth/hw12_13_14_15_calendar/internal/storage/memory"
	db "github.com/studypyth/hw12_13_14_15_calendar/internal/storage/sql"
)

func New(conf config.StorageConf) (app.Storage, error) {
	switch strings.ToLower(conf.Type) {
	case "db":
		return db.New(conf.Db.Host, conf.Db.Db, conf.Db.User, conf.Db.Pass, conf.Db.Port)
	case "memory":
		return memory.New(), nil
	default:
		return nil, errors.New("inccorect type Storage or parametr DB")
	}

}
