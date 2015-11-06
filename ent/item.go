package ent

type MaybeItem struct {
	Value Item
	Err   error
}

type Item struct {
	Title string
}
