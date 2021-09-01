package main

import "fmt"

func main() {
	input := []int{1, 2, 0, 5, 0, 2, 5, 12, 27, 04, 0}

	var temp, result []int

	for _, v := range input {
		if v == 0 {
			temp = append(temp, v)
		} else {
			result = append(result, v)
		}
	}
	result = append(result, temp...)
	fmt.Println(input)
	fmt.Println(result)
}
