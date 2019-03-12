func subarraySum(nums []int, k int) int {
    m := make(map[int]int)
    sum := 0
    cnt := 0
    m[0] = 1
    for _, v := range nums {
        sum += v
        cnt += m[sum-k]
        m[sum]++
    }
    return cnt
}
