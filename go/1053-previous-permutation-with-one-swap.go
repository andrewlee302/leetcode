func prevPermOpt1(A []int) []int {
    Outter:
    for i := len(A) - 1; i > 0; i-- {
        if A[i] < A[i-1] {
            for j := len(A) - 1; j >= i; j-- {
                if A[j] < A[i-1] {
                    A[j], A[i-1] = A[i-1], A[j]
                    break Outter
                }
            }
        }
    }
    return A
}
