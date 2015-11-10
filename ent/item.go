package ent

//go:generate ../monads/generate Item > result_item.go

type Item struct {
	Title string
}

// getters
func (i Item) GetTitle() string {
	return i.Title
}
