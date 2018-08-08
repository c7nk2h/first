package main

import (
	"fmt"
)

func main() {
	temp := []int{}
	matrix := make([][]int, 3)

	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	row := 0
	for i, v := range is {

		temp = append(temp, v)
		fmt.Println(i, temp)

		if len(temp)%4 == 0 {

			matrix[row] = temp
			fmt.Println(matrix)
			
			temp = nil
			row++
		}

	}

	fmt.Println(is)
	fmt.Println(matrix)
}
