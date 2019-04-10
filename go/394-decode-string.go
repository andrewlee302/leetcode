import "strconv"
import "strings"
const (
    NUM_TAG = "0"
    STR_TAG = "1"
)
func decodeString(s string) string {
    var stack [][]string
    var builder strings.Builder
    for i := 0; i < len(s); {
        if s[i] >= '0' && s[i] <= '9' {
            start := i
            for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { }
            stack = append(stack, []string{NUM_TAG, s[start:i]})
        } else if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
            start := i
            for ; i < len(s) && (s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z'); i++ { }
            stack = append(stack, []string{STR_TAG, s[start:i]})
        } else if s[i] == ']' {
            str := ""
            for len(stack) != 0 {
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if top[0] == NUM_TAG {
                    repeatedCnt, _ := strconv.Atoi(top[1])
                    for i := 0; i < repeatedCnt; i++ {
                        builder.WriteString(str)
                    }
                    break
                } else {
                    str = top[1] + str
                }
            }
            stack = append(stack, []string{STR_TAG, builder.String()})
            builder.Reset()
            i++
        } else { // s[i] == '['
            i++
        }
    }
    for i := 0; i < len(stack); i++ {
        builder.WriteString(stack[i][1])
    }
    return builder.String()
}
