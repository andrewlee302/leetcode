const INT32_MAX = int(^uint32(0) >> 1)
const INT32_MIN = ^INT32_MAX
func myAtoi(str string) int {
    i := 0
    for ; i < len(str) && str[i] == ' '; i++ { }
    if i == len(str) { return 0 }
    positive := true
    if str[i] == '-' { positive = false }
    if str[i] == '+' || str[i] == '-' { i++ } 
    if i == len(str) || str[i] < '0' || str[i] > '9' { return 0 }
    num := 0
    for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
        if positive && num > (INT32_MAX - int(str[i] - '0')) / 10 { return INT32_MAX }
        if !positive && num > (INT32_MAX + 1 - int(str[i] - '0')) / 10 { return INT32_MIN }
        num = num * 10 + int(str[i] - '0')
    }
    if positive { return num } else { return -num }
}
