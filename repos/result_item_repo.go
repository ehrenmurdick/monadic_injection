package repos

import . "github.com/ehrenmurdick/monadic_injection/ent"

func (m ResultItemRepo) Get(key string) ResultItem {
	return m.FmapItem(func(i ItemRepo) (Item, error) {
		return i.Get(key)
	})
}

func (m ResultItemRepo) Save(key string, item Item) ResultItem {
	return m.FmapItem(func(i ItemRepo) (Item, error) {
		return i.Save(key, item)
	})
}

//go:generate ../monads/result result_item_repo.go ItemRepo
//go:generate ../monads/fmap result_item_repo.go ItemRepo Item
// GENERATED
type ResultItemRepo struct {
	Value ItemRepo
	Err   error
}

func ReturnResultItemRepo(v ItemRepo) ResultItemRepo {
	return ResultItemRepo{
		Value: v,
	}
}

func (m ResultItemRepo) Bind(bl func(ItemRepo) ItemRepo) ResultItemRepo {
	if m.Err != nil {
		return ResultItemRepo{
			Err: m.Err,
		}
	} else {
		return ResultItemRepo{
			Value: bl(m.Value),
		}
	}
}

func (m ResultItemRepo) Handle(bl func(error) error) ResultItemRepo {
	if m.Err != nil {
		return ResultItemRepo{
			Err: bl(m.Err),
		}
	} else {
		return m
	}
}

func (m ResultItemRepo) FmapItem(bl func(ItemRepo) (Item, error)) (out ResultItem) {
	if m.Err != nil {
		out.Err = m.Err
	} else {
		out.Value, out.Err = bl(m.Value)
	}
	return
}
