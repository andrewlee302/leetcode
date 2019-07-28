// Greedy, O(N).
// like Leetcode 121.
func maxAbsValExpr(arr1 []int, arr2 []int) int {
    fs := [][2]int{{1,1},{1,-1},{-1,1},{-1,-1}}
    ret := 0
    for _, f := range fs {
        minV := f[0]*arr1[0] + f[1]*arr2[0] + 0
        for i := 1; i < len(arr1); i++ {
            v := f[0]*arr1[i] + f[1]*arr2[i] + i
            ret = max(ret, v-minV)
            minV = min(minV, v)
        }
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }
func min(i, j int) int { if i < j { return i } else { return j } }
/*
f(i, j)=
|arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j| = max(
arr1[i] - arr1[j] + arr2[i] - arr2[j] + i - j
arr1[i] - arr1[j] + arr2[j] - arr2[i] + i - j
arr1[j] - arr1[i] + arr2[i] - arr2[j] + i - j
arr1[j] - arr1[i] + arr2[j] - arr2[i] + i - j)
= max(
(arr1[i] + arr2[i] + i) - (arr1[j] + arr2[j] + j)
(arr1[i] - arr2[i] + i) - (arr1[j] - arr2[j] + j)
(-arr1[i] + arr2[i] + i) - (-arr1[j] + arr2[j] + j)
(-arr1[i] - arr2[i] + i) - (-arr1[j] - arr2[j] + j))

f(i, j) = max(g1(i)-g1(j), g2(i)-g2(j), g3(i)-g3(j), g4(i)-g4(j))
*/
