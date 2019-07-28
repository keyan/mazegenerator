package main

import (
	"testing"
)

func TestValidCell(t *testing.T) {
	g := newGraph(5, 5)

	v := validCell(0, 0, g)
	if !v {
		t.Fail()
	}

	v = validCell(0, -1, g)
	if v {
		t.Fail()
	}
}

func TestExploreTouchesAllCells(t *testing.T) {
	g := newGraph(5, 5)

	exploreCell(0, 0, g)

	for _, row := range *g {
		for _, cell := range row {
			if cell == 0 {
				t.Fail()
			}
		}
	}
}

func TestWallBitMathJoinDirections(t *testing.T) {
	if !(((N | S | E) & S) == S) {
		t.Fail()
	}

	if ((N | S | E) & W) != 0 {
		t.Fail()
	}
}

func TestWallBitMathCheckDirectionsOpen(t *testing.T) {
	isOpen := (N & N) != 0
	if !isOpen {
		t.Fail()
	}

	isOpen = (N & S) != 0
	if isOpen {
		t.Fail()
	}
}
