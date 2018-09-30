package main

/* This function reverses a 2D array (assumes the array lengths are equal) */
func reverseBoardArray(array [][]int) [][]int {
	size := len(array)
	reversed := make([][]int, size)

	for x := 0; x < size; x++ {
		reversed[x] = make([]int, size)
		for y := 0; y < size; y++ {
			reversed[x][y] = array[7-x][7-y]
		}
	}

	return reversed
}
