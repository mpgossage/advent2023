package main

import (
	"fmt"
)

// overall an easy day
// grids & path movement

func laser_path(grid []string, start CoordCompass) int {
	width, height := len(grid[0]), len(grid)
	todo := []CoordCompass{start}
	visited := make(map[CoordCompass]bool)

	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		// if off screen ignore
		if p.x < 0 || p.y < 0 || p.x >= width || p.y >= height {
			continue
		}
		// if already explored
		if _, found := visited[p]; found {
			continue
		}
		// mark explored
		visited[p] = true
		// process
		cell := grid[p.y][p.x]
		if cell == '.' {
			// empty just move:
			p.Coord = MoveCoordCompass(p.Coord, p.dir)
			todo = append(todo, p)
		} else if cell == '\\' {
			// direction change
			switch p.dir {
			case North:
				p.dir = West
			case East:
				p.dir = South
			case South:
				p.dir = East
			case West:
				p.dir = North
			}
			p.Coord = MoveCoordCompass(p.Coord, p.dir)
			todo = append(todo, p)
		} else if cell == '/' {
			// direction change
			switch p.dir {
			case North:
				p.dir = East
			case East:
				p.dir = North
			case South:
				p.dir = West
			case West:
				p.dir = South
			}
			p.Coord = MoveCoordCompass(p.Coord, p.dir)
			todo = append(todo, p)
		} else if cell == '|' {
			if p.dir == East || p.dir == West {
				todo = append(todo, CoordCompass{dir: North, Coord: MoveCoordCompass(p.Coord, North)})
				todo = append(todo, CoordCompass{dir: South, Coord: MoveCoordCompass(p.Coord, South)})
			} else {
				p.Coord = MoveCoordCompass(p.Coord, p.dir)
				todo = append(todo, p)
			}
		} else if cell == '-' {
			if p.dir == North || p.dir == South {
				todo = append(todo, CoordCompass{dir: West, Coord: MoveCoordCompass(p.Coord, West)})
				todo = append(todo, CoordCompass{dir: East, Coord: MoveCoordCompass(p.Coord, East)})
			} else {
				p.Coord = MoveCoordCompass(p.Coord, p.dir)
				todo = append(todo, p)
			}
		}
		//fmt.Printf("todo %d coverage %v\n",len(todo), len(visited))
	}

	//fmt.Printf("completed todo %d coverage %v\n",len(todo), len(visited))
	//fmt.Printf("%v\n",visited)

	coverage := make(map[Coord]bool)
	for c := range visited {
		coverage[c.Coord] = true
	}
	//fmt.Printf("total visited %d\n",len(coverage))

	return len(coverage)
}

func day16a(filename string) {
	grid, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	total := laser_path(grid, CoordCompass{Coord{0, 0}, East})
	fmt.Printf("total visited %d\n", total)
}

func day16b(filename string) {
	grid, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	best := 0
	width, height := len(grid[0]), len(grid)
	for x := 0; x < width; x++ {
		best = Max(best, laser_path(grid, CoordCompass{Coord{x, 0}, South}))
		best = Max(best, laser_path(grid, CoordCompass{Coord{x, height - 1}, North}))
	}
	for y := 0; y < height; y++ {
		best = Max(best, laser_path(grid, CoordCompass{Coord{0, y}, East}))
		best = Max(best, laser_path(grid, CoordCompass{Coord{width - 1, y}, West}))
	}
	fmt.Printf("best visited %d\n", best)
}
