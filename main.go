package main

import (
	"fmt"
	"main/list"
)

func main() {
	alpha := list.LinkedList{}
	alpha.Add(8)
	alpha.Add(2)
	alpha.Add(8)
	alpha.Add(1)
	alpha.AddOnIndex(244, 3)
	val, _ := alpha.Get(3)
	fmt.Printf("%d is the first value in the Alpha list", val)
}
