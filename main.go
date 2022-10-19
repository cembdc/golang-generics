package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	// var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// arr := []int{1, 2, 4}
	// fmt.Println("keys:", MapKeys(m))

	// got := Map(arr, func(item int) int { return item })

	// fmt.Println("arr:", got)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

}
