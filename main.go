package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Each cell in the graph is a bitfield representing all directions
// which are "open", as opposed to walls.
//
// | cell bitfield to combine directions
// & cell bitfield to test openness of a particular direction
type Graph = [][]int

// Direction bit values
const (
	N int = 1
	S int = 2
	E int = 4
	W int = 8
)

type offset struct {
	x, y int
}

// Note that N/S offsets are inverse to what you'd expect because the
// postive y-axis moves downwards.
var MovementOffsets = map[int]offset{
	N: {0, -1},
	S: {0, 1},
	E: {1, 0},
	W: {-1, 0},
}

var OppositeDirections = map[int]int{
	N: S,
	S: N,
	E: W,
	W: E,
}

// Recursively visit every unvisited cell starting from (x, y).
func exploreCell(x, y int, g *Graph) {
	dirs := []int{N, S, E, W}
	rand.Shuffle(len(dirs), func(i, j int) {
		dirs[i], dirs[j] = dirs[j], dirs[i]
	})

	var nx, ny int
	for _, dir := range dirs {
		nx = x + MovementOffsets[dir].x
		ny = y + MovementOffsets[dir].y
		if validCell(nx, ny, g) && (*g)[ny][nx] == 0 {
			// Open the direction of travel for the current cell
			(*g)[y][x] |= dir
			// Open the direction of travel relative to the next cell
			(*g)[ny][nx] |= OppositeDirections[dir]

			exploreCell(nx, ny, g)
		}
	}
}

// Return true if the passed (x, y) coord is within the graph dimensions.
func validCell(x, y int, g *Graph) bool {
	return 0 <= x && x < len((*g)[0]) && 0 <= y && y < len(*g)
}

// Output a graph as ASCII to stdout.
func drawMaze(g *Graph) {
	mazeWidth := len((*g)[0]) * 2

	// North border is always closed
	fmt.Printf(" %s\n", strings.Repeat("_", mazeWidth-1))
	for _, row := range *g {
		// West border is always closed
		fmt.Printf("|")

		// For each cell print the East and South borders if closed,
		// no need to print the West and North borders as those
		// are added implicitly by the neighboring cells.
		for i, cell := range row {
			if cell&S != 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("_")
			}

			if cell&E != 0 && row[i+1]&S != 0 {
				fmt.Printf("_")
			} else if cell&E != 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("|")
			}
		}
		fmt.Printf("\n")
	}
}

func newGraph(h, w int) *Graph {
	graph := make(Graph, h)
	for i := range graph {
		graph[i] = make([]int, w)
	}

	return &graph
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var h, w int
	if len(os.Args) < 3 {
		h, w = 10, 10
	} else {
		h, _ = strconv.Atoi(os.Args[1])
		w, _ = strconv.Atoi(os.Args[2])
	}

	graph := newGraph(h, w)

	exploreCell(0, 0, graph)

	drawMaze(graph)
}
