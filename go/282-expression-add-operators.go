import "strconv"
func addOperators(num string, target int) []string {
    if len(num) == 0 { return nil }
    var ret []string
    buf := make([]byte, len(num)*2)
    recursion(0, 0, 0, num, 0, buf, 0, target, &ret)
    return ret
}

// lastAddValue = lastLeftAddOperand + lastRightAddOperand.
func recursion(lastAddValue, lastRightAddOperand, buildOperand int, num string, idx int, buf []byte, bufIdx int, target int, ret *[]string) {
    if idx == len(num) {
        if lastAddValue == target && buildOperand == 0 {
            *ret = append(*ret, string(buf[1:bufIdx]))
        }
        return
    }
    digit := int(num[idx]-'0')
    buildOperand = buildOperand * 10 + digit
    if buildOperand != 0 {
        recursion(lastAddValue, lastRightAddOperand, buildOperand, num, idx+1, buf, bufIdx, target, ret)
    }
    s := strconv.Itoa(buildOperand)
    copy(buf[bufIdx+1:], []byte(s))
    buf[bufIdx] = '+'
    recursion(lastAddValue+buildOperand, buildOperand, 0, num, idx+1, buf, bufIdx+len(s)+1, target, ret)
    if bufIdx > 0 {
        buf[bufIdx] = '-'
        recursion(lastAddValue-buildOperand, -buildOperand, 0, num, idx+1, buf, bufIdx+len(s)+1, target, ret)
        buf[bufIdx] = '*'
        lastLeftAddOperand := lastAddValue - lastRightAddOperand
        currRightAddOperand := lastRightAddOperand*buildOperand
        recursion(lastLeftAddOperand+currRightAddOperand, currRightAddOperand, 0, num, idx+1, buf, bufIdx+len(s)+1, target, ret)
    }
}
