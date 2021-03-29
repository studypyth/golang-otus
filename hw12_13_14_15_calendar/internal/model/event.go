package model

import "time"

type Event struct {
	ID               string    `db:"ID"`
	Title            string    `db:"Title"`
	Description      string    `db:"Description"`
	AuthorId         string    `db:"AuthorId"`
	Datetime         time.Time `db:"Datetime"`
	Duration         time.Duration
	NotificationTime time.Duration
}
