func longestArithSeqLength(A []int) int {
    longestLen := 0
    for i := 0; i < len(A); i++ {
        for j := i+1; j < len(A); j++ {
            // A[i], A[j] ...
            diff := A[j] - A[i]
            cnt := 2
            prev := A[j]
            for k := j+1; k < len(A); k++ {
                if diff == A[k] - prev {
                    cnt++
                    prev = A[k]
                }
            }
            if cnt > longestLen { longestLen = cnt }
        }
    }
    return longestLen
}
