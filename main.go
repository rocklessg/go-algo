package main

import (
	"fmt"
	"github.com/rocklessg/go-algo/playground"
)

func main() {
	// num_in_list test
	//result := add(3, 4)
	// fmt.Println(playground.NumInList([]int{1, 2, 3, 4, 5}, 3)) // should return true
	// fmt.Println(playground.NumInList([]int{1, 2, 3, 4, 5}, 6)) // should return false

	// sum test
	// fmt.Println(playground.Sum([]int {1, 2, 3, 4, 5})) // should return 15
	// fmt.Println(playground.Sum([]int {10, 20, 30})) // should return 60

	// ReverseString test
	fmt.Println(playground.ReverseString("cat"))      // should return "tac"
	fmt.Println(playground.ReverseString("alphabet")) // should return "tebahpla"
}