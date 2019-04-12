import "strconv"
func addOperators(num string, target int) []string {
    if len(num) == 0 { return nil }
    buf := make([]byte, 2*len(num)) // including the first one as the guard
    var ret []string
    recursion(num, 0, 0, 0, 0, target, buf, 0, &ret)
    return ret
}

func recursion(num string, idx, preOperand, currOperand, value, target int, buf []byte, bufIdx int, ret *[]string) {
    if idx == len(num) {
        if value == target && currOperand == 0 {
            *ret = append(*ret, string(buf[1:bufIdx]))
        }
        return
    }
    currOperand = currOperand * 10 + int(num[idx] - '0')
    if currOperand != 0 {
        // concat with subsequent digits.
        recursion(num, idx+1, preOperand, currOperand, value, target, buf, bufIdx, ret)
    }
    
    currOperandStr := strconv.Itoa(currOperand)
    copy(buf[bufIdx+1:], []byte(currOperandStr))
    bufIdxAdd := 1 + len(currOperandStr)
    
    // every buf has '+', in the end, discard it.
    // try +
    buf[bufIdx] = '+'
    recursion(num, idx+1, currOperand, 0, value+currOperand, target, buf, bufIdx+bufIdxAdd, ret)
    
    // if idx > 0 { // Note!!!
    if bufIdx > 0 {
        // try -
        buf[bufIdx] = '-'
        recursion(num, idx+1, -currOperand, 0, value-currOperand, target, buf, bufIdx+bufIdxAdd, ret)
        
        // try *
        buf[bufIdx] = '*'
        recursion(num, idx+1, preOperand*currOperand, 0, value-preOperand+preOperand*currOperand, target, buf, bufIdx+bufIdxAdd, ret)
    }
}

// encounter a digit, consider as a number, and use +/-/* before it. Or, keep it as a part of number, digest seqential digits, then as a whole number, then choose +/-/*. Use currOperand to keep it.
// initial: idx: 0, preOperand: 0, currOperand: 0
// do:
// first, idx == 0 just do +, can't do */-. Can continue to digest seqential digits.
// idx > 0: all can do.
// success condition: idx = len(num), currOperand == 0, value == target.

// corner case:
// 1+0#1, only if num[i]*10+currOperand != 0, we do concat.
