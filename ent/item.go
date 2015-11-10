package ent

type Item struct {
	Title string
}

func (i Item) GetTitle() (string, error) {
	return i.Title, nil
}
