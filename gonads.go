package main

import (
	"fmt"
	"os"
)

// primitives

type database struct {
	dburl string
}

// inits
func fromString(str string) maybeString {
	return maybeString{
		value: str,
	}
}

// monads
type maybeDatabase struct {
	value database
	err   error
}

type maybeString struct {
	value string
	err   error
}

// operations
func (in maybeString) andThen(bl func(string) maybeString) maybeString {
	if in.err == nil {
		return bl(in.value)
	} else {
		return in
	}
}

func (in maybeString) within(bl func(string) error) maybeString {
	if in.err == nil {
		return maybeString{
			value: in.value,
			err:   bl(in.value),
		}
	} else {
		return in
	}
}

func (in maybeString) handle(bl func(error) error) maybeString {
	if in.err != nil {
		bl(in.err)
	}
	return in
}

func (in maybeString) andThenDatabase(bl func(string) maybeDatabase) maybeDatabase {
	if in.err == nil {
		return bl(in.value)
	} else {
		return maybeDatabase{
			err: in.err,
		}
	}
}

// functors
func getEnv(key string) maybeString {
	return maybeString{
		value: os.Getenv(key),
	}
}

// side effects
func printStr(str string) (err error) {
	_, err = fmt.Println(str)
	return
}

func printErr(err error) error {
	_, err = fmt.Println(err.Error())
	return err
}

// main
func main() {
	getEnv("DB").
		within(printStr).
		handle(printErr)
}
