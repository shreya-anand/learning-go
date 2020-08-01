package main

import (
    "fmt"
    "strings"
)

func computer(board [][]string) [][]string {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "_" {
                
            }
        }
    }
    return b
}
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