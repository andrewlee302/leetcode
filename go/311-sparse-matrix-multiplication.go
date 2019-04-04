func multiply(A [][]int, B [][]int) [][]int {
    if len(A) == 0 || len(A[0]) == 0 || len(B) == 0 || len(B[0]) == 0 { return nil }
    mapsA := make([]map[int]int, len(A))
    for i := 0; i < len(A); i++ {
        mapsA[i] = make(map[int]int)
        for j := 0; j < len(A[0]); j++ {
            if A[i][j] != 0 { mapsA[i][j] = A[i][j] }
        }
    }
    ret := make([][]int, len(A))
    for i := 0; i < len(A); i++ {
        ret[i] = make([]int, len(B[0]))
        for j := 0; j < len(B[0]); j++ {
            sum := 0
            for idx, val := range mapsA[i] {
                sum += val * B[idx][j]
            }
            ret[i][j] = sum
        }
    }
    return ret
}
