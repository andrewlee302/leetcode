func isGoodArray(nums []int) bool {
    x := nums[0]
    for _, num := range nums {
        x = gcd(x, num)
    }
    return x == 1
}

func gcd(i, j int) int {
    for j > 0 {
        a := i % j
        i = j
        j = a
    }
    return i
}
