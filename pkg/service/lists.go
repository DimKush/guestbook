package service

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type ListsServiceWorker struct {
	db_lists repository.ListService
	db_users repository.UsersService
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

func (data *ListsServiceWorker) GetAutoListId() (int, error) {
	id, err := data.db_lists.GetAutoListId()

	if err != nil {
		return 0, err
	}

	id++ // nextval

	return id, nil
}

func (data *ListsServiceWorker) CreateList(newList List.List) error {
	// check required fields
	if newList.Title == "" {
		return fmt.Errorf("The Field : Title cannot be empty.")
	}

	if len(newList.Title) > 255 {
		return fmt.Errorf("The Field : Title cannot be more that 255 symbols.")
	}

	if newList.Owner == "" {
		return fmt.Errorf("Field Owner cannot be empty.")
	}

	// get owner's id
	user, err := data.db_users.GetUserByUsername(newList.Owner)

	if err != nil {
		return err
	}
	if (user == User.User{}) {
		return fmt.Errorf("The username %s doesn't exists.", newList.Owner)
	}

	newList.OwnerUserId = user.Id

	// check if list with list_id exists
	list_id := newList.Id

	list, _ := data.GetListById(list_id)

	if (list != List.List{}) {
		err := fmt.Errorf("Cannot create a new list.The list with id = %d already exists.", list_id)
		log.Error().Msgf(err.Error())
		return err
	}

	err = data.db_lists.CreateList(newList)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (data *ListsServiceWorker) GetListById(list_id int) (List.List, error) {
	return data.db_lists.GetListById(list_id)
}

func (data *ListsServiceWorker) DeleteListById(list_id int) error {

	audit_ch := make(chan error)

	go func(list_id int) {
		audit_ch <- Audit.WriteEventParams("ListsServiceWorker",
			"DeleteListById",
			AUDIT_INFO,
			time.Now(),
			false,
			fmt.Sprintf("Delete list with id =  %d", list_id),
		)
	}(list_id)

	list, err := data.GetListById(list_id)
	if err != nil {
		return err
	}

	if (list == List.List{}) {
		return fmt.Errorf("List with id %d doesn't exist.", list_id)
	}

	err = data.db_lists.DeleteListById(list_id)
	if err != nil {
		return err
	}

	return nil
}

func InitListsServiceWorker(repos repository.ListService, repos_users repository.UsersService) *ListsServiceWorker {
	return &ListsServiceWorker{db_lists: repos, db_users: repos_users}
}
