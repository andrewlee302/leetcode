// Use golang library.
import "math/rand"
import "sort"

type Solution struct {
    accum []int
    sum int
}

func Constructor(w []int) Solution {
    accum := make([]int, len(w))
    sum := 0
    for i, n := range w {
        accum[i] = sum
        sum += n
    }
    return Solution{accum, sum}
}

func (this *Solution) PickIndex() int {
    v := rand.Intn(this.sum)
    idx := sort.Search(len(this.accum), func (i int) bool{ return this.accum[i] > v }) // first position that satisfy the bool function.
    return idx - 1
}

// Iterative.
import "math/rand"
type Solution struct {
    accum []int
    sum int
}

func Constructor(w []int) Solution {
    accum := make([]int, len(w))
    sum := 0
    for i, n := range w {
        accum[i] = sum
        sum += n
    }
    return Solution{accum, sum}
}

func (this *Solution) PickIndex() int {
    target := rand.Intn(this.sum)
    l, r := 0, len(this.accum)-1
    for l <= r {
        mid := l + (r-l)/2
        midVal := this.accum[mid]
        if midVal <= target {
            l = mid + 1
        } else {
            if mid == l || this.accum[mid-1] <= target {
                return mid-1
            } else {
                r = mid-1
            }
        }
    }
    return len(this.accum)-1
}
