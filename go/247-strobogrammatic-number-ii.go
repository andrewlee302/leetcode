var pairs = map[byte]byte{'0':'0', '1':'1', '8':'8', '6':'9', '9':'6'}
var sameDict = []byte{'0', '1', '8'}

func findStrobogrammatic(n int) []string {
    buff := make([]byte, n)
    var ret []string
    place(buff, 0, &ret)
    return ret
}

func place(buff []byte, idx int, ret *[]string) {
    if idx >= len(buff)/2 {
        if len(buff)%2 == 0 {
            *ret = append(*ret, string(buff))
        } else {
            for _, v := range sameDict { 
                buff[idx] = v
                *ret = append(*ret, string(buff))
            }
        }
        return
    }
    for k, v := range pairs {
        if idx == 0 && k == '0' && len(buff) > 1 { continue }
        buff[idx], buff[len(buff)-idx-1] = k, v
        place(buff, idx+1, ret)
    }
}

// 0, 1, 8, 6-9
// odd, xxx 0/1/8 xxx
// even, xxx xxx
