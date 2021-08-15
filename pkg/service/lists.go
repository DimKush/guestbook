package service

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type ListsServiceWorker struct {
	db_lists repository.ListService
}

func (data *ListsServiceWorker) GetAllLists() ([]List.List, error) {
	log.Info().Msg("GetAllLists process request")

	lists, err := data.db_lists.GetAllLists()
	if err != nil {
		log.Error().Msg(err.Error())
		return []List.List{}, err
	}

	for _, val := range lists {
		fmt.Println(val)
	}

	return lists, nil
}

func InitListsServiceWorker(repos repository.ListService) *ListsServiceWorker {
	return &ListsServiceWorker{db_lists: repos}
}
