// recursion.
func duplicateZeros(arr []int)  {
    helper(arr, 0)
}

func helper(arr []int, cnt int) {
    if len(arr) <= cnt { return }
    i := 0
    for i = 0; i < len(arr); i++ {
        if arr[i] == 0 {
            break
        }
    }
    if i != len(arr) { helper(arr[i+1:], cnt+1) }
    if i+cnt+1 < len(arr) { arr[i+cnt+1] = 0 }
    for j := min(i, len(arr)-cnt-1); j >= 0; j-- {
        arr[j+cnt] = arr[j]
    }
}

func min(a, b int) int { if a < b { return a } else { return b } }

// use another copy of the array.
func duplicateZeros(arr []int)  {
    cp := make([]int, len(arr))
    copy(cp, arr)
    p := 0
    for i := 0; i < len(cp) && p < len(cp); i++ {
        arr[p] = cp[i]
        p++
        if cp[i] == 0 && p < len(cp) {
            arr[p] = 0
            p++
        }
    }
}

