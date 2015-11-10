package repos

import "errors"
import "github.com/ehrenmurdick/monadic_injection/ent"

// This is the scary basement.
// If we were in swift, all the Maybe* classes
// Could be covered by one implementation using
// generics. Alas!
// At least the implementation of any Maybe* is trivial.
// Could probably write a code generator for monads, ala
// counterfeiter.

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

// I didn't implement any monad identity functions, as
// I'm always going from a repo to an item.
// A MaybeItemRepo -> MaybeItemRepo functor would
// be trivial though, see ent/string.go

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
