package db

import "github.com/sraynitjsr/server/model"

var itemsDB map[int]model.ItemStruct

func init() {
	itemsDB = make(map[int]model.ItemStruct)
	itemsDB[1] = model.ItemStruct{ID: 1, Name: "Item 1", Message: "Message One"}
	itemsDB[2] = model.ItemStruct{ID: 2, Name: "Item 2", Message: "Message Two"}
	itemsDB[3] = model.ItemStruct{ID: 3, Name: "Item 3", Message: "Message Three"}
	itemsDB[4] = model.ItemStruct{ID: 4, Name: "Item 4", Message: "Message Four"}
	itemsDB[5] = model.ItemStruct{ID: 5, Name: "Item 5", Message: "Message Five"}
}

func GetItemByID(id int) (model.ItemStruct, bool) {
	item, found := itemsDB[id]
	return item, found
}

func GetItemByName(name string) (model.ItemStruct, bool) {
	for _, item := range itemsDB {
		if item.Name == name {
			return item, true
		}
	}
	return model.ItemStruct{}, false
}

func GetAllItems() []model.ItemStruct {
	items := make([]model.ItemStruct, 0, len(itemsDB))
	for _, item := range itemsDB {
		items = append(items, item)
	}
	return items
}
