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
	if types, err := data.GetItemTypesByParams(Item.ItemType{Fullname: item.ItemTypeName}); err != nil {
		return err
	} else {
		if len(types) != 1 {
			err := fmt.Errorf("Incorrect return type from a returned result. Returned %d. Must return 1", len(types))
			log.Error().Msg(err.Error())
			return err
		}

		item.ItemTypeId = types[0].TypeId
	}

	err := data.items_repo.CreateNewItem(item)

	return err
}

func (data *ItemsServiceWorker) GetItemTypesByParams(item Item.ItemType) ([]Item.ItemType, error) {
	items, err := data.items_repo.GetItemTypesByParams(item)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, fmt.Errorf("Internal database error.")
	}

	return items, nil
}

func (data *ItemsServiceWorker) GetItemsAvailability(list_id int) (int64, error) {

	items_count, err := data.items_repo.GetItemsAvailability(list_id) // TODO: refactor query (select count )
	if err != nil {
		return 0, fmt.Errorf("Error while executing in the database.")
	}

	return items_count, nil
}

func (data *ItemsServiceWorker) GetItemById(item_id int) (Item.Item, error) {
	return data.items_repo.GetItemById(item_id)
}

func InitItemsServiceWorker(items repository.ItemsService) *ItemsServiceWorker {
	return &ItemsServiceWorker{
		items_repo: items,
	}
}

func (data *ItemsServiceWorker) UpdateItemById(item *Item.Item) error {
	err := data.items_repo.UpdateItemById(item)
	if err != nil {
		log.Error().Msg(err.Error())
		return fmt.Errorf("Error while executing in the database.")
	} else {
		return nil
	}
}

func (data *ItemsServiceWorker) DeleteItemById(item_id int) error {
	err := data.items_repo.DeleteItemById(item_id)
	if err != nil {
		log.Error().Msg(err.Error())
		return fmt.Errorf("Error while executing in the database.")
	} else {
		return nil
	}
}
