/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 * 
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */
import "strconv"
func cleanRoom(robot *Robot) {
    moves := [][]int{{-1,0},{0,1},{1,0},{0,-1}} // clockwise: up, right, down, left.
    visited := make(map[string]bool)
    helper(robot, visited, moves, 0, 0, 0)
}

// relatvie position and direction tag.
func helper(robot *Robot, visited map[string]bool, moves [][]int, row, col, d int) {
    idStr := strconv.Itoa(row) + "," + strconv.Itoa(col)
    visited[idStr] = true
    robot.Clean()
    for i := 0; i < 4; i++ {
        nextR, nextC := row+moves[d][0], col+moves[d][1]
        nextIdStr := strconv.Itoa(nextR) + "," + strconv.Itoa(nextC)
        if !visited[nextIdStr] && robot.Move() {
            helper(robot, visited, moves, nextR, nextC, d)
            back(robot)
        }
        robot.TurnRight() // the last one is to come back to the original direction.
        d = (d+1)%4
    }
}

// retreat and recover the direction.
func back(robot *Robot) {
    robot.TurnRight()
    robot.TurnRight()
    robot.Move()
    robot.TurnRight()
    robot.TurnRight()
}
