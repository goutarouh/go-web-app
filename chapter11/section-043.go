package main

import "fmt"

type Jedi interface {
	HasForce() bool
}

type Knight struct {
}

func (k *Knight) HasForce() bool {
	return true
}

// ある構造体が特定のインターフェースを満たしているかどうか確かめる。
func main() {
	var _ Jedi = (*Knight)(nil)
	fmt.Println("")
}
