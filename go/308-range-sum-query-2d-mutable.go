// 2D binary indexed tree.
type NumMatrix struct {
    matrix [][]int
    tree [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    tree := make([][]int, len(matrix)+1)
    numMatrix := NumMatrix{matrix, tree}
    if len(matrix) == 0 { return numMatrix }
    for i := 0; i < len(tree); i++ { tree[i] = make([]int, len(matrix[0]) + 1) }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            numMatrix.AddDiff(i, j, matrix[i][j])
        }
    }
    return numMatrix
}

func (this *NumMatrix) AddDiff(row, col, delta int) {
    for r := row+1; r < len(this.tree); r += r & (-r) {
        for c := col+1; c < len(this.tree[0]); c += c & (-c) {
            this.tree[r][c] += delta
        }
    }
}

func (this *NumMatrix) Update(row int, col int, val int)  {
    delta := val - this.matrix[row][col]
    this.matrix[row][col] = val
    this.AddDiff(row, col, delta)
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    // r2_c2 - r2_c1_minus - r1_minus_c2 + r1_minus_c1_minus
    r2_c2, r2_c1_minus, r1_minus_c2, r1_minus_c1_minus := 0, 0, 0, 0
    for r2 := row2+1; r2 > 0; r2 -= r2 & (-r2) {
        for c2 := col2+1; c2 > 0; c2 -= c2 & (-c2) {
            r2_c2 += this.tree[r2][c2]
        }
        for c1_minus := col1; c1_minus > 0; c1_minus -= c1_minus & (-c1_minus) {
            r2_c1_minus += this.tree[r2][c1_minus]
        }
    }
    for r1_minus := row1; r1_minus > 0; r1_minus -= r1_minus & (-r1_minus) {
        for c2 := col2+1; c2 > 0; c2 -= c2 & (-c2) {
            r1_minus_c2 += this.tree[r1_minus][c2]
        }
        for c1_minus := col1; c1_minus > 0; c1_minus -= c1_minus & (-c1_minus) {
            r1_minus_c1_minus += this.tree[r1_minus][c1_minus]
        }
    }
    return r2_c2 - r2_c1_minus - r1_minus_c2 + r1_minus_c1_minus
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
