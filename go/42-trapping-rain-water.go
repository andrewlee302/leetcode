// brute force.
func trap(height []int) int {
    sum := 0
    for i := 1; i < len(height) - 1; i++ {
        maxLeft, maxRight := 0, 0
        for j := i; j >= 0; j-- { if height[j] > maxLeft { maxLeft = height[j] } }
        for j := i; j < len(height); j++ { if height[j] > maxRight { maxRight = height[j] } }
        water := maxLeft - height[i]
        if maxLeft > maxRight { water = maxRight - height[i] }
        sum += water
    }
    return sum
}

// Improved from the leftMax & rightMax version (two iterations).
func trap(height []int) int {
    sum := 0
    left, right := 0, len(height) - 1
    leftMax, rightMax := 0, 0
    // left == right, biggest one. So 
    for left <= right {
        if height[left] > leftMax { leftMax = height[left] }
        if height[right] > rightMax { rightMax = height[right] }
        if leftMax < rightMax {
            // leftMax is real leftMax, real rightMax may be bigger than rightMax.
            // So we're sure min(leftMax, rightMax) = leftMax. 
            sum += leftMax - height[left]
            left++
        } else {
            sum += rightMax - height[right]
            right--
        }
    }
    return sum
}
