import "math/rand"
type RandomizedSet struct {
    set map[int]int
    nums []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{set: make(map[int]int)}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.set[val]; ok {
        return false
    } else {
        this.nums = append(this.nums, val)
        this.set[val] = len(this.nums) - 1
        return true
    }
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if pos, ok := this.set[val]; !ok {
        return false
    } else {
        delete(this.set, val)
        // Note, the one that will be deleted happens to be the last element.
        if pos < len(this.nums) - 1 { 
            lastVal := this.nums[len(this.nums)-1]
            this.nums[pos] = lastVal
            this.set[lastVal] = pos
        }
        this.nums = this.nums[:len(this.nums)-1]
        return true
    }
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
