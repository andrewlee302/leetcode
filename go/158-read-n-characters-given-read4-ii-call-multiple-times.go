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
    var buffer = make([]byte, 4)
    var pos = 0
    var length = 0 // buffer content length
    // implement read below.
    return func(buf []byte, n int) int {
        rest := n
        for rest > 0 {
            if length == 0 {
                m := read4(buffer[pos:])
                if m == 0 {
                    break
                } else {
                    length = m
                }
            }
            readLen := rest
            if length < readLen {
                readLen = length
            }
            copy(buf[n-rest:], buffer[pos:pos+readLen])
            length -= readLen
            pos += readLen
            if length == 0 {
                pos = 0
            }
            rest -= readLen
        }
        return n-rest
    }
}
