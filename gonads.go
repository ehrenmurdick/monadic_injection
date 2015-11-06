package main

import "fmt"

type maybe struct {
	err   error
	value string
}

func fromValue(v string) maybe {
	return maybe{
		value: v,
	}
}

func (m maybe) andThen(bl func(string) maybe) maybe {
	if m.err != nil {
		return m
	} else {
		return bl(m.value)
	}
}

func (m maybe) handle(bl func(error) maybe) maybe {
	if m.err != nil {
		return bl(m.err)
	} else {
		return m
	}
}

func getUrl(url string) maybe {
	return maybe{
		value: "url",
	}
}

func fetch(url string) maybe {
	return maybe{
		value: fmt.Sprintf("fetched \"%s\" and got \"%s\"", url, "html"),
	}
}

func pr(response string) maybe {
	println(response)
	return maybe{
		value: response,
	}
}

func raise(err error) maybe {
	panic(err.Error())
	return maybe{
		err: err,
	}
}

func main() {
	getUrl("path").
		andThen(fetch).
		andThen(pr).
		handle(raise)
}
