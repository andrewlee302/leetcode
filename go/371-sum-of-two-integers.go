func getSum(a int, b int) int {
    var carry uint = 0
    var ret uint = 0
    var power uint = 1
    ua, ub := uint(a), uint(b)
    for ua != 0 || ub != 0 {
        av, bv := ua & 1, ub & 1
        if av | bv | carry == 1 {
            if av & bv & carry == 1 {
                ret |= power
                carry = 1
            } else if av ^ bv ^ carry == 1 {
                ret |= power
                carry = 0 // Note!
            } else {
                carry = 1
            }
        } 
        power, ua, ub = power << 1, ua >> 1, ub >> 1
    }
    if carry == 1 {
        ret |= power
    }
    return int(ret)
}
