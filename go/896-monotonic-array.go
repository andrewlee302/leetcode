func isMonotonic(A []int) bool {
    flag := false
    increasing := true
    for i := 1; i < len(A); i++ {
        if flag && (A[i] > A[i-1] && !increasing || A[i] < A[i-1] && increasing) {
            return false
        }
        if A[i] != A[i-1] { flag = true }
        if A[i] < A[i-1] { increasing = false }
    }
    return true
}
