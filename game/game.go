package game

import (
	"fmt"
	"time"
)

// Board is a set of points that are 'alive'
type Board map[Point]bool

// Point is a x,y coord on the Board
type Point struct {
	X int
	Y int
}

// Neighbors produces a slice of the nearby cells to Point p.
func Neighbors(p Point) []Point {

	pts := []Point{}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			pts = append(pts, Point{X: p.X + x, Y: p.Y + y})
		}
	}

	// have to remove the point at index 4, since it is the
	// original point itself.
	return append(pts[:4], pts[5:]...)
}

// Advance applies the rules to the given Board and produces a new
// Board that is the next generation.
func Advance(b Board) Board {
	next := make(Board)

	// make a set of cells to check, initialize with all current cells
	check := make(Board)
	for p := range b {
		check[p] = true
	}

	// put all neighbors of actives in check
	for p := range b {
		for _, n := range Neighbors(p) {
			check[n] = true
		}
	}

	//apply rules to all cells to check
	for p := range check {
		// find number of p's neighbors which are alive in current generation
		count := 0
		for _, n := range Neighbors(p) {
			if b[n] {
				count++
			}
		}

		// 3 neighbors-> live or spawn; 2 and currently alive-> live
		if count == 3 || (b[p] && count == 2) {
			next[p] = true
		}
	}

	return next
}

// Show displays a Board on the screen
func Show(b Board) {

	// figure out extents to draw
	high := Point{}
	low := Point{}

	for p := range b {
		if p.X > high.X {
			high.X = p.X
		}
		if p.X < low.X {
			low.X = p.X
		}
		if p.Y > high.Y {
			high.Y = p.Y
		}
		if p.Y < low.Y {
			low.Y = p.Y
		}
	}

	fmt.Println("\033[2J\033[0;0H") // clear screen & move cursor to 0,0

	var output string // faster than multiple prints to screen
	for y := low.Y - 1; y < high.Y+2; y++ {
		for x := low.X - 1; x < high.X+2; x++ {
			if b[Point{x, y}] {
				output += "\u2588"
			} else {
				output += "\u2591"
			}
		}
		output += fmt.Sprintf(" %d\n", y)
	}

	fmt.Println(output) // display all at once w/ blank line on end
}

// Animate performs the given number of `iterations` on the initial `board`
// and displays it on the screen each time, `pause`ing between each. It does
// not alter `board`.
func Animate(board Board, iterations int, pause time.Duration) {
	b := make(Board)
	for p := range board {
		b[p] = true
	}

	for ; iterations > 0; iterations-- {
		Show(b)
		b = Advance(b)
		time.Sleep(pause)
	}
}
