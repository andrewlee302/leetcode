// max-heap & min-heap, cut it into two halves. 
// O(logN) for every AddNum(), O(1) for every FindMedian().
import "container/heap"
type MedianFinder struct {
    // Left half is one longer than or equal to right one.
    left MaxHeap
    right MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int)  {
    if this.left.Len() == this.right.Len() {
        // push to the right queue, then pop the minimal value into the left queue.
        heap.Push(&this.right, num)
        heap.Push(&this.left, heap.Pop(&this.right).(int))
    } else { // left.Len() > this.right.Len()
        // push to the left queue, then pop the maximum value into the right queue
        heap.Push(&this.left, num)
        heap.Push(&this.right, heap.Pop(&this.left).(int))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.left.Len() == this.right.Len() {
        return float64(this.left.IntHeap[0] + this.right.IntHeap[0]) / 2.0
    } else {
        return float64(this.left.IntHeap[0])
    }
}


type MaxHeap struct { IntHeap }
type MinHeap struct { IntHeap }
func (h MinHeap) Less (i ,j int) bool { return h.IntHeap[i] < h.IntHeap[j] }
func (h MaxHeap) Less (i ,j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
type IntHeap []int
func (h IntHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] } 
func (h IntHeap) Len() int { return len(h) }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) } // add x as element Len()
func (h *IntHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}

// insert sort.
// O(N) for every AddNum(), O(1) for every FindMedian().
type MedianFinder struct {
    nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int)  {
    pos := lowerBound(this.nums, num)
    this.nums = append(this.nums, 0)
    for i := len(this.nums) - 2; i >= pos; i-- {
        this.nums[i+1] = this.nums[i]
    }
    this.nums[pos] = num
}

func lowerBound(nums []int, num int) int {
    if len(nums) == 0 { return 0 }
    l, r := 0, len(nums) - 1
    var mid, pos int
    for l <= r {
        mid = l + (r-l)/2
        if num <= nums[mid] {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    if num <= nums[mid] {
        pos = mid
    } else {
        pos = mid + 1
    }
    return pos
}

func (this *MedianFinder) FindMedian() float64 {
    return float64(this.nums[(len(this.nums)-1)/2] + this.nums[len(this.nums)/2])/2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
