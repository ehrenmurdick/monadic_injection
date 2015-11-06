package ent

type MaybeItem struct {
	Value Item
	Err   error
}

type Item struct {
	Title string
}

// getters
func (i Item) GetTitle() string {
	return i.Title
}

// side effects
func (in MaybeItem) Handle(handler func(error) error) MaybeItem {
	if in.Err != nil {
		handler(in.Err)
	}
	return in
}

// functors
func (in MaybeItem) AndThenString(bl func(Item) string) (out MaybeString) {
	if in.Err == nil {
		out.Value = bl(in.Value)
	}
	out.Err = in.Err
	return
}

func (in MaybeItem) GetTitle() (out MaybeString) {
	return in.AndThenString(Item.GetTitle)
}
