func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return nil }
    moves := [][]int{{0,1}, {1,0},{0,-1},{-1,0}}
    direction := 0
    rl, rr, cl, cr := 0, len(matrix)-1, 0, len(matrix[0])-1
    currRow, currCol := 0, 0
    ret := make([]int, 0, len(matrix)*len(matrix[0]))
    ret = append(ret, matrix[currRow][currCol])
    for len(ret) != cap(ret) {
        move := moves[direction]
        for {
            nextRow, nextCol := currRow + move[0], currCol + move[1]
            if nextRow < rl || nextRow > rr || nextCol < cl || nextCol > cr { break }
            currRow, currCol = nextRow, nextCol
            ret = append(ret, matrix[currRow][currCol])
        }
        switch (direction) {
        case 0: rl+=1
        case 1: cr-=1
        case 2: rr-=1
        case 3: cl+=1
        }
        direction = (direction+1)%4
    }
    return ret
}
