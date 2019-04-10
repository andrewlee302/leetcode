import "math/rand"
type Solution struct {
    sum []int
    w []int
    total int
}

func Constructor(w []int) Solution {
    sum := make([]int, len(w))
    total := 0
    for i, v := range w {
        total += v
        sum[i] = total
    }
    return Solution{sum, w, total}
}

func (this *Solution) PickIndex() int {
    target := rand.Intn(this.total) + 1
    l, r := 0, len(this.sum)-1 
    for l <= r {
        mid := l + (r-l)/2
        midVal := this.sum[mid]
        if midVal < target {
            l = mid + 1 
        } else if midVal > target {
            if mid == l || this.sum[mid-1] < target { 
                return mid 
            } else {
                r = mid-1
            }
        } else {
            return mid
        }
    }
    return 0
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
