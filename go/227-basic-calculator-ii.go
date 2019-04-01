func calculate(s string) int {
    if len(s) == 0 { return 0 }
    var stack []int
    var sign byte = '+'
    for i := 0; i < len(s); {
        for ; i < len(s) && s[i] == ' '; i++ { }
        if i >= len(s) { break }
        if s[i] >= '0' && s[i] <= '9' {
            num := 0
            for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { num = num * 10 + int(s[i] - '0') }
            switch sign {
                case '+': stack = append(stack, num)
                case '-': stack = append(stack, -num)
                case '*': stack[len(stack)-1] *= num
                case '/': stack[len(stack)-1] /= num
            }
        } else {
            sign = s[i]
            i++
        }
    }
    sum := 0
    for i := 0; i < len(stack); i++ { sum += stack[i] }
    return sum
}
