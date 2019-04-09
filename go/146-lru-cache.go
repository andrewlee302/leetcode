type Node struct {
    Key, Value int
    Prev, Next *Node
}

type LRUCache struct {
    dummy *Node
    kvs map[int]*Node
    limit int
}


func Constructor(capacity int) LRUCache {
    dummy := &Node{}
    dummy.Prev, dummy.Next = dummy, dummy
    return LRUCache{dummy: dummy, kvs: make(map[int]*Node), limit: capacity}
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.kvs[key]; !ok {
        return -1
    } else {
        this.Prioritize(node)
        return node.Value
    }
}


func (this *LRUCache) Put(key int, value int)  {
    dummy := this.dummy
    if node, ok := this.kvs[key]; !ok {
        node = &Node{Key:key, Value:value, Prev:dummy, Next:dummy.Next}
        this.kvs[key] = node
        if len(this.kvs) > this.limit {
            evictNode := dummy.Prev
            delete(this.kvs, evictNode.Key)
            evictNode.Prev.Next = dummy
            dummy.Prev = evictNode.Prev
        }
        dummy.Next.Prev = node
        dummy.Next = node
    } else {
        node.Value = value
        this.Prioritize(node)
    }
}

func (this *LRUCache) Prioritize(node *Node) {
    dummy := this.dummy
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
    node.Prev, node.Next = dummy, dummy.Next
    dummy.Next.Prev = node
    dummy.Next = node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
