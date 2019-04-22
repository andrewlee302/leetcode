type StreamChecker struct {
    trie *Trie
    buf []byte
}


func Constructor(words []string) StreamChecker {
    trie := TrieConstructor()
    for _, word := range words {
        trie.Insert(word)    
    }
    return StreamChecker{trie:&trie}
}


func (this *StreamChecker) Query(letter byte) bool {
    this.buf = append(this.buf, letter)
    for i := len(this.buf)-1; i >=0; i-- {
        if isStartWith, isMatch := this.trie.StartsWith(this.buf[i:]); isStartWith {
            if isMatch { return true }
        } else {
            return false
        }
    }
    return false
}


/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

type Trie struct {
    nodes [26]*Trie
    isEnd bool
}


/** Initialize your data structure here. */
func TrieConstructor() Trie {
    return Trie{isEnd:false}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    curr := this
    for i := len(word)-1; i >= 0; i-- {
        v := word[i] - 'a'
        if curr.nodes[v] == nil {
            trie := TrieConstructor()
            curr.nodes[v] = &trie
        }
        curr = curr.nodes[v]
        if i == 0 {
            curr.isEnd = true
        }
    }
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix []byte) (bool, bool) {
    curr := this
    for i := len(prefix) - 1; i >= 0; i-- {
        v := prefix[i] - 'a'
        if curr.nodes[v] == nil {
            return false, false
        }
        curr = curr.nodes[v]
    }
    return true, curr.isEnd == true
}
