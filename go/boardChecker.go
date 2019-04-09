package main

import (
	"fmt"
)

func (e *npuzzle) setFinalState() [][]int {
	var val, side, start int
	brd := make([][]int, e.size)
	for i, _ := range brd {
		brd[i] = make([]int, e.size)
	}
	val = 1
	side = e.size - 1
	for side > 0 {
		end := e.size - start - 1
		for is:=0; is<4; is++ {
			for i:=0; i<side; i++ {
				crd := [][]int{{start, start + i}, {start + i, end}, {end, end - i}, {end - i, start}} 
				brd[crd[is][0]][crd[is][1]] = val
				val++
			}
		}
		side -= 2
		start++
	}
	if e.size % 2 == 0 {
		brd[e.size/2][e.size/2-1] = 0
	}
	return brd
}

func (e *npuzzle) checkValues(board [][]int, size int) {
	ntiles := size * size
	checker := make([]int, ntiles)

	for y, raw := range board {
		for  x, val := range raw {
			if (val < 0) || (val > ntiles) {
				putError(fmt.Sprintf("wrong value at [%d:%d]", y, x))	
			} else {
				checker[val] += 1
				if checker[val] > 1 {
					putError(fmt.Sprintf("duplicate Value at [%d:%d]", y, x))
				}
			}
		}
	}
}

func (e *npuzzle) getInversions(flat_board []int) int {
	sum := 0
	tmp := 0

	for i := 0; i < e.size * e.size; i++ {
		tmp = 0
		for n := i + 1; n < e.size * e.size; n++ {
			if flat_board[n] != 0 && flat_board[i] > flat_board[n] {
				tmp += 1
			}
		}
		sum += tmp
	}
	return sum
}

func (e *npuzzle) checkSolvability() bool {
	e.fbrd = flatten_board(e.setFinalState())
	e.brd = flatten_board(e.board)
	izero := get_index(e.brd, 0)
	fizero := get_index(e.fbrd, 0)
	nbinv := e.getInversions(e.brd)
	fnbinv := e.getInversions(e.fbrd)
	if e.size % 2 == 0 {	
		nbinv += izero / e.size
		fnbinv += fizero / e.size
	}
	return (nbinv % 2 == fnbinv % 2)
}

func (e *npuzzle) checkBoard() bool {
	e.checkValues(e.board, e.size)
	solv := e.checkSolvability()
	e.putDebug("solvability?", solv)
	if solv == false {
		putError("this puzzle is unsolvable")
	}
	return solv
}
