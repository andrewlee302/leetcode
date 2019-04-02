import "strings"
import "strconv"
func groupStrings(ss []string) [][]string {
    var ret [][]string
    m := make(map[string][]string)
    var builder strings.Builder
    for _, s := range ss {
        builder.WriteString("#0")
        for i := 1; i < len(s); i++ {
            builder.WriteByte('#')
            diff := (s[i] - s[0] + 26) % 26
            builder.WriteString(strconv.Itoa(int(diff)))
        }
        mapping := builder.String()
        m[mapping] = append(m[mapping], s)
        builder.Reset()
    }
    for _, strs := range m { ret = append(ret, strs) }
    return ret
}
// Note: ["az","ba"]
