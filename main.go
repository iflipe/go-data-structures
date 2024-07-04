package main

import (
	"main/list"
	"math"
)

func main() {
	list := list.ArrayList{}
	list.Init(10)
	for i := 0; i < 10; i++ {
		list.Add(i * int(math.Pow(-1.0, float64(i))))
	}
	for i := 0; i < 10; i++ {
		a, _ := list.Get(i)
		println(a)
	}
	println("----------------------")
	list.InsertionSort()
	for i := 0; i < 10; i++ {
		a, _ := list.Get(i)
		println(a)
	}
}
