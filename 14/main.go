package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type robot struct {
	x, y, vx, vy int
}

func step(robots []robot, steps, L, H int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for i, r := range robots {
		robots[i].x = (r.x + steps*r.vx) % L
		robots[i].y = (r.y + steps*r.vy) % H
		if robots[i].x < 0 {
			robots[i].x += L
		}
		if robots[i].y < 0 {
			robots[i].y += H
		}

		if robots[i].x < L/2 && robots[i].y < H/2 {
			q1++
		} else if robots[i].x < L/2 && robots[i].y > H/2 {
			q3++
		} else if robots[i].x > L/2 && robots[i].y < H/2 {
			q2++
		} else if robots[i].x > L/2 && robots[i].y > H/2 {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func printRobots(robots []robot, H, L int) {
	grid := make([][]int, H)
	for i := range grid {
		grid[i] = make([]int, L)
	}
	for _, r := range robots {
		grid[r.y][r.x]++
	}
	fmt.Println()
	for _, row := range grid {
		for _, val := range row {
			if val > 0 {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func main() {
	// file, _ := os.ReadFile("test_data")
	file, _ := os.ReadFile("data")

	input := string(file)
	robots := make([]robot, 0)
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			break
		}
		fields := strings.Fields(row)
		point := strings.Split(fields[0][2:], ",")
		velocity := strings.Split(fields[1][2:], ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		vx, _ := strconv.Atoi(velocity[0])
		vy, _ := strconv.Atoi(velocity[1])
		robots = append(robots, robot{
			x: x, y: y, vx: vx, vy: vy,
		})
	}

	// H := 7
	// L := 11
	H := 103
	L := 101
	// steps := 100
	// step(robots, 100)

	N := 0
	numRobots := len(robots)
	// step(robots, N, L, H)
	for n := N; n < 200000; n++ {
		step(robots, 1, L, H)

		x := 0
		y := 0
		for _, r := range robots {
			x += r.x
			y += r.y
		}
		x = x / numRobots
		y = y / numRobots
		xx := 0
		yy := 0
		for _, r := range robots {
			xx += (r.x - x) * (r.x - x)
			yy += (r.y - y) * (r.y - y)
		}

		// clear()
		// printRobots(robots, H, L)
		if xx < 350000 && yy < 350000 {
			printRobots(robots, H, L)
			fmt.Println(n, x, y, xx, yy)
			time.Sleep(time.Millisecond * 50)
		}
	}

}
