// stack
import "strconv"
type Item struct {
    tag int // 0: '(', 1: '+', 2: '-', 3: number
    val int
}

func calculate(s string) int {
    var stack []Item
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '(':
            stack = append(stack, Item{tag:0}) 
        case ')': // just one number on the top.
            stack[len(stack)-2] = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        case '+': stack = append(stack, Item{tag:1}) 
        case '-': 
            // dummy effect.
            if stack[len(stack)-1].tag == 0 {
                stack = append(stack, Item{tag:3, val:0})
            } 
            stack = append(stack, Item{tag:2}) 
        case ' ':
        default: // Number can't be divided by other character, so combine it.
            start := i
            for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { }
            num, _ := strconv.Atoi(s[start:i])
            i-- // Note!
            stack = append(stack, Item{tag:3, val:num})
        }
        if len(stack) >= 3 {
            o, p, q := stack[len(stack)-3], stack[len(stack)-2], stack[len(stack)-1]
            if o.tag == 3 && q.tag == 3 && (p.tag == 1 || p.tag == 2) {
                stack = stack[:len(stack)-3]
                if p.tag == 1 {
                    stack = append(stack, Item{tag:3, val:o.val+q.val})
                } else {
                    stack = append(stack, Item{tag:3, val:o.val-q.val})
                }
            }
        }
    }
    return stack[0].val
}
