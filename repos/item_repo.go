package repos

import "../ent"
import "errors"

type ItemRepo interface {
	Get(int) ent.MaybeItem
	Save(int, ent.Item) ent.MaybeItem
}

func NewItemRepo() ItemRepo {
	return itemRepo{
		items: make(map[int]ent.Item),
	}
}

type itemRepo struct {
	items map[int]ent.Item
}

func (repo itemRepo) Get(key int) (out ent.MaybeItem) {
	if value, ok := repo.items[key]; ok {
		out.Value = value
	} else {
		out.Err = errors.New("Item not found")
	}
	return
}

func (repo itemRepo) Save(key int, value ent.Item) (out ent.MaybeItem) {
	out.Value = value
	repo.items[key] = value
	return
}
