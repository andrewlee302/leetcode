func numSubmatrixSumTarget(matrix [][]int, target int) int {
    ret := 0
    for i := 0; i < len(matrix); i++ { // all the rectangle starts with the i-th row.
        prefixSums := make([]int, len(matrix[0]))
        for j := i; j < len(matrix); j++ { // ends with the j-th row.
            prefixSumMap := make(map[int]int) // prefix sums between i-th and j-th row.
            prefixSumMap[0] = 1
            rowSum := 0 // prefix sum in the j-th row.
            for k := 0; k < len(matrix[0]); k++ { 
                rowSum += matrix[j][k]
                prefixSums[k] += rowSum
                ret += prefixSumMap[prefixSums[k] - target]
                prefixSumMap[prefixSums[k]]++
            }
        }
    }
    return ret
}
