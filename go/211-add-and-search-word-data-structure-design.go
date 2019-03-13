type Trie struct {
    nodes [26]*Trie
    isEnd bool
}

func TrieConstructor() Trie {
    return Trie{isEnd:false}
}

func (this *Trie) Insert(word string)  {
    curr := this
    for i := 0; i < len(word); i++ {
        v := word[i] - 'a'
        if curr.nodes[v] == nil {
            trie := TrieConstructor()
            curr.nodes[v] = &trie
        }
        curr = curr.nodes[v]
        if i == len(word) - 1 {
            curr.isEnd = true
        }
    }
}

// backtracking & support "."
func (this *Trie) Search(word string) bool {
    if len(word) == 0 {
        return this.isEnd
    }
    if word[0] == '.' {
        isMatch := false
        for _, node := range this.nodes {
            if node != nil {
                isMatch = isMatch || node.Search(word[1:])
            }
        }
        return isMatch
    } 
    v := word[0] - 'a'
    if this.nodes[v] == nil {
            return false
    }
    return this.nodes[v].Search(word[1:])  
}

type WordDictionary struct {
    trie Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{trie:TrieConstructor()}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    this.trie.Insert(word)
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    return this.trie.Search(word)
}



/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
