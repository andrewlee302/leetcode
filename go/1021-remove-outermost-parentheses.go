import "strings"
func removeOuterParentheses(S string) string {
    if len(S) == 0 { return "" }
    var stack []byte
    var parts []string
    var builder strings.Builder
    for i := 0; i < len(S); i++ {
        builder.WriteByte(S[i])
        if S[i] == '(' { stack = append(stack, '(') } else { stack = stack[:len(stack)-1] }
        if len(stack) == 0 {
            parts = append(parts, builder.String())
            builder.Reset()
        }
    }
    for _, part := range parts {
        fmt.Println(part)
        builder.WriteString(part[1:len(part)-1])
    }
    return builder.String()
}
