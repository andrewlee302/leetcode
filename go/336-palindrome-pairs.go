// Trie tree + reverse search.
type TrieNode struct {
    nodes [26]*TrieNode
    idx int // corresponding index of word in words.
    match map[int]bool // indices of word whose rest part is a palindrome.
}

// word could be "", but is unique.
func palindromePairs(words []string) [][]int {
    var ret [][]int
    root := &TrieNode{idx:-1, match:make(map[int]bool)}
    for i, word := range words { addWord(root, word, i) }
    for i, word := range words { 
        matchSet := searchWord(root, word)
        for idx, _ := range matchSet {
            if idx != i { ret = append(ret, []int{idx, i}) }
        }
    }
    return ret
}

// Check the word could be the right part of a palindrome.
// E.g. "lls": "lls"+"", "ll"+"s", "l"+"ls", ""+"lls".
func searchWord(root *TrieNode, word string) map[int]bool {
    ret := make(map[int]bool)
    for i := len(word)-1; i>=0; i-- {
        if root.idx != -1 && isPalindrome(word[:i+1]) {
            ret[root.idx] = true
        }
        offset := word[i]-'a'
        if root.nodes[offset] == nil {
            return ret
        }
        root = root.nodes[offset]
    }
    for k, _ := range root.match {
        ret[k] = true
    }
    return ret
}

// Store the word and tag the palindrome property of the rest of the word. 
// E.g. "abcc": ""+"abcc", "a"+"bcc", "ab"+"cc"(true), "abc"+"c"(true), "abcc"+""(true)
func addWord(root *TrieNode, word string, idx int) {
    for i := 0; i < len(word); i++ {
        if isPalindrome(word[i:]) {
            root.match[idx] = true
        }
        offset := word[i]-'a'
        if root.nodes[offset] == nil {
            root.nodes[offset] = &TrieNode{idx:-1, match: make(map[int]bool)}  
        }
        root = root.nodes[offset]
    }
    root.idx = idx
    root.match[idx] = true // "" is the rest of any string, and is also a palindrome.
}

func isPalindrome(s string) bool {
    for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
        if s[l] != s[r] {
            return false
        }
    }
    return true
}
