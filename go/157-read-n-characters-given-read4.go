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
    return func(buf []byte, n int) int {
        _buf := make([]byte, 4)
        rest := n
        for rest > 0 {
            m := read4(_buf)
            if m == 0 {
                break
            }
            l := rest 
            if m < l {
            	l = m
            }
            copy(buf[n-rest:], _buf[:l])
            rest -= l
        }
        return n-rest
    }
}


