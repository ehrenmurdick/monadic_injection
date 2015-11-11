package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ehrenmurdick/monadic_injection/ent"
	"github.com/ehrenmurdick/monadic_injection/repos"
)

// initialize a repo (wrapped in a Result).
// Failure to open can be handled
// here by adding:
// .Handle(panicErr)
// That would quit the
// server if the db is configured
// incorrectly, for example.
// otherwise, you could defer
// handling the opening db error
// until a page request, which is more
// like what you see with Rails
// for example. That's the behavior right now.
var repo = repos.
	ReturnResultItemRepo(repos.NewItemRepo())

// Note that it would need a little more love to
// make the resultRepo re-try getting a good repo
// per request, like what you see with Rails

// function which returns a handler function with
// the response writer curried in
func writeResponse(w http.ResponseWriter) func(string) string {
	return func(str string) string {
		fmt.Fprintf(w, str)
		return str
	}
}

// function which returns a handler function with
// the response writer curried in
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

	// Each step in this chain returns another Result,
	// with the right methods on it, regardless of
	// whether the line failed or not. If a line fails,
	// then subsequent lines do nothing but wait until a
	// .Handle call
	repo.
		Get(key).
		// a trivial example. a more
		// realistic handler could render
		// an entity to json, for example.
		GetTitle().
		// Bind means "get the wrapper value out if it's there
		// and call this function. If the wrapped value is bad, ie there's
		// and error, don't call this function and just return self.
		Bind(writeResponse(w)).
		// Handle is the same as Bind, but for the error value instead
		// of the wrapped value. So it calls the the passed in
		// function if there IS an error, and ignores the passed in function
		// if there isn't.
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
