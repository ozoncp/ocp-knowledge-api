package repo

import "github.com/ozoncp/ocp-knowledge-api/internal/models"

type Repo interface {
	AddKnowledgies(entities []models.Knowledge) error
	ListKnowledgies(limit, offset uint64) ([]models.Knowledge, error)
	DescribeKnowledge(entityId uint64) (*models.Knowledge, error)
}
