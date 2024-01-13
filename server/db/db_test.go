package db

import (
	"testing"
	"github.com/sraynitjsr/server/model"
)

func TestGetItemByID(t *testing.T) {
	db := NewItemDB()

	// Test for an existing item
	expectedItem := model.ItemStruct{ID: 2, Name: "Item 2", Message: "Message Two"}
	actualItem, found := db.GetItemByID(2)
	if !found || actualItem != expectedItem {
		t.Errorf("Expected item %+v, but got %+v", expectedItem, actualItem)
	}

	// Test for a non-existing item
	nonExistentItemID := 100
	_, found = db.GetItemByID(nonExistentItemID)
	if found {
		t.Errorf("Expected item with ID %d to not exist, but it was found", nonExistentItemID)
	}
}
