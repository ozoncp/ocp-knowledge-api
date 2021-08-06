package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-knowledge-api/internal/models"
)

func TestChunkKnowledge(t *testing.T) {
	type args struct {
		in        []models.Knowledge
		chunkSize int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]models.Knowledge
		wantErr bool
	}{
		{
			name:    "chunk size negative value",
			args:    args{chunkSize: -1},
			wantErr: true,
		},
		{
			name:    "chunk size zero value",
			args:    args{chunkSize: 0},
			wantErr: true,
		},
		{
			name:    "input slice nil",
			args:    args{chunkSize: 1},
			wantErr: true,
		},
		{
			name: "chunk size 1",
			args: args{
				chunkSize: 1,
				in: []models.Knowledge{
					{Id: 1, UserId: 2, Topic: 3, Text: "test1"},
					{Id: 4, UserId: 5, Topic: 6, Text: "test2"},
					{Id: 7, UserId: 8, Topic: 9, Text: "test3"},
				},
			},
			want: [][]models.Knowledge{
				{{Id: 1, UserId: 2, Topic: 3, Text: "test1"}},
				{{Id: 4, UserId: 5, Topic: 6, Text: "test2"}},
				{{Id: 7, UserId: 8, Topic: 9, Text: "test3"}},
			},
		},
		{
			name: "chunk size 2",
			args: args{
				chunkSize: 2,
				in: []models.Knowledge{
					{Id: 1, UserId: 2, Topic: 3, Text: "test1"},
					{Id: 2, UserId: 5, Topic: 6, Text: "test2"},
					{Id: 3, UserId: 8, Topic: 9, Text: "test3"},
					{Id: 4, UserId: 5, Topic: 6, Text: "test4"},
				},
			},
			want: [][]models.Knowledge{
				{
					{Id: 1, UserId: 2, Topic: 3, Text: "test1"},
					{Id: 2, UserId: 5, Topic: 6, Text: "test2"},
				},
				{
					{Id: 3, UserId: 8, Topic: 9, Text: "test3"},
					{Id: 4, UserId: 5, Topic: 6, Text: "test4"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChunkKnowledge(tt.args.in, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestChunkKnowledge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestChunkKnowledge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKnowledge(t *testing.T) {
	tests := []struct {
		name    string
		args    []models.Knowledge
		want    map[uint64]models.Knowledge
		wantErr bool
	}{
		{
			name:    "input slice nil",
			wantErr: true,
		},
		{
			name:    "input slice empty",
			args:    []models.Knowledge{},
			wantErr: true,
		},
		{
			name: "input slice has non unique key",
			args: []models.Knowledge{
				{Id: 1, UserId: 2, Topic: 3, Text: "test1"},
				{Id: 1, UserId: 2, Topic: 3, Text: "test1"},
			},
			wantErr: true,
		},
		{
			name: "input slice empty",
			args: []models.Knowledge{
				{Id: 1, UserId: 2, Topic: 3, Text: "test1"},
				{Id: 2, UserId: 5, Topic: 6, Text: "test2"},
			},
			want: map[uint64]models.Knowledge{
				1: {Id: 1, UserId: 2, Topic: 3, Text: "test1"},
				2: {Id: 2, UserId: 5, Topic: 6, Text: "test2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapKnowledge(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestMapKnowledge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestMapKnowledge() got = %v, want %v", got, tt.want)
			}
		})
	}
}
