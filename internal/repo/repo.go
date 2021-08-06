package repo

import "github.com/ozoncp/ocp-knowledge-api/internal/models"

type Repo interface {
	AddKnowledge(entities []models.Knowledge) error
	ListKnowledge(limit, offset uint64) ([]models.Knowledge, error)
	DescribeKnowledge(entityId uint64) (*models.Knowledge, error)
}
