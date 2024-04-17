package main

import "fmt"

/*
	    To get row and column of an matrix global index
		row = index / number of col
		col = index % number of col
*/
func main() {
	matrix := [][]int{
		{1},
	}
	fmt.Println(searchMatrixOptimize(matrix, 7))
}

func searchMatrix(matrix [][]int, target int) bool {
	spreadSlice := []int{}
	for _, v := range matrix {
		spreadSlice = append(spreadSlice, v...)
	}
	low, high := 0, len(spreadSlice)-1
	for low <= high {
		mid := (low + high) / 2
		if spreadSlice[mid] < target {
			low = mid + 1
		} else if spreadSlice[mid] > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrixOptimize(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	low, high := 0, rows*cols-1
	for low <= high {
		mid := (low + high) / 2
		rowMid := mid / cols
		colMid := mid % cols
		valMid := matrix[rowMid][colMid]
		if valMid > target {
			high = mid - 1
		} else if valMid < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
