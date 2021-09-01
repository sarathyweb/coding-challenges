package main

import "fmt"

func main() {
	input := []int{1, 2, 0, 5, 0, 2, 5, 12, 27, 04, 0}

	var insertZeroPosition int
	var result []int

	for _, v := range input {
		if v != 0 {
			result = append(result, v)
			insertZeroPosition++
		}
	}

	for i := insertZeroPosition;i<len(input);i++{
		result = append(result, 0)

	}

	fmt.Println(input)
	fmt.Println(result)

}
