package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	runGame()

}

func runGame() {
	var matrix = make([][]string, 20)

	matrix = initMatrix(matrix)

	for {
		time.Sleep(1 * time.Second)
		render(matrix)
		//testRender(matrix)
		matrix = updateState(matrix)
	}

}

func initMatrix(matrix [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	for r := 0; r < 20; r++ {
		matrix[r] = make([]string, 20)
		for c := range matrix {
			seed := rand.Intn(5)
			if seed < 3 {
				matrix[r][c] = "x"
			} else {
				matrix[r][c] = " "
			}
		}

	}
	return matrix
}

func updateState(m [][]string) [][]string {
	updatedMatrix := make([][]string, 20)
	// if 3 neighbors and dead, make alive
	// if > 3 neighbors, kill
	// if < 2 kill
	for r := range m {
		updatedMatrix[r] = make([]string, 20)
		copy(updatedMatrix[r], m[r])

		for c := range m[r] {
			//
			count := countNeighbors(m, r, c)
			if count == 3 {
				updatedMatrix[r][c] = "x"
			} else if count > 3 {
				updatedMatrix[r][c] = " "
			} else if count < 2 {
				updatedMatrix[r][c] = " "
			}
		}
	}

	return updatedMatrix
}

func testRender(m [][]string) {
	for r := range m {
		testRow := make([]string, 20)
		for c := range testRow {
			testRow[c] = strconv.Itoa(countNeighbors(m, r, c))
		}
		fmt.Println(testRow)
	}
	fmt.Println("=========================================")

}

func countNeighbors(m [][]string, r int, c int) int {
	var count = 0

	for rowOffset := 0; rowOffset <= 2; rowOffset++ {
		for colOffset := 0; colOffset <= 2; colOffset++ {
			rowIdx := r - 1 + rowOffset
			colIdx := c - 1 + colOffset

			if rowIdx >= 0 && rowIdx < 20 && colIdx >= 0 && colIdx < 20 {
				if !(rowIdx == r && colIdx == c) {
					if m[rowIdx][colIdx] == "x" {
						count++
					}
				}
			}

		}
	}
	return count
}

func render(m [][]string) {
	for r := range m {
		fmt.Println(m[r])
	}
	fmt.Println("=========================================")
}
