// 1D binary indexed tree.
type NumArray struct {
    nums []int
    tree []int
}

func Constructor(nums []int) NumArray {
    tree := make([]int, len(nums)+1)
    numArray := NumArray{nums:nums, tree: tree}
    for i := 0; i < len(nums); i++ {
        numArray.AddDiff(i, nums[i])
    }
    return numArray
}

func (this *NumArray) AddDiff(i int, delta int)  {
    for i++; i < len(this.tree); i += i & (-i) {
        this.tree[i] += delta
    }
}

func (this *NumArray) Update(i int, val int)  {
    delta := val - this.nums[i]
    this.nums[i] = val
    this.AddDiff(i, delta)
}

func (this *NumArray) SumRange(i int, j int) int {
    pre_i_minus, pre_j := 0, 0
    for ; i > 0; i -= i & (-i) { pre_i_minus += this.tree[i] }
    for j++; j > 0; j -= j & (-j) { pre_j += this.tree[j] }
    return pre_j - pre_i_minus
}
