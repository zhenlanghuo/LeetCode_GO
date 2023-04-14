package main

import "fmt"

func main() {
	fmt.Println(isRobotBounded("GGLLGG"))
	fmt.Println(isRobotBounded("GG"))
	fmt.Println(isRobotBounded("GL"))

}

func isRobotBounded(instructions string) bool {
	directs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	point := []int{0, 0}
	directIndex := 0
	for _, instruction := range instructions {
		switch instruction {
		case 'G':
			point[0] += directs[directIndex][0]
			point[1] += directs[directIndex][1]
		case 'L':
			directIndex = (directIndex - 1 + 4) % 4
		case 'R':
			directIndex = (directIndex + 1) % 4
		}
	}
	fmt.Println(point, directIndex)

	return !((point[0] != 0 || point[1] != 0) && directIndex == 0)
}
