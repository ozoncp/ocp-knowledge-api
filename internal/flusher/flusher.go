package flusher

import (
	"log"

	"github.com/ozoncp/ocp-knowledge-api/internal/models"
	"github.com/ozoncp/ocp-knowledge-api/internal/repo"
	"github.com/ozoncp/ocp-knowledge-api/internal/utils"
)

type Flusher interface {
	Flush(entities []models.Knowledge) []models.Knowledge
}

type flusher struct {
	chunkSize  int
	entityRepo repo.KnowledgeRepo
}

func NewFlusher(chunkSize int, repo repo.KnowledgeRepo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: repo,
	}
}

func (f flusher) Flush(knowledge []models.Knowledge) []models.Knowledge {
	chunks, err := utils.ChunkKnowledge(knowledge, f.chunkSize)
	if err != nil {
		log.Println(err)
		return knowledge
	}

	for idx := range chunks {
		if err := f.entityRepo.AddKnowledges(chunks[idx]); err != nil {
			log.Println(err)
			return knowledge[idx:]
		}
	}

	return nil
}
