package db

import "github.com/sraynitjsr/server/model"

type ItemDB struct {
	Items map[int]model.ItemStruct
}

func NewItemDB() *ItemDB {
	itemsDB := make(map[int]model.ItemStruct)
	itemsDB[1] = model.ItemStruct{ID: 1, Name: "Item 1", Message: "Message One"}
	itemsDB[2] = model.ItemStruct{ID: 2, Name: "Item 2", Message: "Message Two"}
	itemsDB[3] = model.ItemStruct{ID: 3, Name: "Item 3", Message: "Message Three"}
	itemsDB[4] = model.ItemStruct{ID: 4, Name: "Item 4", Message: "Message Four"}
	itemsDB[5] = model.ItemStruct{ID: 5, Name: "Item 5", Message: "Message Five"}

	return &ItemDB{Items: itemsDB}
}

func (db *ItemDB) GetItemByID(id int) (model.ItemStruct, bool) {
	item, found := db.Items[id]
	return item, found
}

func (db *ItemDB) GetItemByName(name string) (model.ItemStruct, bool) {
	for _, item := range db.Items {
		if item.Name == name {
			return item, true
		}
	}
	return model.ItemStruct{}, false
}

func (db *ItemDB) GetAllItems() []model.ItemStruct {
	items := make([]model.ItemStruct, 0, len(db.Items))
	for _, item := range db.Items {
		items = append(items, item)
	}
	return items
}

func (db *ItemDB) DeleteItemByName(name string) bool {
	for id, item := range db.Items {
		if item.Name == name {
			delete(db.Items, id)
			return true
		}
	}
	return false
}
