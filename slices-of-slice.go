/**

slice 的 slice
slice 可以包含任意的类型，包括另一个 slice。
**/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe bord
	game := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns

	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)

}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}

}
