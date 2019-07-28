func alphabetBoardPath(target string) string {
    var retBytes []byte
    currR, currC := 0, 0
    for i := 0; i < len(target); i++ {
        b := int(target[i] - 'a')
        r, c := b/5, b%5
        var action byte
        diffR, diffC := abs(r-currR), abs(c-currC)
        if r != 5 {
            if r > currR { action = 'D' } else { action = 'U' }
            for j := 0; j < diffR; j++ { retBytes = append(retBytes, action) }
            if c > currC { action = 'R' } else { action = 'L' }
            for j := 0; j < diffC; j++ { retBytes = append(retBytes, action) }
        } else {
            if c > currC { action = 'R' } else { action = 'L' }
            for j := 0; j < diffC; j++ { retBytes = append(retBytes, action) }
            if r > currR { action = 'D' } else { action = 'U' }
            for j := 0; j < diffR; j++ { retBytes = append(retBytes, action) }
        }
        retBytes = append(retBytes, '!')
        currR, currC = r, c 
    }
    return string(retBytes)
}

func abs(i int) int { if i >= 0 { return i } else { return -i } }
