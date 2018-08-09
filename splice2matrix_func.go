package main

import (
	"fmt"
)

func main() {
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	matrix := make([][]int, 3)
	
	fmt.Println(splitslice(is, matrix))
}

func splitslice(a []int, b [][]int) [][]int {
	temp := []int{}
	row := 0
	for _, v := range a {

		temp = append(temp, v)

		if len(temp)%4 == 0 {

			b[row] = temp

			temp = nil
			row++
		}

	}
	return b
}
