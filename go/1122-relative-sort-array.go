import "sort"
func relativeSortArray(arr1 []int, arr2 []int) []int {
    invertIndex := make([]int, 1001)
    for idx, val := range arr2 {
        invertIndex[val] = idx+1
    }
    sort.Slice(arr1, func (i, j int) bool {
        v1, v2 := arr1[i], arr1[j]
        if invertIndex[v1] == 0 && invertIndex[v2] == 0 {
            return v1 < v2
        } else if invertIndex[v1] == 0 {
            return false
        } else if invertIndex[v2] == 0 {
            return true
        } else {
            return invertIndex[v1] < invertIndex[v2]
        }
    })
    return arr1
}
