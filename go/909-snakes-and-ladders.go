import "container/list"

func snakesAndLadders(board [][]int) int {
    N := len(board)
    visited := make([]bool, N*N+1)
    queue := list.New()
    queue.PushBack(1)
    depth := 0
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            v := e.Value.(int)
            queue.Remove(e)
            for vv := v+1; vv <= min(N*N, v+6); vv++ {
                dest := vv
                r, c := decode(N, dest)
                if board[r][c] != -1 { 
                    dest = board[r][c]
                }
                if !visited[dest] {
                    if dest == N*N { return depth + 1 }
                    visited[dest] = true
                    queue.PushBack(dest)
                }
            }
        }
        depth += 1
    }
    return -1
}

func decode(N, v int) (int, int) { 
    var r, c int
    r = (v-1)/N
    if r % 2 == 0 { c = (v-1)%N } else { c = N - 1 - ((v-1)%N) }
    r = N-1-r
    return r, c
}

func min(i, j int) int { if i < j { return i } else { return j } }
