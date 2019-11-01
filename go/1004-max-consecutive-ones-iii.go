func longestOnes(A []int, K int) int {
    if len(A) == 0 { return 0 }
    left, right := 0, 0
    zeroCnt := 0
    ret := 0
    for right < len(A) {
        if A[right] == 0 { zeroCnt++ }
        for zeroCnt > K {
            if A[left] == 0 { zeroCnt-- }
            left++
        }
        ret = max(ret, right-left+1)
        right++
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }
