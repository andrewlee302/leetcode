// length of the final result (non-positive two numbers)
// multiplication: len(num1)+len(num2)-1 or len(num1)+len(num2). (Note: zero is exception.)
// sum: max(len(num1), len(num2)) or max(len(num1), len(num2)) + 1 (Note: no exception.)
func multiply(num1 string, num2 string) string {
    if len(num1) == 1 && num1[0] == '0' || len(num2) == 1 && num2[0] == '0' { return "0" }
    retBytes := make([]byte, len(num1)+len(num2))
    for i := 0; i < len(retBytes); i++ { retBytes[i] = '0' }
    for i := len(num1) - 1; i >= 0; i-- {
        carry := 0 // int
        for j := len(num2) - 1; j >= 0; j-- {
            // represented by int
            accum := int(retBytes[i+j+1] - '0' + (num1[i] - '0') * (num2[j] - '0')) + carry
            retBytes[i+j+1] = byte(accum%10) + '0'
            carry = accum/10
        }
        retBytes[i] = byte(carry) + '0' // Note, carry mustn't be more than 9.
    }
    if retBytes[0] == '0' { return string(retBytes[1:]) } else { return string(retBytes) }
}
