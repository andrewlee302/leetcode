// Improved from the leftMax & rightMax version (two iterations). O(N), O(1).
func trap(height []int) int {
    water := 0
    left, right := 0, len(height) - 1
    leftMax, rightMax := 0, 0
    // left == right, biggest one. So
    for left <= right {
        leftMax, rightMax = max(leftMax, height[left]), max(rightMax, height[right])
        if leftMax < rightMax {
            // leftMax is real leftMax, real rightMax may be bigger than rightMax.
            // So we're sure leftMax == min(leftMax, real rightMax).
            water += leftMax - height[left]
            left++
        } else {
            water += rightMax - height[right]
            right--
        }
    }
    return water
}

// LeftMax & rightMax. O(N), O(N)
func trap(height []int) int {
    water := 0
    leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
    for i := 0; i < len(height); i++ {
        if i == 0 { leftMax[i] = height[i] } else { leftMax[i] = max(height[i], leftMax[i-1]) }
    }
    for i := len(height)-1; i >= 0; i-- {
        if i == len(height)-1 { rightMax[i] = height[i] } else { rightMax[i] = max(height[i], rightMax[i+1]) }
    }
    for i := 0; i < len(height); i++ {
        water += min(leftMax[i], rightMax[i]) - height[i]
    }
    return water
}

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
