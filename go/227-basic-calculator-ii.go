import "strconv"
func calculate(s string) int {
    var stack []int
    var sign byte = '+'
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            start := i
            for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { }
            num, _ := strconv.Atoi(s[start:i])
            i--
            switch (sign) {
                case '-': stack = append(stack, -num)
                case '+': stack = append(stack, num)
                case '*': stack[len(stack)-1] *= num
                case '/': stack[len(stack)-1] /= num
            }
        } else if s[i] != ' ' {
            sign = s[i]
        }
    }
    val := 0
    for i := 0; i < len(stack); i++ { val += stack[i] }
    return val
}
