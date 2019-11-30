import "strings"
func simplifyPath(path string) string {
    var stack []string
    strs := strings.Split(path, "/") // strs[0] == ""
    for i := 1; i < len(strs); i++ {
        if str := strs[i]; str == ".." {
            if len(stack) > 0 { stack = stack[:len(stack)-1] } // pop
        } else if str != "." && str != "" {
            stack = append(stack, str)
        } else { } // ".", "" nothing
    }
    var builder strings.Builder
    for i := 0; i < len(stack); i++ {
        builder.WriteString("/")
        builder.WriteString(stack[i])
    }
    ret := builder.String()
    if ret == "" { return "/" } else { return ret }
}
