package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	DishRock = 'O'
	DishClear = '.'
)

// part A I took a chance on
// part B will need 4 billion mutations. I doubt the string version can cope
// changed to grid & tested, it worked with the basic code in 5 mins
// adding memorisation, it needed a bit of work to consider settling time
// but worked within 1 second

func grid_set(grid []string, x,y int, c rune) {
	l := grid[y]
	grid[y] = l[:x] + string(c) + l[x+1:]
}

func move_north(grid []string) {
	// move all rocks north
	for y := range grid {
		if y == 0 {
			continue // y0 will not move up
		}
		for x := range grid[0] {
			if grid[y][x] == 'O' && grid[y-1][x] == '.' {
				// move up until no space
				final_y := y - 1
				for final_y > 0 && grid[final_y-1][x] == '.' {
					final_y--
				}
				//fmt.Printf("moving %d,%d to %d,%d\n",x,y,x,final_y)
				grid_set(grid, x, y, '.')
				grid_set(grid, x, final_y, 'O')
			}
		}
	}
}

func move_cycle(grid [][]rune) {
	// move all rocks north,west,south,east in that order
	// changing to a 2d grid so we don't need to string slice
	width := len(grid[0])
	height := len(grid)

	// north
	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' && grid[y-1][x] == DishClear {
				// move up until no space
				final_y := y - 1
				for final_y > 0 && grid[final_y-1][x] == DishClear {
					final_y--
				}
				//fmt.Printf("moving %d,%d to %d,%d\n",x,y,x,final_y)
				grid[y][x] = DishClear
				grid[final_y][x] = DishRock
			}
		}
	}
	// west
	for x := 1; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[y][x] == 'O' && grid[y][x-1] == DishClear {
				// move until no space
				final_x := x - 1
				for final_x > 0 && grid[y][final_x-1] == DishClear {
					final_x--
				}
				grid[y][x] = DishClear
				grid[y][final_x] = DishRock
			}
		}
	}
	// south
	for y := height - 2; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' && grid[y+1][x] == DishClear {
				// move until no space
				final_y := y + 1
				for final_y < height-1 && grid[final_y+1][x] == DishClear {
					final_y++
				}
				grid[y][x] = DishClear
				grid[final_y][x] = DishRock
			}
		}
	}
	// east
	for x := width - 2; x >= 0; x-- {
		for y := 0; y < height; y++ {
			if grid[y][x] == 'O' && grid[y][x+1] == DishClear {
				// move until no space
				final_x := x + 1
				for final_x < width-1 && grid[y][final_x+1] == DishClear {
					final_x++
				}
				grid[y][x] = DishClear
				grid[y][final_x] = DishRock
			}
		}
	}
}

func calculate_load(grid []string) int {
	height:=len(grid)
	load:=0
	for y,l:=range grid {
		num:=strings.Count(l,"O")
		load+=num*(height-y)
	}
	return load
}

func day14a(filename string) {
	grid, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	for _, g := range grid {
		println(g)
	}
	fmt.Printf("\n\nmoving north\n")
	move_north(grid)
	for _, g := range grid {
		println(g)
	}
	fmt.Printf("total load %d\n", calculate_load(grid))
}

func day14b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// to 2d array
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	total_moves := 1000 * 1000 * 1000
	cache := make(map[string]int)
	cy1, cy2 := 0, 0
	// move
	start_time := time.Now()
	for cycle := 0; cycle < total_moves; cycle++ {
		g := ""
		for _, l := range grid {
			g += string(l)
		}
		if val, found := cache[g]; found {
			taken := time.Now().Sub(start_time).Seconds()
			fmt.Printf("found match %d matches %d in %.2f seconds\n", cycle, val, taken)
			cy1, cy2 = cycle, val
			break
		}
		cache[g] = cycle

		move_cycle(grid)

		if cycle%1000000 == 0 {
			taken := time.Now().Sub(start_time).Seconds()
			fmt.Printf("cycle %d time %.2f\n", cycle, taken)
		}
	}

	if cy1 > 0 && cy2 > 0 {
		// we know cy1 maps to cy2
		delta := cy1 - cy2
		cycle := total_moves % delta
		// we cannot use stuff before cy2 as its still settling
		for cycle < cy2 {
			cycle += delta
		}
		fmt.Printf("looking for the state %d\n", cycle)
		// find it
		for g, c := range cache {
			if c == cycle {
				// we need to get the score from g
				width := len(grid[0])
				height := len(grid)
				var final []string
				for y := 0; y < height; y++ {
					idx := width * y
					final = append(final, g[idx:idx+width])
				}

				fmt.Printf("total load %d\n", calculate_load(final))
				return
			}
		}
	}
	fmt.Printf("error\n")
}
