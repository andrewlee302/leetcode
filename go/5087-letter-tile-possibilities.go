// permutations + combinations + recursion
var permus []int = []int{1, 1, 2, 6, 24, 120, 720, 5040}
func numTilePossibilities(tiles string) int {
    dict := make([]int, 26)
    for i := 0; i < len(tiles); i++ { dict[tiles[i]-'A']++ }
    buf := make([]int, 26)
    sum := 0
    helper(dict, 0, buf, 0, &sum)
    return sum - 1
}

func helper(dict []int, dictIdx int, buf []int, bufIdx int, sum *int) {
    if dictIdx == len(dict) {
        fmt.Println(buf)
        cnt, divisor := 0, 1
        for i := 0; i < len(buf); i++ {
            cnt += buf[i]
            if buf[i] == 0 { continue }
            divisor *= permus[buf[i]]
        }
        *sum += permus[cnt]/divisor
        return
    }
    for i := 0; i <= dict[dictIdx]; i++ {
        buf[bufIdx] = i
        helper(dict, dictIdx+1, buf, bufIdx+1, sum)
    }
}

