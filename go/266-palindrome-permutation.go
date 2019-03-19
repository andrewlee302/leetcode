func canPermutePalindrome(s string) bool {
    m := make(map[byte]int) 
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }
    oddCnt := 0
    for _, cnt := range m {
        if cnt % 2 == 1 { oddCnt++ }
    }
    if len(s) % 2 == 0 && oddCnt == 0 || len(s) % 2 == 1 && oddCnt == 1 {
        return true
    }
    return false
}
