package service

import (
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

	return lists, nil
}

func (data *ListsServiceWorker) GetListsByParams(list List.List) ([]List.List, error) {
	log.Info().Msg("GetListsByParams process request")

	lists, err := data.db_lists.GetListsByParams(list)
	if err != nil {
		log.Error().Msg(err.Error())
		return []List.List{}, err
	}

	return lists, nil
}

func InitListsServiceWorker(repos repository.ListService) *ListsServiceWorker {
	return &ListsServiceWorker{db_lists: repos}
}
