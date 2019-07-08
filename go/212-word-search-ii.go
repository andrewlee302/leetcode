func findWords(board [][]byte, words []string) []string {
    if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 { return []string{} }
    root := &TrieNode{}
    wl := 0
    for _, word := range words {
        wl = max(wl, len(word))
        curr := root
        for i := 0; i < len(word); i++ {
            offset := word[i]-'a'
            if curr.nodes[offset] == nil { curr.nodes[offset] = &TrieNode{} }
            curr = curr.nodes[offset]
            if i == len(word)-1 { curr.leaf = true }
        }
    }
    buf := make([]byte, wl)
    m := make(map[string]bool)
    moves := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            DFS(board, moves, root, buf, 0, i, j, m)
        }
    }
    ret, cnt := make([]string, len(m)), 0
    for k, _ := range m {
        ret[cnt] = k
        cnt++
    }
    return ret
}

type TrieNode struct {
    nodes [26]*TrieNode
    leaf bool
}

// root cannot be nil.
func DFS(board [][]byte, moves [][]int, root *TrieNode, buf []byte, idx, i, j int, m map[string]bool) {
    root = root.nodes[board[i][j]-'a']
    if root == nil { return }
    buf[idx] = board[i][j]
    if root.leaf { m[string(buf[:idx+1])] = true }
    board[i][j] = '#'
    for _, move := range moves {
        ii, jj := i+move[0], j+move[1]
        if ii >= 0 && ii < len(board) && jj >= 0 && jj < len(board[0]) && board[ii][jj] != '#' {
            DFS(board, moves, root, buf, idx+1, ii, jj, m)
        }
    }
    board[i][j] = buf[idx]
}

func max(i, j int) int { if i > j { return i } else { return j } }
