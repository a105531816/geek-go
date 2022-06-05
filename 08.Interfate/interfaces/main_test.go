package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	var a interface{}
	switch c := a.(type) {
	case string:
		fmt.Println(c, "cstrings")
	case int:
		fmt.Println("int")
	}
	typee, ok := a.(string)
	fmt.Println(typee,ok)
}

type TestBox struct {
}

func (tb *TestBox) Close() error {
	return nil
}
