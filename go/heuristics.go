package main

import (
	"math"
)

func (e *npuzzle) h_manhattan_helper(c chan int, brd []int, val int) {
	index := get_index(brd, val)
	final_index := get_index(e.fbrd, val)

	c <- (abs(index / e.size - final_index / e.size) + abs(index % e.size - final_index % e.size))
}

func (e *npuzzle) h_manhattan(brd []int) int {
	ret := 0
	ln := e.size * e.size
	c := make(chan int)
	for i:=1;i<ln;i++ {
		go e.h_manhattan_helper(c, brd, i)
	}

	for i:=1;i<ln;i++ {
		ret += <- c
	}
	return ret
}

func (e *npuzzle) h_misplaced_tiles (brd []int) int {
	ret := 0
	ln := e.size * e.size
	for i:=0;i<ln;i++ {
		if (brd[i] != 0 && e.fbrd[i] != brd[i]) {
			ret++
		}
	}
	return ret
}



func (e *npuzzle) h_horizontal_conflict(c chan int, brd []int, index int) {
	ret := 0
	a := 0
	b := 0
	for i:=0;i < e.size;i++ {
		a = index * e.size + i
		for j:=i+1; j<e.size; j++ {
			b = index * e.size + j
			if (brd[a] != 0 && brd[b] != 0 && brd[a] == e.fbrd[b] && brd[b] == e.fbrd[a]) {
				ret++;
			}
		}
	}
	c <- ret
}

func (e *npuzzle) h_vertical_conflict(c chan int, brd []int, index int) {
	ret := 0
	a := 0
	b := 0
	for i:=0;i < e.size;i++ {
		a = i * e.size + index
		for j:=i+1; j<e.size; j++ {
			b = j * e.size + index
			if (brd[a] != 0 && brd[b] != 0 && brd[a] == e.fbrd[b] && brd[b] == e.fbrd[a]) {
				ret++;
			}
		}
	}
	c <- ret
}

func (e *npuzzle) h_linear_conflict (brd []int) int {
	ret := 0
	ln := e.size * 2
	c := make(chan int)
	for i:=0;i<e.size;i++ {
		go e.h_horizontal_conflict(c, brd, i)
		go e.h_vertical_conflict(c, brd, i)
	}

	for i:=0;i<ln;i++ {
		ret += <- c
	}
	return ret
}

func (e *npuzzle) h_mlc (brd []int) int {
	return (e.h_manhattan(brd) + 2 * e.h_linear_conflict(brd))
}

func (e *npuzzle) h_euclidian_distance_helper(c chan int, brd []int, val int) {
	index := get_index(brd, val)
	final_index := get_index(e.fbrd, val)

	a := abs(index / e.size - final_index / e.size)
	b := abs(index % e.size - final_index % e.size)

	c <- (int(math.Sqrt(float64(a * a + b * b))))
}

func (e *npuzzle) h_euclidian_distance(brd []int) int {
	ret := 0
	ln := e.size * e.size
	c := make(chan int)
	for i:=1;i<ln;i++ {
		go e.h_euclidian_distance_helper(c, brd, i)
	}

	for i:=1;i<ln;i++ {
		ret += <- c
	}
	return ret
}

func (e *npuzzle) h_diagonal_distance_helper(c chan int, brd []int, val int) {
	index := get_index(brd, val)
	final_index := get_index(e.fbrd, val)

	c <- (max(abs(index / e.size - final_index / e.size), abs(index % e.size - final_index % e.size)))
}

func (e *npuzzle) h_diagonal_distance(brd []int) int {
	ret := 0
	ln := e.size * e.size
	c := make(chan int)
	for i:=1;i<ln;i++ {
		go e.h_diagonal_distance_helper(c, brd, i)
	}

	for i:=1;i<ln;i++ {
		ret += <- c
	}
	return ret
}