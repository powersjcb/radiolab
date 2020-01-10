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

type Config struct {
	SampleFrequency uint64
	SampleRate      float64
}

func NewStore(samples int) Store {
	return Store{
		size: samples,
		m:  sync.Mutex{},
		iq: make([]Sample, 0),
		config: Config{},
	}
}

type Store struct {
	size int
	m   sync.Mutex
	iq  []Sample
	config Config
	idx int
}

func (s *Store) GetConfig() Config {
	s.m.Lock()
	defer s.m.Unlock()
	return s.config
}

func (s *Store) SetConfig(c Config) error {
	s.m.Lock()
	defer s.m.Unlock()
	s.config = c
	// todo: validate config
	return nil
}

func (s *Store) GetIQ() []Sample {
	s.m.Lock()
	defer s.m.Unlock()
	result := s.iq
	sort.Slice(result, func(i, j int) bool {
		return result[i].OccurredAt.Before(s.iq[j].OccurredAt)
	})
	return result
}

func (s *Store) PushIQ(buf []int8) {
	s.m.Lock()
	defer s.m.Unlock()
	newSample := Sample{
		OccurredAt: time.Now(),
		Data: buf,
	}

	if len(s.iq) < s.size {
		// add to ring buffer
		s.iq = append(s.iq, newSample)
	} else {
		// rotate through ring buffer
		nextIndex := (s.idx+1) % len(s.iq)
		s.iq[nextIndex] = newSample
		s.idx = nextIndex
	}
}

func NewApplication() (Application, error) {
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
	a.Store.PushIQ(buf)
	return nil
}