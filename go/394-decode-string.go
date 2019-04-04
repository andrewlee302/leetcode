import "container/list"
import "strings"
import "strconv"
func decodeString(s string) string {
    stack := list.New()
    num := 0
    var builder strings.Builder
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' { 
            num = num * 10 + int(s[i] - '0')
            continue
        }
        if s[i] == '[' { 
            stack.PushBack(strconv.Itoa(num))
            stack.PushBack("[")
            num = 0
        } else if s[i] == ']' {
            var str = ""
            for e := stack.Back(); e != nil; e = stack.Back() {
                stack.Remove(e)
                v := e.Value.(string) // maybe ""
                // fmt.Println("here", e.Value, e.Prev(), stack.Back().Value)
                if v == "[" { break }
                str = v + str
            } 
            e := stack.Back()
            stack.Remove(e)
            cnt, _ := strconv.Atoi(e.Value.(string))
            for k := 0; k < cnt; k++ {
                builder.WriteString(str)    
            }
            fmt.Println(i, e.Value.(string), str, builder.String())
            stack.PushBack(builder.String())
            builder.Reset()
        } else {
            for ; i < len(s) && isValidString(s[i]); i++ {
                builder.WriteByte(s[i])
            }
            stack.PushBack(builder.String()) // maybe ""
            builder.Reset()
            if i != len(s) { i-- }
        }
    }
    for e := stack.Front(); e != nil; e = e.Next() {
        str := e.Value.(string)
        builder.WriteString(str)
    }
    return builder.String()
}

func isDigit(b byte) bool {
    return b >= '0' && b <= '9'
}

func isValidString(b byte) bool {
    return !isDigit(b) && b != ']' && b != '['
}
