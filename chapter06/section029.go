package main

import "fmt"

type Dog struct{}

func (d *Dog) Bark() string {
	return "Bow"
}

type BullDog struct {
	Dog
}

type ShibaInu struct {
	Dog
}

func (s *ShibaInu) Bark() string {
	return "ワン"
}
func DogVoice(d *Dog) string { return d.Bark() }

func main() {
	bd := &BullDog{}
	fmt.Println(bd.Bark())

	// コンポジション、ミックスイン、トレイトに近い概念
	si := &ShibaInu{}
	fmt.Println(si.Bark())
}
