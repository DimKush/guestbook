package service

import (
	"sort"

	"github.com/DimKush/guestbook/tree/main/internal/entities/Item"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type ItemsServiceWorker struct {
	events_repo repository.ItemsService
}

func (data *ItemsServiceWorker) GetItemsByParams(event Item.Item) ([]Item.Item, error) {
	items, err := data.events_repo.GetItemsByParams(event)

	if err != nil {
		return nil, err
	}

	if len(items) != 0 {
		sort.Slice(items, func(i int, j int) bool {
			return items[i].Id < items[j].Id
		})
	} else {
		items = []Item.Item{}
	}

	return items, nil
}

func InitItemsServiceWorker(items repository.ItemsService) *ItemsServiceWorker {
	return &ItemsServiceWorker{
		events_repo: items,
	}
}
