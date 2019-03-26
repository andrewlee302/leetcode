func multiply(num1 string, num2 string) string {
    l1, l2 := len(num1), len(num2)
    minLen, littleNum, bigNum := l1, num1, num2
    if l2 < l1 { minLen, littleNum, bigNum = l2, num2, num1 }
    if minLen == 1 && littleNum[0] == '0' { return "0" }
    part := make([][]byte, minLen)
    for i := 0; i < minLen; i++ {
        var flag byte = 0
        littleV := littleNum[minLen-i-1] - '0'
        part[i] = make([]byte, len(bigNum) + 1)
        for j := len(bigNum) - 1; j >= 0; j-- {
            v := (bigNum[j] - '0') * littleV + flag
            flag = v/10
            part[i][j+1] = v%10
        }
        part[i][0] = flag    
    }
    var flag int = 0
    ret := make([]byte, len(bigNum) + len(littleNum) + 1)
    for k := 0; k < len(bigNum) + len(littleNum); k++ {
        sum := flag // sum must be int, much bigger than byte
        for i := 0; i < minLen; i++ {
            idx := len(bigNum) + i - k
            if idx < 0 || idx > len(bigNum) { continue }
            sum += int(part[i][idx])
        }
        flag = sum/10
        ret[len(ret)-k-1] = byte(sum%10) + '0'
    }
    ret[0] = byte(flag) + '0'
    start := 0
    for ; start < len(ret) && ret[start] == '0'; start++ { }
    if start == len(ret) { return "0" } else { return string(ret[start:]) }
}
