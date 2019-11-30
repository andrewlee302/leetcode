func largestRectangleArea(heights []int) int {
    stack := [][]int{{-1, 0}} // index, height.
    ret := 0
    heights = append(heights, 0)
    for i, h := range heights {
        for len(stack) != 0 {
            top := stack[len(stack)-1]
            if top[1] <= h { break }
            ret = max(ret, top[1]*(i-stack[len(stack)-2][0]-1))
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, []int{i, h})
    }
    return ret
}

func max(a, b int) int { if a > b { return a } else { return b } }
