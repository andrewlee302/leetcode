import "container/list"
func updateBoard(board [][]byte, click []int) [][]byte {
    x, y := click[0], click[1]
    if board[x][y] == 'M' { 
        board[x][y] = 'X'
        return board
    }
    moves := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}, {1,1}, {1,-1}, {-1,1}, {-1,-1}}
    queue := list.New()
    queue.PushBack([]int{x, y})
    board[x][y] = 'B'
    for queue.Len() != 0 {
        e := queue.Front()
        queue.Remove(e)
        c := e.Value.([]int)
        c_x, c_y := c[0], c[1]
        mineCnt := byte(0)
        for _, move := range moves {
            _x, _y := c_x + move[0], c_y + move[1]
            if _x >= 0 && _x < len(board) && _y >= 0 && _y < len(board[0]) {
                if board[_x][_y] == 'M' { mineCnt++ }
            }
        }
        if mineCnt > 0 { board[c_x][c_y] = mineCnt + '0' } else {
            for _, move := range moves {
                _x, _y := c_x + move[0], c_y + move[1]
                if _x >= 0 && _x < len(board) && _y >= 0 && _y < len(board[0]) {
                    if board[_x][_y] == 'E' { 
                        queue.PushBack([]int{_x, _y}) 
                        board[_x][_y] = 'B'
                    }
                }
            }
        }
    }
    return board
}
