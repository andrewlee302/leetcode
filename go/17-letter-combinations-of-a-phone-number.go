var m = map[byte]string{'2':"abc",'3':"def",'4':"ghi",'5':"jkl",'6':"mno",'7':"pqrs",'8':"tuv",'9':"wxyz"}
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    output := make([]string, 0, 10)
    buff := make([]byte, len(digits))
    backtracking(0, buff, digits, &output)
    return output
}

func backtracking(k int, buff []byte, digits string, output *[]string) {
    if k == len(digits) {
        *output = append(*output, string(buff))
        return
    }
    ss := m[digits[k]]
    
    for i := 0; i < len(ss); i++ {
        buff[k] = ss[i]
        backtracking(k+1, buff, digits, output)
    }
}
