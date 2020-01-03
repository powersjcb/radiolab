package pkg

import (
	"sort"
	"sync"
	"time"
)

type Sample struct {
	OccurredAt time.Time
	Data []int8
}

func NewStore(samples int) Store {
	return Store{
		size: samples,
		m:  sync.Mutex{},
		db: make([]Sample, 0),
	}
}

type Store struct {
	size int
	m   sync.Mutex
	db  []Sample
	idx int
}

func (s *Store) Get() []Sample {
	s.m.Lock()
	defer s.m.Unlock()
	result := s.db
	sort.Slice(result, func(i, j int) bool {
		return result[i].OccurredAt.Before(s.db[j].OccurredAt)
	})
	return result
}

func (s *Store) Push(buf []int8) {
	s.m.Lock()
	defer s.m.Unlock()
	newSample := Sample{
		OccurredAt: time.Now(),
		Data: buf,
	}

	if len(s.db) < s.size {
		// add to ring buffer
		s.db = append(s.db, newSample)
	} else {
		// rotate through ring buffer
		nextIndex := (s.idx+1) % len(s.db)
		s.db[nextIndex] = newSample
		s.idx = nextIndex
	}
}

func NewApplication() (Application, error) {
	// todo: extract
	// allocate some storage
	samples := 10 // samples of IQ data
	store := NewStore(samples)
	return Application{
		Store: store,
	}, nil
}

type Application struct {
	Store Store
}

func (a *Application) NoopCallback(buf []int8) error {
	a.Store.Push(buf)
	return nil
}