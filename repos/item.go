package repos

import "../ent"
import "errors"

type Item interface {
	Get(int) ent.MaybeItem
	Save(int, ent.Item) ent.MaybeItem
}

func NewItem() Item {
	return item{
		items: make(map[int]ent.Item),
	}
}

type item struct {
	items map[int]ent.Item
}

func (repo item) Get(key int) (out ent.MaybeItem) {
	if value, ok := repo.items[key]; ok {
		out.Value = value
	} else {
		out.Err = errors.New("Item not found")
	}
	return
}

func (repo item) Save(key int, value ent.Item) (out ent.MaybeItem) {
	out.Value = value
	repo.items[key] = value
	return
}
