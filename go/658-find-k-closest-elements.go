func findClosestElements(arr []int, k int, x int) []int {
    if len(arr) == 0 || k > len(arr) { return nil }
    left, right := binarySearch(arr, 0, len(arr) - 1, x)
    fmt.Println(left, right)
    cnt := 0
    if left == right { 
        cnt++
        left--
        right++
    }
    for ; cnt < k && left >= 0 && right < len(arr); cnt++ {
        diffLeft := x - arr[left]
        diffRight := arr[right] - x
        if diffLeft <= diffRight {
            // left is priorited
            left--
        } else {
            right++
        }
    }
    if left < 0 { 
        right += (k - cnt) 
    } else if right >= len(arr) { left -= (k - cnt) }
    return arr[left+1: right]
}

func binarySearch(arr []int, left, right, x int) (int, int) {
    if left > right { return right, left }
    mid := left + (right - left) / 2
    if x < arr[mid] { 
        return binarySearch(arr, left, mid-1, x) 
    } else if x > arr[mid] { 
        return binarySearch(arr, mid+1, right, x) 
    } else { return mid, mid }
}
