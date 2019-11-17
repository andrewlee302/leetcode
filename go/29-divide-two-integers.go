const INT32_MAX = int((^uint32(0)) >> 1)
const INT32_MIN = ^INT32_MAX

func divide(dividend int, divisor int) int {
    if dividend == INT32_MIN && divisor == -1 { return INT32_MAX }
    flag := true
    if dividend < 0 { dividend, flag = -dividend, !flag }
    if divisor < 0 { divisor, flag = -divisor, !flag }
    ret := 0
    for dividend >= divisor {
        shiftDivisor := divisor
        for i := 0; dividend >= shiftDivisor; i++ {
            ret += (1 << uint(i))
            dividend -= shiftDivisor
            shiftDivisor <<= 1
        }
    }
    if flag { return ret } else { return -ret }
}
