func isRobotBounded(instructions string) bool {
    moves := [][]int{{0,1}, {1,0}, {0,-1}, {-1,0}} // N, E, S, W
    direction := 0 // N
    x, y := 0, 0
    for i := 0; i < len(instructions); i++ {
        action := instructions[i]
        switch (action) {
            case 'G': {
                x += moves[direction][0]
                y += moves[direction][1]
            }
            case 'L': direction = (direction + 3) % 4
            case 'R': direction = (direction + 1) % 4
        }
    }
    return x == 0 && y == 0 || direction != 0
}
