package repos

import "errors"
import "../ent"

type MaybeItemRepo struct {
	Value ItemRepo
	Err   error
}

func OpenBadItemRepo() (out MaybeItemRepo) {
	out.Err = errors.New("failed to open item repo")
	return
}

func OpenItemRepo() (out MaybeItemRepo) {
	out.Value = NewItemRepo()
	return
}

func (in MaybeItemRepo) Get(key string) ent.MaybeItem {
	return in.AndThenItem(func(repo ItemRepo) ent.MaybeItem {
		return repo.Get(key)
	})
}

func (in MaybeItemRepo) Save(key string, item ent.Item) ent.MaybeItem {
	return in.AndThenItem(func(repo ItemRepo) ent.MaybeItem {
		return repo.Save(key, item)
	})
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

func (in MaybeItemRepo) Handle(handler func(error) error) MaybeItemRepo {
	if in.Err != nil {
		handler(in.Err)
	}
	return in
}
