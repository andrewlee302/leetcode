func romanToInt(s string) int {
    m := map[byte]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
    sum := 0
    for i := 0; i < len(s); {
        b := m[s[i]]
        if i+1<len(s) {
            bNext := m[s[i+1]]
            if bNext > b {
                sum += (bNext - b)
                i += 2
                continue
            } 
        } 
        sum += b
        i++
    }
    return sum
}
