package utility

import (
	"github.com/sraynitjsr/server/db"
	"github.com/sraynitjsr/server/model"
	"sort"
)

func GetSortedItems(itemDB *db.ItemDB) []model.ItemStruct {
	allItems := itemDB.GetAllItems()
  
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].ID < allItems[j].ID
	})
  
	return allItems
}
