import "math/rand"
type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{nums}
}


func (this *Solution) Pick(target int) int {
    cnt := 0
    var chosen int
    for i := 0; i < len(this.nums); i++ {
        if this.nums[i] == target { cnt++ } else { continue }
        if rand.Intn(cnt) == 0 { chosen = i }
    }
    return chosen
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
