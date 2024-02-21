package main

import (
	"slices"
)

func InitializeMatrix(rows int, columns int) [][]int {
	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)
	}

	// Initialize the top row (0 to len(word2))
	for j := 0; j < columns; j++ {
		matrix[0][j] = j
	}

	// Initialize the first column (0 to len(word1))
	for i := 0; i < rows; i++ {
		matrix[i][0] = i
	}

	return matrix
}

func LogiDistance(word1 string, word2 string) int {
	rows := len(word1) + 1
	columns := len(word2) + 1

	matrix := InitializeMatrix(rows, columns)

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if word1[i-1] == word2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				best_action := []int{matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1]}
				matrix[i][j] = 1 + slices.Min(best_action)
			}
		}
	}

	return matrix[rows-1][columns-1]
}
