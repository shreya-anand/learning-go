package main

import (
    "fmt"
    "strings"
)

func main() {
    board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
    }
    for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", strings.Join(board[i], " "))
    }
    var col int
    var row int
    fmt.Println("Enter the column.")
    fmt.Scanln(&col)
    fmt.Println("Enter the row.")
    fmt.Scanln(&row)
    board[row-1][col-1] = "X"
    for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", strings.Join(board[i], " "))
    }
}