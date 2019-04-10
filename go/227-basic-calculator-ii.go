import "strconv"
func calculate(s string) int {
    if len(s) == 0 { return 0 }
    var stack []int
    var operation byte = '+'
    for i := 0; i < len(s); {
        if s[i] >= '0' && s[i] <= '9' {
            start := i
            for ;i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { }
            num, _ := strconv.Atoi(s[start:i])
            switch (operation) {
                case '+': stack = append(stack, num)
                case '-': stack = append(stack, -num)
                case '*': stack[len(stack)-1] *= num
                case '/': stack[len(stack)-1] /= num
            }
        } else if s[i] == ' ' {
            i++
        } else {
            operation = s[i]
            i++
        }
    }
    ret := 0
    for i := 0; i < len(stack); i++ { ret += stack[i] }
    return ret
}
