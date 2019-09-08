func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
    r, c := 0, len(matrix[0])-1
    for r < len(matrix) && c >= 0 {
        if target < matrix[r][c] {
            c--
        } else if target > matrix[r][c] {
            r++
        } else {
            return true
        }
    }
    return false
}
