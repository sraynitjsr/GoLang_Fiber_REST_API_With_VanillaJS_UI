package db

import (
	"testing"
	"github.com/sraynitjsr/server/model"
)

// MockItemDB is a mock implementation of the ItemDB type
type MockItemDB struct {
	Items map[int]model.ItemStruct
}

func (mockDB *MockItemDB) GetItemByID(id int) (model.ItemStruct, bool) {
	item, found := mockDB.Items[id]
	return item, found
}

func TestGetItemByID(t *testing.T) {
	// Create a mock database with sample data
	mockDB := &MockItemDB{
		Items: map[int]model.ItemStruct{
			1: {ID: 1, Name: "Item 1", Message: "Message One"},
			2: {ID: 2, Name: "Item 2", Message: "Message Two"},
			3: {ID: 3, Name: "Item 3", Message: "Message Three"},
		},
	}

	// Test for an existing item
	expectedItem := model.ItemStruct{ID: 2, Name: "Item 2", Message: "Message Two"}
	actualItem, found := mockDB.GetItemByID(2)
	if !found || actualItem != expectedItem {
		t.Errorf("Expected item %+v, but got %+v", expectedItem, actualItem)
	}

	// Test for a non-existing item
	nonExistentItemID := 100
	_, found = mockDB.GetItemByID(nonExistentItemID)
	if found {
		t.Errorf("Expected item with ID %d to not exist, but it was found", nonExistentItemID)
	}
}
