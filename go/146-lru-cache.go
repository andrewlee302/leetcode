type LRUCache struct {
    cache map[int]*Node
    capacity int
    head *Node
    tail *Node
}

type Node struct {
    prev *Node
    next *Node
    key int
    val int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{cache:make(map[int]*Node), capacity:capacity}
}

func (this *LRUCache) Get(key int) int {
    if v, ok := this.cache[key]; ok {
        this.update(v)
        return v.val
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int)  {
    if v, ok := this.cache[key]; ok {
        this.update(v)
        v.val = value
    } else {
        if len(this.cache) >= this.capacity {
            delete(this.cache, this.head.key)
            this.head.next.prev = this.tail
            this.tail.next = this.head.next
            this.head = this.head.next
        }
        newNode := &Node{key: key, val:value}
        if this.head == nil {
            this.head, this.tail = newNode, newNode
        } else {
            this.head.prev, this.tail.next = newNode, newNode
        }
        newNode.prev, newNode.next = this.tail, this.head
        this.tail = newNode
        this.cache[key] = newNode
    }
}

func (this *LRUCache) update(v *Node) {
    if v == this.tail {
        return
    }
    v.prev.next, v.next.prev = v.next, v.prev
    if this.head == v {
        this.head = v.next
    }
    // move to tail
    v.next, v.prev = this.head, this.tail
    this.tail.next, this.head.prev = v, v
    this.tail = v
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
