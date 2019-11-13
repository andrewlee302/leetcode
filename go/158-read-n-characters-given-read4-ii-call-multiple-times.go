/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */
var solution = func(read4 func([]byte) int) func([]byte, int) int {
    // implement read below.
    var buf4 [4]byte
    p, num := 0, 0
    EOF := false
    return func(buf []byte, n int) int {
        if len(buf) | n == 0 { return 0 }
        if p == 0 && EOF { return 0 }
        n = min(len(buf), n)
        start := 0
        for n > 0 {
            if !EOF && num == 0 {
                num = read4(buf4[:])
                p = 0
                if num == 0 {
                    EOF = true
                    return start
                }
            }
            rn := min(n, num)
            copy(buf[start:], buf4[p: p+rn])
            start += rn
            n -= rn
            p += rn
            num -= rn
        }
        return start
    }
}

func min(i, j int) int { if i < j { return i } else { return j } }
