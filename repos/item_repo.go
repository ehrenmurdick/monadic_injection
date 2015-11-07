package repos

import "../ent"
import "errors"

type MaybeItemRepo struct {
	Value ItemRepo
	Err   error
}

func (in MaybeItemRepo) AndThenItem(bl func(ItemRepo) ent.MaybeItem) ent.MaybeItem {
	if in.Err != nil {
		return ent.MaybeItem{
			Err: in.Err,
		}
	} else {
		return bl(in.Value)
	}
}

func (in MaybeItemRepo) Get(key int) ent.MaybeItem {
	if in.Err != nil {
		return ent.MaybeItem{
			Err: in.Err,
		}
	} else {
		return in.Value.Get(key)
	}
}

func (in MaybeItemRepo) Save(key int, item ent.Item) ent.MaybeItem {
	if in.Err != nil {
		return ent.MaybeItem{
			Err: in.Err,
		}
	} else {
		return in.Value.Save(key, item)
	}
}

func (in MaybeItemRepo) Handle(handler func(error) error) MaybeItemRepo {
	if in.Err != nil {
		handler(in.Err)
	}
	return in
}

type ItemRepo interface {
	Get(int) ent.MaybeItem
	Save(int, ent.Item) ent.MaybeItem
}

func OpenItemRepo() (out MaybeItemRepo) {
	out.Value = NewItemRepo()
	return
}

func OpenBadItemRepo() (out MaybeItemRepo) {
	out.Err = errors.New("failed to open item repo")
	return
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
