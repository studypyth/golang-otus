package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/studypyth/hw12_13_14_15_calendar/internal/model"
)

type Storage struct {
	IdEventMap   map[string]model.Event
	DateEventMap map[time.Time]model.Event
	Mu           *sync.RWMutex
}

func New() *Storage {
	mu := sync.RWMutex{}
	idMap := map[string]model.Event{}
	dateMap := map[time.Time]model.Event{}
	return &Storage{IdEventMap: idMap, DateEventMap: dateMap, Mu: &mu}
}

func (s *Storage) WriteEvent(ctx context.Context, event model.Event) error {
	s.Mu.Lock()
	s.IdEventMap[event.ID] = event
	s.DateEventMap[event.Datetime] = event
	s.Mu.Unlock()
	return nil
}
