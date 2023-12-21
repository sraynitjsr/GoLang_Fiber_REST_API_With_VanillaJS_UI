package service

import "github.com/sraynitjsr/server/model"

type ItemService interface {
	GetItemByID(id int) (model.ItemStruct, bool)
	GetItemByName(name string) (model.ItemStruct, bool)
	GetAllItems() []model.ItemStruct
	DeleteItemByID(id int) bool
	DeleteItemByName(name string) bool
}
