func kClosest(points [][]int, K int) [][]int {
    sort(points, 0, len(points) - 1, K)
    return points[:K]
}

// K < right - left + 1
func sort(points [][]int, left, right, K int) {
    if left >= right {
        return
    }
    pivotDist := dist(points[right])
    leftPos := left
    for i := left; i < right; i++ {
        if dist(points[i]) < pivotDist {
            points[i], points[leftPos] = points[leftPos], points[i]
            leftPos++
        }
    }
    points[leftPos], points[right] = points[right], points[leftPos]
    leftLen := leftPos - left + 1
    if leftLen == K || leftLen == K + 1 {
        return
    } else if leftLen > K + 1 {
        sort(points, left, leftPos-1, K)
    } else {
        sort(points, leftPos + 1, right, K-leftLen)
    }
}


func dist(p []int) int {
    return p[0] * p[0] + p[1] * p[1]
}
