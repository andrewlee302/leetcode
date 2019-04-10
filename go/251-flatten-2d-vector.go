type Vector2D struct {
    v [][]int
    r, c int
}

// [[],[3]]
func Constructor(v [][]int) Vector2D {
    return Vector2D{v, 0, 0}
}


func (this *Vector2D) Next() int {
    this.HasNext()
    v := this.v[this.r][this.c]
    this.c++ // Note.
    return v
}


func (this *Vector2D) HasNext() bool {
    for this.r < len(this.v) { // Note. Can't be: this.r < len(this.v) && this.c < len(this.v[this.r])
        if this.r < len(this.v) && this.c < len(this.v[this.r]) {
            return true
        }
        if this.c == len(this.v[this.r]) {
            this.r++
            this.c = 0
        } else {
            this.c++
        }
    }
    return false
}


/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
