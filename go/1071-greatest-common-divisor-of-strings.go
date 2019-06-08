func gcdOfStrings(str1 string, str2 string) string {
    if len(str1) > len(str2) { str1, str2 = str2, str1 }
    if len(str1) == 0 { return str2 }
    if len(str1) == len(str2) && str1 == str2 {
        return str1
    }
    cnt := 0
    for strings.HasPrefix(str2, str1) {
        cnt++
        str2 = str2[len(str1):]
    }
    if cnt == 0 { return "" }
    return gcdOfStrings(str2, str1)
}
