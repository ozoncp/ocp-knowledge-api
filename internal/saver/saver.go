package saver

import (
	"time"

	"github.com/ozoncp/ocp-knowledge-api/internal/flusher"
	"github.com/ozoncp/ocp-knowledge-api/internal/models"
)

type Saver interface {
	Save(entity models.Knowledge)
	Close()
}

type saver struct {
	flusher  flusher.Flusher
	entities []models.Knowledge
	saveCh   chan models.Knowledge
	closeCh  chan struct{}
}

func NewSaver(cap uint, fl flusher.Flusher) Saver {
	s := saver{
		flusher:  fl,
		entities: make([]models.Knowledge, 0, cap),
		saveCh:   make(chan models.Knowledge),
		closeCh:  make(chan struct{}),
	}

	s.init()

	return &s
}

func (s *saver) init() {
	ticker := time.NewTicker(time.Second * 1)

	go func() {
		defer ticker.Stop()
		defer close(s.saveCh)
		defer close(s.closeCh)

		for {
			select {
			case entity := <-s.saveCh:
				s.entities = append(s.entities, entity)
			case <-ticker.C:
				s.entities = s.flusher.Flush(s.entities)
			case <-s.closeCh:
				s.flusher.Flush(s.entities)
				return
			}
		}
	}()
}

func (s *saver) Save(entity models.Knowledge) {
	s.saveCh <- entity
}

func (s *saver) Close() {
	s.closeCh <- struct{}{}
}
