package io

import "fmt"

func PrintErr(err error) error {
	fmt.Println("error:", err)
	return err
}

func PrintStr(str string) error {
	_, err := fmt.Println(str)
	return err
}
