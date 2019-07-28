package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

var MovementOffsets = map[int]offset{
	N: {0, 1},
	S: {0, -1},
	E: {1, 0},
	W: {-1, 0},
}

var OppositeDirections = map[int]int{
	N: S,
	S: N,
	E: W,
	W: E,
}

// Recursively visit every unvisited cell starting from the passed (x, y) coord.
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

func drawMaze(g *Graph) {
	for _, row := range *g {
		fmt.Printf("|")
		for _, cell := range row {
			fmt.Printf("%d ", cell)
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

	h, _ := strconv.Atoi(os.Args[1])
	w, _ := strconv.Atoi(os.Args[2])
	graph := newGraph(h, w)

	exploreCell(0, 0, graph)

	drawMaze(graph)
}
