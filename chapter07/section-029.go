package main

import "fmt"

type MyErr struct{}

func (me *MyErr) Error() string { return "" }

func Apply() error {
	var err *MyErr = nil
	return err
}

func Apply2() error {
	var err error = nil
	return err
}

// Goのインターフェースは具象型の型情報と値の二つを要素とするデータ構造
func main() {
	fmt.Println(Apply() == nil)
	fmt.Println(Apply2() == nil)
}
