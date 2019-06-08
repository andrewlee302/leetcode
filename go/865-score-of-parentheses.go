// Counter.
func scoreOfParentheses(S string) int {
    diff := 0
    score := 0
    for i := 0; i < len(S); i++ {
        if S[i] == '(' { 
            diff++ 
        } else { 
            if S[i-1] == '(' {
                score += 1 << uint(diff - 1)    
            }
            diff-- 
        }
    }
    return score
}

// Stack.
func scoreOfParentheses(S string) int {
    // Example: (())()
    var stack []int
    stack = append(stack, 0) // the outermost frame. Including two parts: (()), ()
    for i := 0; i < len(S); i++ {
        if S[i] == '(' {
            stack = append(stack, 0)
        } else {
            v := stack[len(stack)-1] // value in the frame.
            v = max(v*2, 1) // value for the frame
            stack[len(stack)-2] += v  // last frame increases
            stack = stack[:len(stack)-1] // last frame is on the top
        }
    }
    return stack[0]
}

func min(a, b int) int { if a < b { return a } else { return b } }
func max(a, b int) int { if a < b { return b } else { return a } }
