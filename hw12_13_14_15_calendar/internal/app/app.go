package app

import (
	"context"
	"time"

	"github.com/studypyth/hw12_13_14_15_calendar/internal/model"
)

type App struct {
	Logger  Logger
	Storage Storage
}

type Logger interface {
	ErrorMsg(msg string)
	InfoMsg(msg string)
}

type Storage interface {
	WriteEvent(ctx context.Context, event model.Event) error
}

func New(logger Logger, storage Storage) *App {
	return &App{Logger: logger, Storage: storage}
}

func (a *App) CreateEvent(
	ctx context.Context,
	id, title, description, authorId string,
	datetime time.Time,
	duration, notificationTime time.Duration,
) error {
	// TODO context???
	event := model.Event{
		ID:               id,
		Title:            title,
		Description:      description,
		AuthorId:         authorId,
		Datetime:         datetime,
		Duration:         duration,
		NotificationTime: notificationTime,
	}
	err := a.Storage.WriteEvent(ctx, event)
	return err
}
