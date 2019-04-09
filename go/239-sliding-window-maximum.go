// monotonic queue
import "container/list"
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) < k { return []int{} }
    var ret []int
    queue := list.New()
    for i := 0; i < len(nums); i++ {
        for queue.Len() != 0 {
            e := queue.Back()
            if e.Value.(int) < nums[i] {
                queue.Remove(e)
            } else { break }
        }
        queue.PushBack(nums[i])
        // We can have a window.
        if i >= k - 1 {
            e := queue.Front()
            currWindowMaximum := e.Value.(int)
            ret = append(ret, currWindowMaximum)
            if nums[i-k+1] == currWindowMaximum {
                queue.Remove(e)
            }
        }
    }
    return ret
}
