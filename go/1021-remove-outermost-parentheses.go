// Counter
import "strings"
func removeOuterParentheses(S string) string {
    if len(S) == 0 { return "" }
    var builder strings.Builder
    diff := 0
    var part []byte
    for i := 0; i < len(S); i++ {
        part = append(part, S[i])
        if S[i] == '(' { diff++ } else { diff-- }
        if diff == 0 {
            builder.Write(part[1:len(part)-1])
            part = part[:0]
        }
    }
    return builder.String()
}

// Stack
import "strings"
func removeOuterParentheses(S string) string {
    if len(S) == 0 { return "" }
    var stack []byte
    var builder strings.Builder
    var part []byte
    for i := 0; i < len(S); i++ {
        part = append(part, S[i])
        if S[i] == '(' { stack = append(stack, '(') } else { stack = stack[:len(stack)-1] }
        if len(stack) == 0 {
            builder.Write(part[1:len(part)-1])
            part = part[:0]
        }
    }
    return builder.String()
}
