package main

import "fmt"

/*
工具函数，通用，类型变了，可能是泛型

*/

func mapKeys[T comparable](mp map[T]int, filter func(k T) bool) []T {
	res := []T{}
	for k := range mp {
		if filter(k) {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	mp := map[int]int{
		4: 1,
		3: 2,
		2: 4,
	}

	fmt.Println(mapKeys(mp, func(k int) bool {
		return k >= 3
	}))

}
