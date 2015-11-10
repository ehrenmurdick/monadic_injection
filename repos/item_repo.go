package repos

import "github.com/ehrenmurdick/monadic_injection/ent"
import "errors"

type ItemRepo interface {
	Get(string) (ent.Item, error)
	Save(string, ent.Item) (ent.Item, error)
}

func NewItemRepo() ItemRepo {
	return itemRepo{
		items: make(map[string]ent.Item),
	}
}

type itemRepo struct {
	items map[string]ent.Item
}

func (repo itemRepo) Get(key string) (out ent.Item, err error) {
	if value, ok := repo.items[key]; ok {
		out = value
	} else {
		err = errors.New("Item not found")
	}
	return
}

func (repo itemRepo) Save(key string, value ent.Item) (out ent.Item, err error) {
	out = value
	repo.items[key] = value
	return
}
