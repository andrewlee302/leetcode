func myPow(x float64, n int) float64 {
    if n == 0 { return 1 } else if n == 1 { return x }
    if n < 0 {
        x = 1/x
        n = -n
    }
    part := myPow(x, n/2)
    if n % 2 == 0 { return part * part } else { return part * part * x}
}
