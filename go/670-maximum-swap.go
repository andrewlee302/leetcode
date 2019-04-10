import "strconv"
func maximumSwap(num int) int {
    numBytes := []byte(strconv.Itoa(num))
    lateIndices := make([]int, 10) // latest index of '0'~'9'
    for i, b := range numBytes { lateIndices[int(b-'0')] = i }
    for i := 0; i < len(numBytes); i++ {
        for j := 9; j > int(numBytes[i] - '0'); j-- {
            occur := lateIndices[j]
            if lateIndices[j] > i {
                numBytes[i], numBytes[occur] = numBytes[occur], numBytes[i]
                ret, _ := strconv.Atoi(string(numBytes))
                return ret
            }
        }
    }
    return num
}
