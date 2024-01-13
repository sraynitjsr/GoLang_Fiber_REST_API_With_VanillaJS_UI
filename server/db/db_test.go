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

func (mockDB *MockItemDB) GetItemByName(name string) (model.ItemStruct, bool) {
	for _, item := range mockDB.Items {
		if item.Name == name {
			return item, true
		}
	}
	return model.ItemStruct{}, false
}

func (mockDB *MockItemDB) GetAllItems() []model.ItemStruct {
	items := make([]model.ItemStruct, 0, len(mockDB.Items))
	for _, item := range mockDB.Items {
		items = append(items, item)
	}
	return items
}

func (mockDB *MockItemDB) DeleteItemByName(name string) bool {
	for id, item := range mockDB.Items {
		if item.Name == name {
			delete(mockDB.Items, id)
			return true
		}
	}
	return false
}

func (mockDB *MockItemDB) DeleteItemByID(id int) bool {
	_, found := mockDB.Items[id]
	if found {
		delete(mockDB.Items, id)
		return true
	}
	return false
}

func TestGetItemByID(t *testing.T) {
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

func TestGetItemByName(t *testing.T) {
	mockDB := &MockItemDB{
		Items: map[int]model.ItemStruct{
			1: {ID: 1, Name: "Item 1", Message: "Message One"},
			2: {ID: 2, Name: "Item 2", Message: "Message Two"},
			3: {ID: 3, Name: "Item 3", Message: "Message Three"},
		},
	}

	// Test for an existing item by name
	expectedItem := model.ItemStruct{ID: 3, Name: "Item 3", Message: "Message Three"}
	actualItem, found := mockDB.GetItemByName("Item 3")
	if !found || actualItem != expectedItem {
		t.Errorf("Expected item %+v, but got %+v", expectedItem, actualItem)
	}

	// Test for a non-existing item by name
	nonExistentItemName := "Nonexistent Item"
	_, found = mockDB.GetItemByName(nonExistentItemName)
	if found {
		t.Errorf("Expected item with name %s to not exist, but it was found", nonExistentItemName)
	}
}

func TestGetAllItems(t *testing.T) {
	mockDB := &MockItemDB{
		Items: map[int]model.ItemStruct{
			1: {ID: 1, Name: "Item 1", Message: "Message One"},
			2: {ID: 2, Name: "Item 2", Message: "Message Two"},
			3: {ID: 3, Name: "Item 3", Message: "Message Three"},
		},
	}

	// Check if the number of items returned matches the expected count
	expectedItemCount := 3
	items := mockDB.GetAllItems()
	actualItemCount := len(items)
	if actualItemCount != expectedItemCount {
		t.Errorf("Expected %d items, but got %d", expectedItemCount, actualItemCount)
	}
}

func TestDeleteItemByName(t *testing.T) {
	mockDB := &MockItemDB{
		Items: map[int]model.ItemStruct{
			1: {ID: 1, Name: "Item 1", Message: "Message One"},
			2: {ID: 2, Name: "Item 2", Message: "Message Two"},
			3: {ID: 3, Name: "Item 3", Message: "Message Three"},
		},
	}

	// Test for deleting an existing item by name
	itemNameToDelete := "Item 2"
	if !mockDB.DeleteItemByName(itemNameToDelete) {
		t.Errorf("Failed to delete item with name %s", itemNameToDelete)
	}

	// Check if the item is actually deleted
	_, found := mockDB.GetItemByName(itemNameToDelete)
	if found {
		t.Errorf("Item with name %s still exists after deletion", itemNameToDelete)
	}

	// Test for deleting a non-existing item by name
	nonExistentItemName := "Nonexistent Item"
	if mockDB.DeleteItemByName(nonExistentItemName) {
		t.Errorf("Unexpectedly deleted non-existent item with name %s", nonExistentItemName)
	}
}

func TestDeleteItemByID(t *testing.T) {
	mockDB := &MockItemDB{
		Items: map[int]model.ItemStruct{
			1: {ID: 1, Name: "Item 1", Message: "Message One"},
			2: {ID: 2, Name: "Item 2", Message: "Message Two"},
			3: {ID: 3, Name: "Item 3", Message: "Message Three"},
		},
	}

	// Test for deleting an existing item by ID
	itemIDToDelete := 3
	if !mockDB.DeleteItemByID(itemIDToDelete) {
		t.Errorf("Failed to delete item with ID %d", itemIDToDelete)
	}

	// Check if the item is actually deleted
	_, found := mockDB.GetItemByID(itemIDToDelete)
	if found {
		t.Errorf("Item with ID %d still exists after deletion", itemIDToDelete)
	}

	// Test for deleting a non-existing item by ID
	nonExistentItemID := 100
	if mockDB.DeleteItemByID(nonExistentItemID) {
		t.Errorf("Unexpectedly deleted non-existent item with ID %d", nonExistentItemID)
	}
}
