const INT32_MAX = int(^uint32(0)>>1)
const INT32_MIN = ^INT32_MAX
func divide(dividend int, divisor int) int {
    if dividend == INT32_MIN && divisor == -1 { return INT32_MAX }
    positive := dividend >= 0 && divisor > 0 || dividend < 0 && divisor < 0 
    if dividend < 0 { dividend = -dividend }
    if divisor < 0 { divisor = -divisor }
    for quotient := 0; ; {
        if dividend < divisor { 
            if positive { return quotient } else { return -quotient }
        }
        shiftDivisor := divisor
        for i := 1; ; i++ {
            shiftDivisor = shiftDivisor << 1
            if dividend < shiftDivisor { 
                quotient += 1 << uint(i - 1)
                dividend -= shiftDivisor >> 1
                break
            }
        }
    }
}
