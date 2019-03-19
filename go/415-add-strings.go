func addStrings(num1 string, num2 string) string {
    l := len(num1) + 1
    if len(num2) > len(num1) {
        l = len(num2) + 1
    }
    retBytes := make([]byte, l)
    var flag, b1, b2 byte = 0, 0, 0
    for i := 1; i <= len(num1) || i <= len(num2); i++ {
        if i <= len(num1) { b1 = num1[len(num1)-i] - '0' } else { b1 = 0 }
        if i <= len(num2) { b2 = num2[len(num2)-i] - '0' } else { b2 = 0 }
        v := b1 + b2 + flag
        if v >= 10 { 
            flag = 1
            v -= 10
        } else { flag = 0 }
        retBytes[len(retBytes)-i] = v + '0'
    }
    retBytes[0] = flag + '0'
    if retBytes[0] == '0' {
        return string(retBytes[1:]) 
    } else {
        return string(retBytes)
    }
}
