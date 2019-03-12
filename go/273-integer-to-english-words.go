import "strings"

var LESS_THAN_20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var TENS = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var TAGS = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    var ret = ""
    for i := 0; i < 4; i++{
        lowerNumLtK := num % 1000
        num = num / 1000
        if lowerNumLtK > 0 {
            if TAGS[i] != "" {
                if ret != "" {
                    ret = numberLtKToWords(lowerNumLtK) + " " + TAGS[i] + " " + ret
                } else {
                    ret = numberLtKToWords(lowerNumLtK) + " " + TAGS[i]
                }
            } else {
                 ret = numberLtKToWords(lowerNumLtK)   
            } 
        }
    }
    return ret
}

// 0 < num < 1000
func numberLtKToWords(num int) string {
    if num < 20 {
        return LESS_THAN_20[num]
    }
    if num < 100 {
        if num % 10 != 0 {
            return TENS[num/10] + " " + LESS_THAN_20[num%10]    
        } else {
            return TENS[num/10]
        }
    }
    if num%100 != 0 {
        return LESS_THAN_20[num/100] + " Hundred " + numberLtKToWords(num%100)
    } else {
        return LESS_THAN_20[num/100] + " Hundred"
    }   
}
