func isMonotonic(A []int) bool {
    if len(A) < 3 { return true }
    status := 0 // 0: equal or unkown, 1: increase, 2: decrease
    for i := 1; i < len(A); i++ {
        if A[i-1] < A[i] {
            if status == 2 { return false }
            status = 1
        }
        if A[i-1] > A[i] {
            if status == 1 { return false }
            status = 2
        }
    }
    return true
}
