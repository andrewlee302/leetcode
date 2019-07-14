import "math/rand"
type RandomizedCollection struct {
    array []int
    m map[int]map[int]bool
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
    return RandomizedCollection{m: make(map[int]map[int]bool)}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
    this.array = append(this.array, val)
    if _, ok := this.m[val]; !ok { this.m[val] = make(map[int]bool) }
    this.m[val][len(this.array)-1] = true
    return len(this.m[val]) == 1
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
    if pos := this.m[val]; pos != nil && len(pos) > 0 {
        for p, _ := range pos {
            lastIdx := len(this.array)-1
            lastVal := this.array[lastIdx]
            this.array[p] = lastVal
            this.array = this.array[:lastIdx]
            delete(pos, p) // firstly.
            this.m[lastVal][p] = true // secondly.
            delete(this.m[lastVal], lastIdx) // thirdly.
            break
        }
        return true
    } else {
        return false
    }
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
    return this.array[rand.Intn(len(this.array))]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
