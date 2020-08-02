package main

import (
    "fmt"
    "strings"
)

func computer(altBoard [][]string) (int int) {
    var count int
    if count == 1 {
        //return altBoard
    }
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            if altBoard[row][col] == "_" {
                copyBoard := makeCopy(altBoard)
                copyBoard[row][col] = "O"
                
                computer(copyBoard)
            }
        }
    }
    //return altBoard
}

func makeCopy(board [][]string) [][]string {
    return board;
}

func move(n [][]string) [][]string {
    altBoard := makeCopy(n)
    row, col := computer(altBoard)
    n[row][col] = "O"
    return n
}

func main() {
    board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"}
    }
    for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", strings.Join(board[i], " "))
    }
    var col int
    var row int
    for moves := 0; moves < 10; moves++ {
        fmt.Println("Enter the column.")
        fmt.Scanln(&col)
        fmt.Println("Enter the row.")
        fmt.Scanln(&row)
        board[row-1][col-1] = "X"
        for i := 0; i < 3; i++ {
            fmt.Printf("%v\n", strings.Join(board[i], " "))
        }
        //computer(makeCopy(board))
        board = move(board)
        for i := 0; i < 3; i++ {
            fmt.Printf("%v\n", strings.Join(board[i], " "))
        }
    }
    
}