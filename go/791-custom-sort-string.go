import "strings"
func customSortString(S string, T string) string {
    var count [26]int
    for i := 0; i < len(T); i++ { count[T[i]-'a']++ }
    var builder strings.Builder
    for i := 0; i < len(S); i++ {
        for j := 0; j < count[S[i]-'a']; j++ {
            builder.WriteByte(S[i])
        }
        count[S[i]-'a'] = 0
    }
    for i := 0; i < 26; i++ { 
        for j := 0; j < count[i]; j++ { builder.WriteByte(byte(i)+'a') }
    }
    return builder.String()
}
