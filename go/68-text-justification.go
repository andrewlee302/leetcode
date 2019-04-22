import "strings"
func fullJustify(words []string, maxWidth int) []string {
    var builder strings.Builder
    var lineWords []string
    lineWordsLetterCnt := 0
    var ret []string
    for i := 0; i < len(words); {
        for ; i < len(words); i++ {
            lineWords = append(lineWords, words[i])
            lineWordsLetterCnt += len(words[i])
            if lineWordsLetterCnt + len(lineWords) - 1 > maxWidth {
                lineWords = lineWords[:len(lineWords)-1]
                lineWordsLetterCnt -= len(words[i])
                break
            }
        }
        if i < len(words) { 
            space := maxWidth - lineWordsLetterCnt
            if len(lineWords) == 1 {
                builder.WriteString(lineWords[0])
                for k := 0; k < space; k++ { builder.WriteByte(' ') }
            } else {
                evenSpace := space / (len(lineWords) - 1)
                remainder := space % (len(lineWords) - 1)
                for j, word := range lineWords {
                    builder.WriteString(word)
                    if j < remainder {
                        for k := 0; k < evenSpace + 1; k++ { builder.WriteByte(' ') }
                    } else if j != len(lineWords) - 1 {
                        for k := 0; k < evenSpace; k++ { builder.WriteByte(' ') }
                    }
                }
            }
        } else {
            finalSpace := maxWidth - lineWordsLetterCnt - (len(lineWords) - 1)
            for j, word := range lineWords {
                builder.WriteString(word)
                if j != len(lineWords) - 1 {
                    builder.WriteByte(' ')
                } else {
                    for k := 0; k < finalSpace; k++ { builder.WriteByte(' ') }
                }
            }
        }
        ret = append(ret, builder.String())
        builder.Reset()
        lineWords = lineWords[:0]
        lineWordsLetterCnt = 0
    }
    return ret
}
