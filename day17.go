package main

import "fmt"

// this took way too long
// the path planner was much too slow (very slow for part b)
// it fails on test17b.txt (reason unknown)
// but moving on

func _traverse_grid(start,end Coord, grid[][]int, max_step int) int {
	width, height := len(grid[0]), len(grid)

	type CoordCompassCount struct {
		CoordCompass
		count int // number of steps in the given dir
	}

	type step struct {
		pos      CoordCompass
		prev     CoordCompass
		cost     int
		num_step int // number of steps in the given dir
	}

	todo := []step{{pos: CoordCompass{Coord: start, dir: East}, cost: 0, num_step: 0}}
	// visited location: their previous point & cost
	type visit_hist struct {
		prev CoordCompass
		cost int
	}
	visited := make(map[string]visit_hist)
	//var last_loc CoordCompass

	all_directions := []Compass{North, East, South, West}

	for len(todo) > 0 {
		// lowest cost
		bidx, bcost := 0, todo[0].cost
		for idx, t := range todo {
			if t.cost < bcost {
				bidx, bcost = idx, t.cost
			}
		}
		current := todo[bidx]
		// remove (not preserving ordering)
		todo[bidx] = todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		// lookup for visited, must include steps as a quicker route which cannot do the steps is not as good
		key:=fmt.Sprintf("%d:%d:%d:%d", current.pos.x, current.pos.y, current.pos.dir, current.num_step)
		// if there is a cheaper cost: skip
		visit, found := visited[key]
		if found && visit.cost < current.cost {
			continue
		}
		// add loc to visited
		visited[key] = visit_hist{prev: current.prev, cost: current.cost}
		// if arrived: end (doing here so end is in visited)
		if current.pos.Coord == end {
			fmt.Printf("found: %v\n", current)
			//last_loc = current.pos
			return current.cost
			//break
		}
		// process loc, all all possible directions
		for _, dir := range all_directions {
			// cannot 180
			if dir == OppositeCompass(current.pos.dir) {
				continue
			}
			// cannot keep going for X in the same direction
			if dir == current.pos.dir && max_step == current.num_step {
				continue
			}
			// cannot move off map
			newstep := step{pos: MakeMovedCoordCompass(current.pos.Coord, dir), prev: current.pos}
			if newstep.pos.x < 0 || newstep.pos.x >= width || newstep.pos.y < 0 || newstep.pos.y >= height {
				continue
			}
			if dir == current.pos.dir {
				newstep.num_step = current.num_step + 1
			} else {
				newstep.num_step = 1
			}
			newstep.cost = current.cost + grid[newstep.pos.y][newstep.pos.x]
			todo = append(todo, newstep)
		}
		//fmt.Printf("todo %d visited %d\n", len(todo), len(visited))
	}

	// path is a little difficult, skip for now
	// make the path
	//path := []CoordCompass{last_loc}
	/*for last_loc.Coord != start {
		key:=
		visit, _ := visited[last_loc]
		path = append([]CoordCompass{visit.prev}, path...)
		last_loc = visit.prev
	}

	debug:=true
	if debug {
		for y,l:=range grid {
			for x:=range l{
				dir:=" "
				// see if a path exits
				for _,p:=range path {
					if p.x==x && p.y==y {
						switch p.dir {
						case North:
							dir="^"
						case East:
							dir=">"
						case South:
							dir="v"
						case West:
							dir="<"
						}
					}
				}
				// find cost
				cost:=99
				var p CoordCompass
				p.x=x
				p.y=y
				for _,d:=range all_directions {
					p.dir=d
					if cst,found:=visited[p]; found{
						cost=Min(cost, cst.cost)
					}
				}
				fmt.Printf("%02d%s ",cost,dir)
			}
			fmt.Printf("\n")
		}
	}*/

	//return path
	return -1
}

func traverse_grid(start,end Coord, grid[][]int, max_step int, min_step int) int {
	width, height := len(grid[0]), len(grid)

	type CoordCompassCount struct {
		CoordCompass
		count int // number of steps in the given dir
	}

	type CoordCompassCountCost struct {
		CoordCompassCount
		cost int // cost of step
	}

	type CoordHistory struct {
		// CoordCompassCount current
		prev CoordCompassCount
		cost int // current cost
	}

	var start_step CoordCompassCountCost
	start_step.Coord=start
	start_step.dir=East
	start_step.cost=0
	start_step.count=0

	todo := make(map[CoordCompassCountCost]bool)
	todo[start_step]=true

	// visited is CoordCompassCount(ignore the cost, count matters) => prev step+ others
	visited := make(map[CoordCompassCount]CoordHistory)

	all_directions := []Compass{North, East, South, West}

	for count:=0;len(todo) > 0;count++ {
		// lowest cost
		current := CoordCompassCountCost{cost: 1000000}
		for t := range todo {
			if t.cost < current.cost {
				current= t
			}
		}
		// remove (not preserving ordering)
		delete(todo, current)
		// if there is a cheaper cost: skip
		visit, found := visited[current.CoordCompassCount]
		if found && visit.cost < current.cost {
			continue
		}
		// add loc to visited
		visited[current.CoordCompassCount] = CoordHistory{prev: current.CoordCompassCount, cost: current.cost}
		// if arrived: end (doing here so end is in visited)
		if current.Coord == end {
			fmt.Printf("found: %v\n", current)
			//last_loc = current.pos
			return current.cost
			//break
		}
		// process loc, all all possible directions
		for _, dir := range all_directions {
			// cannot 180
			if dir == OppositeCompass(current.dir) {
				continue
			}
			// cannot keep going for X in the same direction
			if dir == current.dir && max_step == current.count {
				continue
			}
			// needs to keep going for a certain number of steps
			if dir != current.dir && current.count < min_step {
				continue
			}
			// cannot move off map
			newpos:=MakeMovedCoordCompass(current.Coord, dir)
			if newpos.x < 0 || newpos.x >= width || newpos.y < 0 || newpos.y >= height {
				continue
			}
			newstep := CoordCompassCountCost{}
			newstep.CoordCompass=newpos
			if dir == current.dir {
				newstep.count = current.count + 1
			} else {
				newstep.count = 1
			}
			newstep.cost = current.cost + grid[newstep.y][newstep.x]
			todo[newstep]=true
		}
		if count%1000 ==0{
			fmt.Printf("todo %d visited %d\n", len(todo), len(visited))
		}
	}

	// path is a little difficult, skip for now
	return -1
}

func day17a(filename string) {
	grid, err := LoadIntGrid(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	//for _,g:=range grid {
	//	fmt.Printf("%v\n",g)
	//}

	width,height:=len(grid[0]),len(grid)
	fmt.Printf("Grid size %d x %d\n",width,height)

	path:=traverse_grid(Coord{0,0},Coord{width-1,height-1}, grid, 3,0)
	fmt.Printf("result %v\n",path)
}

func day17b(filename string) {
	grid, err := LoadIntGrid(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	width,height:=len(grid[0]),len(grid)
	fmt.Printf("Grid size %d x %d\n",width,height)

	path:=traverse_grid(Coord{0,0},Coord{width-1,height-1}, grid, 10,4)
	fmt.Printf("result %v\n",path)
}