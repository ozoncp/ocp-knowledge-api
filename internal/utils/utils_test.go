package utils

import (
	"reflect"
	"testing"
)

func TestBatchSlice(t *testing.T) {
	var testCases = []struct {
		caseName  string
		batchSize int
		in        []int
		out       [][]int
	}{
		{
			"Batched slice with batch size 1",
			1,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}},
		},
		{
			"Batched slice with batch size 3",
			3,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			reversedResult, err := BatchSlice(testCase.in, testCase.batchSize)
			if err != nil {
				t.Fatalf(err.Error())
			}

			if !reflect.DeepEqual(reversedResult, testCase.out) {
				t.Errorf("got %v, want %v", reversedResult, testCase.out)
			}
		})
	}
}

func TestReverseMap(t *testing.T) {
	var testCases = []struct {
		caseName string
		in       map[int]string
		out      map[string]int
	}{
		{"Empty map", map[int]string{}, map[string]int{}},
		{"Reversed slice", map[int]string{1: "v1", 2: "v2"}, map[string]int{"v1": 1, "v2": 2}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			reversedResult, err := ReverseMap(testCase.in)
			if err != nil {
				t.Fatalf(err.Error())
			}

			if !reflect.DeepEqual(reversedResult, testCase.out) {
				t.Errorf("got %v, want %v", reversedResult, testCase.out)
			}
		})
	}
}

func TestFilterSlice(t *testing.T) {
	var testCases = []struct {
		caseName string
		in       []int
		out      []int
	}{
		{"Empty slice", []int{}, []int{}},
		{"Filtered slice", []int{1, 2, 3}, []int{1, 3}},
		{"Empty result", []int{2, 5}, []int{}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			filterResult, err := FilterSlice(testCase.in)
			if err != nil {
				t.Fatalf(err.Error())
			}

			if !reflect.DeepEqual(filterResult, testCase.out) {
				t.Errorf("got %v, want %v", filterResult, testCase.out)
			}
		})
	}
}
