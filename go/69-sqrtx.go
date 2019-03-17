func mySqrt(x int) int {
    left, right := 0, x
    for left <= right {
        mid := (left + right) / 2
        if mid*mid == x || mid*mid < x && (mid+1)*(mid+1) > x  {
            return mid
        } 
        if mid*mid > x {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}
