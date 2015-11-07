package main

import (
	"fmt"
	"net/http"
	"strconv"

	"./ent"
	"./repos"
)

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
// for example, by simply deleting the
// handle line here.
var repo = repos.
	OpenItemRepo().
	Handle(panicErr)

func writeResponse(w http.ResponseWriter) func(string) error {
	return func(str string) error {
		_, err := fmt.Fprintf(w, str)
		return err
	}
}

func writeError(w http.ResponseWriter) func(error) error {
	return func(err error) error {
		fmt.Fprintf(w, err.Error())
		return err
	}
}

func show(w http.ResponseWriter, r *http.Request) {
	var key string
	key = r.URL.Path[1:]

	repo.
		Get(key).
		GetTitle().
		Within(writeResponse(w)).
		Handle(writeError(w))
}

func main() {
	for x := 0; x < 100; x++ {
		key := strconv.Itoa(x)
		repo.Save(key, ent.Item{Title: fmt.Sprintf("Item %s", key)})
	}

	http.HandleFunc("/", show)
	http.ListenAndServe(":8080", nil)
}
