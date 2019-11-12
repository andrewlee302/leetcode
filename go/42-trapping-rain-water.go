// Improved from the leftMax & rightMax version (two iterations). O(N), O(1).
func trap(height []int) int {
    if len(height) <= 2 { return 0 }
    left, right := 0, len(height) - 1
    leftMax, rightMax := height[left], height[right]
    ret := 0
    for left < right - 1 {
        var bar int
        if leftMax >= rightMax {
            right--
            bar = height[right]
            rightMax = max(bar, rightMax)
        } else {
            left++
            bar = height[left]
            leftMax = max(bar, leftMax)
        }
        ret += max(min(leftMax, rightMax) - bar, 0)
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }

// LeftMax & rightMax. O(N), O(N)
func trap(height []int) int {
    if len(height) <= 2 { return 0 }
    min_max := make([]int, len(height))
    leftMax, rightMax := 0, 0
    for i := 0; i < len(height); i++ {
        min_max[i] = leftMax
        leftMax = max(leftMax, height[i])
    }
    for i := len(height) - 1; i >= 0; i-- {
        min_max[i] = min(min_max[i], rightMax)
        rightMax = max(rightMax, height[i])
    }
    ret := 0
    for i := 0; i < len(height); i++ {
        ret += max(0, min_max[i] - height[i])
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }

// monotonic queue/stack. (strict decreasing). O(N), O(N).
func trap(height []int) int {
    var stack []int
    stack = append(stack, -1)
    water := 0
    for i, h := range height {
        for {
            topIdx := stack[len(stack)-1]
            if topIdx == -1 || h < height[topIdx] {
                stack = append(stack, i)
                break
            } else {
                if leftIdx := stack[len(stack)-2]; leftIdx != -1 {
                    water += (min(h, height[leftIdx]) - height[stack[len(stack)-1]]) * (i-leftIdx-1)
                }
                stack = stack[:len(stack)-1]
            }
        }
    }
    return water
}
