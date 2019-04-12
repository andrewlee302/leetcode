func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
    // assume {r0,c0} is valid.
    ret := make([][]int, 0, R*C)
    moves := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    targetMoveCnt := R*C
    ret = append(ret, []int{r0,c0})
    if len(ret) == targetMoveCnt { return ret }
    moveCnt := 0
    directionCnt := 0
    for {
        directionCnt++
        move := moves[(directionCnt-1)%len(moves)]
        steps := (directionCnt+1)/2
        for i := 0; i < steps; i++ {
            r0, c0 = r0 + move[0], c0 + move[1]
            if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
                moveCnt++
                ret = append(ret, []int{r0, c0})
                if len(ret) == targetMoveCnt {
                    return ret
                }
            }
        }
    }
}
