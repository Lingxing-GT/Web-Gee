package main

import "fmt"

func main() {
	intArray := []int{3, 3}
	indexMap := map[int]int{}
	for k, v := range intArray {
		indexMap[v] = k
	}
	for k, v := range indexMap {
		fmt.Println("map[k]v", k, v)
	}
	for _, v := range intArray {
		fmt.Println("value:", v, "key:", indexMap[v])
	}
}
