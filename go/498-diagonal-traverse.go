func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return nil }
    n, m := len(matrix), len(matrix[0])
    var ret []int
    up := true
    for i := 0; i < m + n + 1; i++ {
        var x, y int
        if up { 
            if i < n { x, y = i, 0 } else { x, y = n - 1, i - n + 1 }
            for ; x >= 0 && y < m; x, y = x - 1, y + 1 { ret = append(ret, matrix[x][y]) }
        } else {
            if i < m { x, y = 0, i } else { x, y = i - m + 1, m - 1 }
            for ; x < n && y >= 0; x, y = x + 1, y - 1 { ret = append(ret, matrix[x][y]) }
        }
        up = !up
    }
    return ret
}
