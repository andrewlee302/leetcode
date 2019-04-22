func maxSumTwoNoOverlap(A []int, L int, M int) int {
    dpLeft, dpRight := maxiumLeftAndRight(A, L, M)
    max := 0
    for i := L-1; i <= len(A) - M - 1; i++ {
        temp := dpLeft[i] + dpRight[i+1] 
        if temp > max { max = temp }
    }
    dpLeft, dpRight = maxiumLeftAndRight(A, M, L)
    for i := M-1; i <= len(A) - L - 1; i++ {
        temp := dpLeft[i] + dpRight[i+1] 
        if temp > max { max = temp }
    }
    return max
}

func maxiumLeftAndRight(A []int, left, right int) ([]int, []int) {
    dpLeft, dpRight := make([]int, len(A)), make([]int, len(A))
    sum := 0
    for i := 0; i < left; i++ { sum += A[i] }
    dpLeft[left-1] = sum
    for i := left; i < len(A); i++ { 
        sum += A[i]
        sum -= A[i-left]
        if sum > dpLeft[i-1] { 
            dpLeft[i] = sum
        } else {
            dpLeft[i] = dpLeft[i-1]
        }
    }
    sum = 0
    for i := len(A) - 1; i >= len(A) - right; i-- { sum += A[i] }
    dpRight[len(A) - right] = sum
    for i := len(A) - right - 1; i >= 0; i-- { 
        sum += A[i]
        sum -= A[i+right]
        if sum > dpRight[i+1] { 
            dpRight[i] = sum
        } else {
            dpRight[i] = dpRight[i+1]
        }
    }
    return dpLeft, dpRight
}
