type Trie struct {
    nodes [26]*Trie
    isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{isEnd:false}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    curr := this
    for i := 0; i < len(word); i++ {
        v := word[i] - 'a'
        if curr.nodes[v] == nil {
            trie := Constructor()
            curr.nodes[v] = &trie
        }
        curr = curr.nodes[v]
        if i == len(word) - 1 {
            curr.isEnd = true
        }
    }
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    curr := this
    for i := 0; i < len(word); i++ {
        v := word[i] - 'a'
        if curr.nodes[v] == nil {
            return false
        }
        curr = curr.nodes[v]
        if i == len(word) - 1 {
            return curr.isEnd == true
        }
    }
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    curr := this
    for i := 0; i < len(prefix); i++ {
        v := prefix[i] - 'a'
        if curr.nodes[v] == nil {
            return false
        }
        curr = curr.nodes[v]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
