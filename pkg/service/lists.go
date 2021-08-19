package service

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

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
	log.Info().Msg("Service GetListsByParams process request")

	//write event into audit_event
	audit_ch := make(chan error)

	go func() {
		out, err := json.Marshal(list)
		if err != nil {
			log.Error().Msg(err.Error())
		}
		audit_ch <- Audit.WriteEventParams("ListsServiceWorker",
			"GetListsByParams",
			AUDIT_INFO,
			time.Now(),
			false,
			fmt.Sprintf("Get lists params %s", string(out)),
		)
	}()

	lists, err := data.db_lists.GetListsByParams(list)
	if err != nil {
		log.Error().Msg(err.Error())
		return []List.List{}, err
	}

	// sorting output
	if len(lists) != 0 {
		sort.Slice(lists, func(i, j int) bool {
			return lists[i].Id < lists[j].Id
		})
	}

	//output log
	log.Info().Msg("Database output:")

	for _, val := range lists {
		log.Info().Msgf("\n%v", val)
	}

	if err := <-audit_ch; err != nil {
		log.Error().Msg(err.Error())
	}

	return lists, nil
}

func InitListsServiceWorker(repos repository.ListService) *ListsServiceWorker {
	return &ListsServiceWorker{db_lists: repos}
}
