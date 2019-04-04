/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        if n == 0 { return -1 }
        candidate := 0
        for i := 1; i < n; i++ {
            if knows(candidate, i) { candidate = i }
        }
        for i := 0; i < n; i++ {
            if i != candidate && (!knows(i, candidate) || knows(candidate, i)) { return -1 }
        }
        return candidate
    }
}
