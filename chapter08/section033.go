package main

import (
	"errors"
	"fmt"
)

type MyErr struct{}

func (me *MyErr) Error() string {
	return "gogogo"
}

func errorMethod() error {
	return &MyErr{}
}
func errorMethod2() error {
	e := errorMethod()
	if e != nil {
		return fmt.Errorf("%w", &MyErr{})
	} else {
		return &MyErr{}
	}
}

func main() {
	e := errorMethod2()

	if e != nil {
		if errors.Is(e, &MyErr{}) {
			fmt.Println("MyErr")
		} else {
			fmt.Println("error not myErr")
		}
	} else {
		fmt.Println("not error")
	}
}
