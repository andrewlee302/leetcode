// recursion with pruning.
func solveSudoku(board [][]byte)  {
    rows, columns, boxes := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
    for i := 0; i < 9; i++ {
        rows[i], columns[i], boxes[i] = make([]bool, 10), make([]bool, 10), make([]bool, 10)
    }
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                changeNum(board, i, j, board[i][j], rows, columns, boxes, true)
            }
        }
    }
    solved := false
    for k := byte(1); k <= 9; k++ {
        backtracking(board, 0, 0, k+'0', rows, columns, boxes, &solved)
        if solved { return }
    }
}

func backtracking(board [][]byte, row, col int, val byte, rows, columns, boxes [][]bool, solved *bool) {
    if *solved { return }
    if row == 9 && col == 0 {
        *solved = true
        return
    }
    newRow, newCol := row, col
    if col == 8 { newRow, newCol = newRow + 1, 0 } else { newCol++ }
    if board[row][col] != '.' {
        backtracking(board, newRow, newCol, val, rows, columns, boxes, solved)
    } else {
        if !isValid(row, col, val, rows, columns, boxes) { return }
        changeNum(board, row, col, val, rows, columns, boxes, true)
        for k := byte(1); k <= 9; k++ {
            backtracking(board, newRow, newCol, k+'0', rows, columns, boxes, solved)
            if *solved { return }
        }
        changeNum(board, row, col, val, rows, columns, boxes, false)
    }   
}

func isValid(row, col int, val byte, rows, columns, boxes [][]bool) bool {
    digit := int(val - '0')
    boxesId := row/3*3 + col/3
    return !(rows[row][digit] || columns[col][digit] || boxes[boxesId][digit])
}

func changeNum(board [][]byte, row, col int, val byte, rows, columns, boxes [][]bool, place bool) { 
    // remove the num if place is false.
    digit := int(val - '0')
    boxesId := row/3*3 + col/3
    if place { board[row][col] = val } else { board[row][col] = '.' }
    rows[row][digit] = place
    columns[col][digit] = place
    boxes[boxesId][digit] = place
}
