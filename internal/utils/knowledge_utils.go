package utils

import (
	"errors"

	"github.com/ozoncp/ocp-knowledge-api/internal/models"
)

// Batch splits input slice by passed batchSize.
func BatchKnowledge(in []models.Knowledge, batchSize int) ([][]models.Knowledge, error) {
	if in == nil {
		return nil, errors.New("input slice is nil")
	}

	if batchSize <= 0 {
		return nil, errors.New("batch size less or equal zero")
	}

	inSize := len(in)
	batchCount := inSize / batchSize
	if inSize%batchSize > 0 {
		batchCount++
	}

	batchedSlice := make([][]models.Knowledge, 0, batchCount)

	for i := 0; i < inSize; i += batchSize {
		end := i + batchSize

		if end > inSize {
			end = inSize
		}

		batch := append([]models.Knowledge{}, in[i:end]...)
		batchedSlice = append(batchedSlice, batch)
	}

	return batchedSlice, nil
}

// MapKnowledge converts slice of Knowledge to map of Knowledge where Knowledge.Id is key of map.
func MapKnowledge(in []models.Knowledge) (map[uint64]models.Knowledge, error) {
	if len(in) == 0 {
		return nil, errors.New("input slice is empty or nil")
	}

	result := make(map[uint64]models.Knowledge, len(in))

	for _, v := range in {
		result[v.Id] = v
	}

	return result, nil
}
