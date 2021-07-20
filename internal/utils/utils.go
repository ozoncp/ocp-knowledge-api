package utils

import (
	"errors"
)

// BatchSlice batch input slice by passed batchSize.
func BatchSlice(in []int, batchSize int) ([][]int, error) {
	if batchSize <= 0 {
		return nil, errors.New("batch size less or equal zero")
	}

	inSize := len(in)
	if inSize <= batchSize {
		batchedSlice := append([]int{}, in...)
		return [][]int{batchedSlice}, nil
	}

	batchCount := inSize / batchSize
	if inSize%batchSize > 0 {
		batchCount++
	}

	batchedSlice := make([][]int, 0, batchCount)

	for i := 0; i < inSize; i += batchSize {
		end := i + batchSize

		if end > inSize {
			end = inSize
		}

		batch := append([]int{}, in[i:end]...)
		batchedSlice = append(batchedSlice, batch)
	}

	return batchedSlice, nil
}

// ReverseMap reverse input map, (key-value) project to (value-key).
func ReverseMap(in map[int]string) (map[string]int, error) {
	if in == nil {
		return nil, errors.New("input map is nil")
	}

	reversedMap := make(map[string]int, len(in))

	for key, value := range in {
		reversedMap[value] = key
	}

	return reversedMap, nil
}

// FilterSlice filter input slice by hardcoded list.
func FilterSlice(in []int) ([]int, error) {
	if in == nil {
		return nil, errors.New("input slice is nil")
	}

	filterList := map[int]bool{
		1: true,
		3: true,
	}

	inLen := len(in)
	filteredSlice := make([]int, 0, inLen)

	for _, value := range in {
		if _, ok := filterList[value]; ok {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return filteredSlice, nil
}
