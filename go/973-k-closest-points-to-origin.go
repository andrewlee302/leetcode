// Iterative.
import "math/rand"
func kClosest(points [][]int, K int) [][]int {
    vs := make([][]int, len(points))
    for i, p := range points {
        vs[i] = []int{p[0]*p[0] + p[1]*p[1], i}
    }
    l, r := 0, len(vs) - 1
    for l <= r {
        pivot := rand.Intn(r-l+1) + l
        vs[r], vs[pivot] = vs[pivot], vs[r]
        p, q := l, l
        for ; q < r; q++ {
            if vs[q][0] < vs[r][0] {
                vs[p], vs[q] = vs[q], vs[p]
                p++
            }
        }
        vs[p], vs[r] = vs[r], vs[p]
        if p == K - 1 || p == K { break } // Note.
        if p > K { r = p - 1} else { l = p + 1 }
    }
    var ret [][]int
    for i := 0; i < K; i++ {
        ret = append(ret, points[vs[i][1]])
    }
    return ret
}

// Recursive.
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
