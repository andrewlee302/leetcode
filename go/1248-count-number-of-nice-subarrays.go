func numberOfSubarrays(nums []int, k int) int {
    var cnts []int
    l := 0
    for _, num := range nums {
        if num & 1 == 0 {
            l++
        } else {
            cnts = append(cnts, l)
            l = 0
        }
    }
    cnts = append(cnts, l)
    fmt.Println(cnts)
    ret := 0 
    for i := 0; i + k < len(cnts); i++ {
        ret += (cnts[i]+1) * (cnts[i+k]+1)
    }
    return ret
}
