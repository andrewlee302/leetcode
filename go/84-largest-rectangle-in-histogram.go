func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    leftLess := make([]int, len(heights))
    rightLess := make([]int, len(heights))
    leftLess[0] = -1
    rightLess[len(heights)-1] = len(heights)
    for i := 1; i < len(heights); i++ {
        k := i - 1
        for ; k >= 0 && heights[k] >= heights[i]; k = leftLess[k] { }
        leftLess[i] = k
    }
    for i := len(heights) - 1; i >= 0 ; i-- {
        k := i + 1
        for ; k < len(heights) && heights[k] >= heights[i]; k = rightLess[k] { }
        rightLess[i] = k
    }
    maxArea := 0
    for i := 0; i < len(heights); i++ {
        if area := heights[i] * (rightLess[i] - leftLess[i] - 1); area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}
