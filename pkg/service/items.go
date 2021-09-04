package service

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/entities/Item"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
	"github.com/rs/zerolog/log"
)

type ItemsServiceWorker struct {
	items_repo repository.ItemsService
}

func (data *ItemsServiceWorker) GetItemsByParams(event Item.Item) ([]Item.Item, error) {
	items, err := data.items_repo.GetItemsByParams(event)

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

func (data *ItemsServiceWorker) CreateNewItem(item Item.Item) error {
	audit_ch := make(chan error)

	go func() {
		out, err := json.Marshal(item)

		if err != nil {
			log.Error().Msg(err.Error())
			audit_ch <- err
			return
		}

		audit_ch <- Audit.WriteEventParams("ItemsServiceWorker",
			"CreateNewItem",
			AUDIT_INFO,
			time.Now(),
			false,
			fmt.Sprintf("Create new item %s", string(out)),
		)
	}()
	// get item_type by the name
	itemTypeName := 

	err := data.items_repo.CreateNewItem(item)

	return err
}

func InitItemsServiceWorker(items repository.ItemsService) *ItemsServiceWorker {
	return &ItemsServiceWorker{
		items_repo: items,
	}
}
