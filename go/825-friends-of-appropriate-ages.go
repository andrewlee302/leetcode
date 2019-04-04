func numFriendRequests(ages []int) int {
    counts := make([]int, 121)
    for i := 0; i < len(ages); i++ { counts[ages[i]]++ }
    sum := 0
    lowBound, setLowBound := 1, true
    for a := 1; a <= 120; a++ {
        if counts[a] == 0 { continue }
        for b := lowBound; b <= a; b++ {
            if b > a/2 + 7 && counts[b] != 0 { 
                if !setLowBound { lowBound, setLowBound = b, true }
                if b != a { sum += counts[b]*counts[a] } else {
                    sum += counts[a] * (counts[a] - 1)            
                }
            }
        }
        setLowBound = false
    }
    return sum
}
