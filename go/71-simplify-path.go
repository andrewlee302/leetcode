import "container/list"
import "strings"
func simplifyPath(path string) string {
    stack := list.New()
    parts := strings.Split(path, "/")
    for _, part := range parts {
        if part == "." || part == "" { continue }
        if part == ".." {
            if stack.Len() >= 1 { stack.Remove(stack.Back()) }
            continue
        } 
        stack.PushBack(part)
        // fmt.Println("x"+part)
    }
    var builder strings.Builder
    if stack.Len() == 0 { return "/" }
    for stack.Len() != 0 {
        e := stack.Front()
        stack.Remove(e)
        v := e.Value.(string)
        builder.WriteString("/")
        builder.WriteString(v)
    }
    return builder.String()
}
