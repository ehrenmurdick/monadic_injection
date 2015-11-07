package main

import (
	"fmt"
	"net/http"
	"strconv"

	"./ent"
	"./repos"
)

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

// To see per-request error handling of a bad repo,
// try swapping the above for this:
// var repo = repos.
// 	OpenBadItemRepo()

// Note that it would need a little more love to
// make the maybeRepo re-try getting a good repo
// per request, like what you see with Rails

// "return" monad operation callback
// (with the correct response writer closed in)
func writeResponse(w http.ResponseWriter) func(string) error {
	return func(str string) error {
		_, err := fmt.Fprintf(w, str)
		return err
	}
}

// Error handler that writes the error to the http response!
// (with the correct response writer closed in)
func writeError(w http.ResponseWriter) func(error) error {
	return func(err error) error {
		fmt.Fprintf(w, err.Error())
		return err
	}
}

// Error handler that writes errors to stdout!
func logError(err error) error {
	fmt.Println(err.Error())
	return err
}

// Error handler that just panics!
func panicErr(err error) error {
	panic(err.Error())
	return err
}

// GET /:key
// renders item's title
func show(w http.ResponseWriter, r *http.Request) {
	var key string
	key = r.URL.Path[1:]

	repo.
		Get(key).
		// a trivial example. a more
		// realistic handler could render
		// an entity to json, for example.
		GetTitle().
		Within(writeResponse(w)).
		Handle(writeError(w)).
		Handle(logError)
}

func main() {
	// pre-load the repo with items up to 10
	for x := 0; x < 10; x++ {
		key := strconv.Itoa(x)
		repo.Save(key, ent.Item{Title: fmt.Sprintf("Item %s", key)})
	}

	http.HandleFunc("/", show)
	http.ListenAndServe(":8080", nil)
}
