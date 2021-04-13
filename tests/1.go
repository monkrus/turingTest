package main

import "fmt"

func makeSquares(array [10]int) {

	for index, elem := range array {
		array[index] = elem * elem
	}

}

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	makeSquares(a)

	fmt.Println(a)
}

// ANSWER [0 1 2 3 4 5 6 7 8 9]
// https://play.golang.org/p/itnZYDiqdIN
