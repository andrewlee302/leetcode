func numFriendRequests(ages []int) int {
    counts := make([]int, 121)
    for i := 0; i < len(ages); i++ { counts[ages[i]]++ }
    sum := 0
    for a := 1; a <= 120; a++ {
        if counts[a] == 0 { continue }
        for b := a/2 + 7 + 1; b <= a; b++ {
            if counts[b] != 0 {
                if b != a {
                    sum += counts[b]*counts[a]
                } else {
                    sum += counts[a] * (counts[a] - 1)
                }
            }
        }
    }
    return sum
}
