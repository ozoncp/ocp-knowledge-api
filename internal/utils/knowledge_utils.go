package utils

import (
	"errors"

	"github.com/ozoncp/ocp-knowledge-api/internal/models"
)

// ChunkKnowledge splits input slice by passed chunkSize.
func ChunkKnowledge(in []models.Knowledge, chunkSize int) ([][]models.Knowledge, error) {
	inSize := len(in)

	if inSize == 0 {
		return nil, errors.New("input slice is nil or empty")
	}
	if chunkSize <= 0 {
		return nil, errors.New("chunk size less or equal zero")
	}

	chunkCount := inSize / chunkSize
	if inSize%chunkSize > 0 {
		chunkCount++
	}

	chunkedSlice := make([][]models.Knowledge, 0, chunkCount)

	for i := 0; i < inSize; i += chunkSize {
		end := i + chunkSize

		if end > inSize {
			end = inSize
		}

		chunk := append([]models.Knowledge{}, in[i:end]...)
		chunkedSlice = append(chunkedSlice, chunk)
	}

	return chunkedSlice, nil
}

// MapKnowledge converts slice of Knowledge to map of Knowledge where Knowledge.Id is key of map.
func MapKnowledge(in []models.Knowledge) (map[uint64]models.Knowledge, error) {
	if len(in) == 0 {
		return nil, errors.New("input slice is empty or nil")
	}

	result := make(map[uint64]models.Knowledge, len(in))

	for _, v := range in {
		if _, ok := result[v.Id]; ok {
			return nil, errors.New("key exists in map")
		}

		result[v.Id] = v
	}

	return result, nil
}
