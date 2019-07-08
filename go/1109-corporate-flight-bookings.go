func corpFlightBookings(bookings [][]int, n int) []int {
    ret := make([]int, n+2) // The first and last one is guard.
    for _, booking := range bookings {
        ret[booking[0]] += booking[2]
        ret[booking[1]+1] -= booking[2]
    }
    for i := 1; i <= n; i++ { ret[i] += ret[i-1] }
    return ret[1:n+1]
}
