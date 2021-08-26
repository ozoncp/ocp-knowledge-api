package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-knowledge-api/internal/models"
)

type KnowledgeRepo interface {
	AddKnowledges(entities []models.Knowledge) error
	AddKnowledge(entities models.Knowledge) (uint64, error)
	ListKnowledge(limit, offset uint64) ([]models.Knowledge, error)
	DescribeKnowledge(entityId uint64) (*models.Knowledge, error)
	RemoveKnowledge(entityId uint64) (bool, error)
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) KnowledgeRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) AddKnowledge(entity models.Knowledge) (uint64, error) {
	var resultId uint64

	query := "INSERT INTO knowledge(user_id, topic, text) VALUES($1, $2, $3) RETURNING id"
	err := r.db.QueryRowx(query, entity.UserId, entity.Topic, entity.Text).Scan(&resultId)

	return resultId, err
}

func (r *repo) AddKnowledges(entities []models.Knowledge) error {
	query := "INSERT INTO knowledge(user_id, topic, text) VALUES(:user_id, :topic, :text)"
	if _, err := r.db.NamedExec(query, entities); err != nil {
		return err
	}

	return nil
}

func (r *repo) ListKnowledge(limit, offset uint64) ([]models.Knowledge, error) {
	var entities []models.Knowledge
	query := "SELECT * FROM knowledge ORDER BY id LIMIT $1 OFFSET $2"
	if err := r.db.Select(&entities, query, limit, offset); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *repo) DescribeKnowledge(entityId uint64) (*models.Knowledge, error) {
	entity := models.Knowledge{}
	query := "SELECT * FROM knowledge r WHERE r.id = $1"
	err := r.db.QueryRowx(query, entityId).StructScan(&entity)

	return &entity, err
}

func (r *repo) RemoveKnowledge(entityId uint64) (bool, error) {
	query := "DELETE FROM knowledge r WHERE r.id = $1"
	result, err := r.db.Exec(query, entityId)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return count == 1, nil
}
