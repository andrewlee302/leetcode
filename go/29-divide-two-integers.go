// 10/4 = 2
// 10/3 = 3
const INT32_MAX = int(^uint32(0)>>1)
const INT32_MIN = ^INT32_MAX
func divide(dividend int, divisor int) int {
    if dividend == INT32_MIN && divisor == -1 { return INT32_MAX }
    isPositive := dividend >= 0 && divisor > 0 || dividend < 0 && divisor < 0
    if dividend < 0 { dividend = -dividend }
    if divisor < 0 { divisor = -divisor }
    ret := 0
    for dividend >= divisor {
        shiftDivisor := divisor
        shiftNum := uint(0)
        for dividend >= shiftDivisor {
            ret += 1 << shiftNum
            dividend -= shiftDivisor
            shiftNum++
            shiftDivisor <<= 1
        }
    }
    if isPositive { return ret } else { return -ret }
}
