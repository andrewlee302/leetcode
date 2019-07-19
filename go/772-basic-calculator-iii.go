func calculate(s string) int {
    i := 0
    return subCalculate(s, &i)
}

func subCalculate(s string, p *int) int {
    var stack []int
    var sign byte = '+'
    for i := *p; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' || s[i] == '(' {
            var num int
            if s[i] >= '0' && s[i] <= '9' {
                start := i
                for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { }
                num, _ = strconv.Atoi(s[start:i])
                i--
            } else {
                *p = i+1
                num = subCalculate(s, p)
                i = *p
            }
            switch (sign) {
                case '-': stack = append(stack, -num)
                case '+': stack = append(stack, num)
                case '*': stack[len(stack)-1] *= num
                case '/': stack[len(stack)-1] /= num
            }
        } else if s[i] == ')' {
            *p = i
            break
        } else if s[i] != ' ' {
            sign = s[i]
        }
    }
    val := 0
    for i := 0; i < len(stack); i++ { val += stack[i] }
    return val
}
