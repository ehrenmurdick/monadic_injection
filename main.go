package main

import "./repos"
import "./ent"
import "./io"

func panicErr(err error) error {
	panic(err.Error())
	return err
}

// initialize a monadic repo.
// failure to open can be handled
// here by adding:
// .Handle(panicErr)
// that would quit the
// server if the db is configured
// incorrectly, for example.
// otherwise, you could defer
// handling the opening db error
// until a page request, which is more
// like what you see with Rails
// for example.
var repo = repos.
	OpenItemRepo()

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
		// handler function is of type
		// func(error) error.
		// In a web handler, you'd
		// write a handler function that
		// returns 404 for example.
		Handle(io.PrintErr)
}
