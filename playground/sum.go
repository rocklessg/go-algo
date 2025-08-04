package playground

// func Sum(list []int) int {
// 	sum := 0
// 	for i := range list {
// 		sum+=list[i]
// 	}
// 	return sum
// }


func Sum(numbers []int) int {
	result := 0
	for _, value := range numbers {
		result += value
	}
	return result
}