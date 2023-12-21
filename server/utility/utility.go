package utility

import (
	"github.com/sraynitjsr/server/db"
	"github.com/sraynitjsr/server/model"
	"sort"
)

type SortOrder int

const (
	Ascending  SortOrder = iota
	Descending SortOrder = iota
)

func GetSortedItems(itemDB *db.ItemDB, sortOrder SortOrder) []model.ItemStruct {
	allItems := itemDB.GetAllItems()

	sortFunc := func(i, j int) bool {
		if sortOrder == Ascending {
			return allItems[i].ID < allItems[j].ID
		}
		return allItems[i].ID > allItems[j].ID
	}

	sort.Slice(allItems, sortFunc)

	return allItems
}
