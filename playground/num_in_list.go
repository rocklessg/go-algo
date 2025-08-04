package playground

import "slices"

// func NumInList(list []int, num int) bool {
// 	for i := range list {
// 		if list[i] == num {
// 			return true
// 		}
// 	}
// 	return false
// }


func NumInList(list []int, num int) bool {
	return slices.Contains(list, num)
}