package todo_test

import (
	"context"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		{
			name:           "Given a Todo of a shop and a search of sh, I should get shop back.",
			toDosToAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "Still returns shop although the case doesn't match",
			toDosToAdd:     []string{"Shopping"},
			query:          "sh",
			expectedResult: []string{"Shopping"},
		},
		{
			name:           "Spaces at the start of the word",
			toDosToAdd:     []string{"Go shopping"},
			query:          "go",
			expectedResult: []string{"Go shopping"},
		},
		{
			name:           "Spaces at the end of the word",
			toDosToAdd:     []string{" space at the beginning"},
			query:          "space",
			expectedResult: []string{" space at the beginning"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Fatal(err)
				}
			}
			got, err := svc.Search(tt.query)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}