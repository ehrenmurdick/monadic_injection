package main

import (
	"errors"
	"fmt"
	"os"
)

type V interface{}

type printable interface {
	stringify() string
}

type maybe struct {
	err   error
	value V
}

type item struct {
	title string
	done  bool
	url   string
}

func newItem(s string, u string) (i item) {
	return item{
		title: s,
		done:  false,
		url:   u,
	}
}

func (i item) stringify() string {
	return fmt.Sprintf("{item title=\"%s\"  done=%+v url=%s}", i.title, i.done, i.url)
}

func fromValue(v V) maybe {
	return maybe{
		value: v,
	}
}

func (m maybe) andThen(bl func(V) maybe) maybe {
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

func getUrl(key V) maybe {
	return maybe{
		value: os.Getenv(key.(string)),
	}
}

var itemId int = 0

func fetch(url V) (m maybe) {
	itemId++
	if s, ok := url.(string); ok {
		m.value = newItem(fmt.Sprintf("item %d", itemId), s)
	} else {
		m.err = errors.New(fmt.Sprintf("couldn't fetch non-string url: %+v", url))
	}
	return
}

func pr(response V) (m maybe) {
	m.value = response
	if s, ok := response.(string); ok {
		println(s)
	} else if i, ok := response.(item); ok {
		println(i.stringify())
	} else {
		m.err = errors.New("couln't print non-printable thing")
	}
	return
}

func raise(err error) maybe {
	panic(err.Error())
	return maybe{
		err: err,
	}
}

func main() {
	getUrl("HOST").
		andThen(fetch).
		andThen(pr).
		handle(raise)

	getUrl("path").
		andThen(fetch).
		andThen(pr).
		handle(raise)
}
