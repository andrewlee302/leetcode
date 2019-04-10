import "container/list"
func isValid(s string) bool {
    m := map[byte]byte{'(':')', '{':'}', '[':']'}
    l := list.New()
    for i := 0; i < len(s); i++ {
        e := l.Back()
        if e != nil && m[e.Value.(byte)] == s[i] {
            l.Remove(e)
        } else {
            l.PushBack(s[i])
        }
    }
    return l.Len() == 0
}
