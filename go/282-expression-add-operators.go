import "strconv"
func addOperators(num string, target int) []string {
    if len(num) == 0 {
        return nil
    }
    buff := make([]byte, len(num)*2) // including the guard
    ret := []string{}
    tryOperators(buff, 0, 0, 0, 0, 0, num, target, &ret)
    return ret
}
// idx <= len(num)
// buffIdx <= len(buff)
func tryOperators(buff []byte, buffIdx, value, preOperand, currentOperand, idx int, num string, target int, output *[]string)  {
    if idx == len(num) {
        if value == target && currentOperand == 0 {
            *output = append(*output, string(buff[1:buffIdx]))
        }
        return
    }
    currentOperand = currentOperand * 10 + int(num[idx] - '0')
    if currentOperand != 0 {
        tryOperators(buff, buffIdx, value, preOperand, currentOperand, idx+1, num, target, output)
    }
    operandBytes := []byte(strconv.Itoa(currentOperand))
    // plus
    buff[buffIdx] = '+'
    copy(buff[buffIdx+1:], operandBytes)
    tryOperators(buff, buffIdx+len(operandBytes)+1, value+currentOperand, currentOperand, 0, idx+1, num, target, output)
    if buffIdx > 0 {
        // minis
        buff[buffIdx] = '-'
        copy(buff[buffIdx+1:], operandBytes)
        tryOperators(buff, buffIdx+len(operandBytes)+1, value-currentOperand, -currentOperand, 0, idx+1, num, target, output)
        // multiply
        buff[buffIdx] = '*'
        copy(buff[buffIdx+1:], operandBytes)
        tryOperators(buff, buffIdx+len(operandBytes)+1, value-preOperand+currentOperand*preOperand, currentOperand*preOperand, 0, idx+1, num, target, output)
    }
    
}
