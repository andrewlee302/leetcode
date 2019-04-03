import "math"
func bulbSwitch(n int) int {
    return int(math.Sqrt(float64(n))) // only perfect square will be on.
}
