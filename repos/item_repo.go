package repos

import "github.com/ehrenmurdick/monadic_injection/ent"
import "errors"

type ItemRepo interface {
	Get(string) ent.MaybeItem
	Save(string, ent.Item) ent.MaybeItem
}

func NewItemRepo() ItemRepo {
	return itemRepo{
		items: make(map[string]ent.Item),
	}
}

type itemRepo struct {
	items map[string]ent.Item
}

func (repo itemRepo) Get(key string) (out ent.MaybeItem) {
	if value, ok := repo.items[key]; ok {
		out.Value = value
	} else {
		out.Err = errors.New("Item not found")
	}
	return
}

func (repo itemRepo) Save(key string, value ent.Item) (out ent.MaybeItem) {
	out.Value = value
	repo.items[key] = value
	return
}
