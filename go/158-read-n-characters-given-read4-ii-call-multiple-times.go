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
    cache := make([]byte, 4)
    cacheIdx := 0
    remain := 0
    EOF := false
    // implement read below.
    return func(buf []byte, n int) int {
        bufIdx := 0
        for n > 0 {
            // if EOF { break } // wrong
            if remain == 0 {
                if EOF { break }
                remain = read4(cache)
                cacheIdx = 0
                if remain < 4 { EOF = true }
            }
            copyLen := min(remain, n)
            copy(buf[bufIdx:], cache[cacheIdx:cacheIdx+copyLen])
            n -= copyLen
            bufIdx += copyLen
            cacheIdx += copyLen
            remain -= copyLen
        }
        fmt.Println(bufIdx)
        return bufIdx
    }
}

func min(i, j int) int { if i < j { return i } else { return j } }
