package main

import "./repos"
import "./ent"
import "./io"

var repo = repos.NewItem()

func main() {
	var i = ent.Item{
		Title: "Hello",
	}

	repo.Save(1, i)

	// Get item 1 and print out title
	// ignores errors
	repo.
		Get(1).
		GetTitle().
		Within(io.PrintStr)

	// Get item 100 and print out title
	// print out any errors from anywhere
	repo.
		Get(100).
		GetTitle().
		Within(io.PrintStr).
		Handle(io.PrintErr)
}
