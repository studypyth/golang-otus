package sqlstorage

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/studypyth/hw12_13_14_15_calendar/internal/model"
)

type Storage struct {
	host, db, user, pass string
	port                 int
	pgConn               *sqlx.DB
}

func New(host, db, user, pass string, port int) (*Storage, error) {
	var s *Storage
	s = &Storage{db: db, user: user, pass: pass, port: port, host: host}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", s.user, s.pass, s.host, s.port, s.db)
	pgConn, err := sqlx.Open("pgx", dsn) // *sql.DB
	if err != nil {
		return nil, err
	}
	err = pgConn.Ping()
	if err != nil {
		return nil, err
	}
	s.pgConn = pgConn
	return s, nil
}

func (s *Storage) Close() error {
	//TODO ctx???
	err := s.pgConn.Close()
	if err != nil {
		log.Fatalf("failed to close db: %v", err)
	}
	return err
}

func (s *Storage) WriteEvent(ctx context.Context, event model.Event) error {
	startDate := event.Datetime.Format(time.RFC3339)
	startTime := event.Datetime.Format("15:04:05")
	end := event.Datetime.Add(event.Duration)
	endDate := end.Format(time.RFC3339)
	endTime := end.Format("15:04:05")
	notificationTime := event.Datetime.Add(-1 * event.NotificationTime).Format("15:04:05")
	sql := "insert into events(id, title, description, authorId,start_date,start_time,end_date,end_time,notification_time) " +
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	rows := s.pgConn.QueryRowxContext(ctx, sql, event.ID, event.Title, event.Description, event.AuthorId, startDate, startTime, endDate, endTime, notificationTime)
	if rows.Err() != nil {
		log.Fatalf("failed to writeEvent: %v", rows.Err())
	}
	return rows.Err()
}
