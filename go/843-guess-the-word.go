/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
 func findSecretWord(wordlist []string, master *Master) {
    match := make([][]int, len(wordlist))
    for i := 0; i < len(wordlist); i++ {
        match[i] = make([]int, len(wordlist))
        for j := 0; j <= i; j++ { // symmetric.
            cnt := 0
            for k := 0; k < 6; k++ { if wordlist[i][k] == wordlist[j][k] { cnt++ } }
            match[i][j], match[j][i] = cnt, cnt
        }
    }
    var candidates []int
    guessSet := make(map[int]bool)
    for i := 0; i < len(wordlist); i++ { candidates = append(candidates, i) }
    for {
        guessIdx := chooseMinmalWorstGuess(match, candidates, guessSet)
        matchCnt := master.Guess(wordlist[guessIdx])
        guessSet[guessIdx] = true
        if matchCnt == 6 { return }
        ptr := 0
        for _, c := range candidates {
            if match[guessIdx][c] == matchCnt {
                candidates[ptr] = c
                ptr++
            }
        }
        candidates = candidates[:ptr]
    }
}

func chooseMinmalWorstGuess(match [][]int, candidates []int, guessSet map[int]bool) int {
    if len(candidates) == 1 { return candidates[0] } // !Note!
    guess := []int{-1, len(match)}
    for i := 0; i < len(match); i++ {
        if guessSet[i] { continue }
        var group [7]int
        for _, c := range candidates { group[match[i][c]]++ }
        maxCnt := 0
        for _, cnt := range group { maxCnt = max(maxCnt, cnt) }
        if maxCnt < guess[1] { guess = []int{i, maxCnt} }
    }
    return guess[0]
}

func max(i, j int) int { if i > j { return i } else { return j } }

