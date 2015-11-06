package main

import "./repos"
import "./ent"

var repo = repos.NewItem()

func main() {
	var i = ent.Item{
		Title: "Hello",
	}

	repo.Save(1, i)

	res := repo.Get(1)

	println(res.Value.Title)
}
