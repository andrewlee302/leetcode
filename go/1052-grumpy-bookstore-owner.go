func maxSatisfied(customers []int, grumpy []int, X int) int {
    var max int
    total := 0
    sum := 0
    for i := 0; i < len(grumpy); i++ {
        if grumpy[i] == 1 { sum += customers[i] } else { total += customers[i] }
        if i == X - 1 { max = sum }
        if i >= X {
            if grumpy[i-X] == 1 { sum -= customers[i-X] }
            if sum > max { max = sum }
        }
    }
    return max + total
}
