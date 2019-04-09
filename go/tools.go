package main

import (
	"os"
	"fmt"
	"strconv"
)

func putUsage() {
	fmt.Printf("Usage: ./npuzzle -f FILENAME\n")
	os.Exit(0)
}

func putNode(n *Node) {
	fmt.Printf("{ f: %5d g: %5d pred: %p    ", n.f, n.g, n.pred)
	for _, v := range n.brd {
		fmt.Printf("%d-", v)
	}
	fmt.Printf(" }\n")
}

func putError(msg string) {
	fmt.Printf("\033[31mError\033[39m: %s\n", msg)
	os.Exit(0)
}

func (e *npuzzle) putDebug(msg string, args ...interface{}) {
	if e.debug == true {
		fmt.Println(fmt.Sprintf("\033[32mDebug\033[39m:%s\n", msg), args, "\n")
	}
}

func (e * npuzzle) putBoard(brd []int) {
	wdth := e.size * 6 + 1
	fmt.Printf("\033[36m ")
	for w:= 0; w < wdth;w++ {
		fmt.Printf("-")
	}
	fmt.Printf("\033[39m\n")

	for i, v := range brd {
		if (i % e.size == 0) {
			fmt.Printf("\033[36m|\033[39m")
		}
		if (v == 0) {
			fmt.Printf("\033[32m%6s\033[39m", "0 ")
		} else {
			fmt.Printf("%5d ", v)
		}
		if ((i + 1) % e.size == 0) {
			fmt.Printf("\033[36m |\n\033[39m")
		}
	}
	fmt.Printf("\033[36m ")
	for w:= 0; w < wdth;w++ {
		fmt.Printf("-")
	}
	fmt.Printf("\033[39m\n\n")
}

func flatten_board(brd [][]int) []int {
	ret := []int{}
	for i:=0; i<len(brd); i++ {
		ret = append(ret, brd[i]...)
	}
	return ret
}

func comp_board(a []int, b []int) bool {
	ln := len(a)
	for i:=0; i<ln; i++ {
		if (a[i] != b[i]) {
			return false
		}
	}
	return true
}

func get_index(brd []int, val int) int {
	ln := len(brd)
	for i:=0;i<ln;i++ {
		if (brd[i] == val) {
			return i
		}
	}
	return -1
}

func node_in_pq(nd *Node, p *PriorityQ) *Node {
	ln := p.Len()

	for i:=0; i<ln; i++ {
		if (comp_board((*p)[i].brd, nd.brd)) {
			return (*p)[i]
		}
	}
	return nil
}

func abs (v int) int {
	if (v < 0) {
		return -v
	}
	return v
}

func max (a int, b int) int {
	if (a < b) {
		return b
	}
	return a
}

func filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
} 

func new_swap(brd []int, a int , b int) []int {
	ret:= make([]int, len(brd))
	copy(ret, brd)
	ret[a], ret[b] = brd[b], brd[a] 
	return ret
}

const (
	UP = 1
	RIGHT = 2
	DOWN = 3
	LEFT = 4
)

func (e * npuzzle) get_node(c chan *Node, nd *Node, index_zero int, index_next int, move int, greedy bool) {
	if (index_next >= e.size * e.size || index_next < 0) {
		c <- &Node{brd: nil}
		return ;
	}
	brd:= new_swap(nd.brd, index_zero, index_next)
	if (nd.pred != nil && comp_board(brd, nd.pred.brd)) {
		c <- &Node{brd: nil}
		return ;
	}
	next := &Node{
		brd:	nil,
		f:		nd.g + 1 + e.heuristic(brd),
		g:		nd.g + 1,
		pred: 	nd,
		to_rm : false,
	}
	switch move {
	case UP:
			next.brd = brd
			next.key = brd_to_key(brd)
			c <- next
	case RIGHT:
		if (index_next % e.size == 0) {
			c <- next
		} else {
			next.brd = brd
			next.key = brd_to_key(brd)
			c <- next
		}
	case LEFT:
		if (index_zero % e.size == 0) {
			c <- next
		} else {
			next.brd = brd
			next.key = brd_to_key(brd)
			c <- next
		}
	case DOWN:
			next.brd = brd
			next.key = brd_to_key(brd)
			c <- next
	default:
		c <- next
	}
}

func (e * npuzzle) get_nodes(nd *Node, c chan *Node, greedy bool) {
	index := get_index(nd.brd, 0)

	go e.get_node(c, nd, index, index + 1, RIGHT, greedy)
	go e.get_node(c, nd, index, index - 1, LEFT, greedy)
	go e.get_node(c, nd, index, index + e.size, DOWN, greedy)
	go e.get_node(c, nd, index, index - e.size, UP, greedy)
}

func (e *npuzzle) build_path(end *Node) []*Node {
	var path []*Node
	for end != nil {
		path = append([]*Node{end}, path...)
		end = end.pred
	}
	return path
}

func brd_to_key(brd []int) string {
	ln := len(brd)
	ret:= ""
	for i:=0;i<ln;i++ {
			ret += "-" + strconv.Itoa(brd[i])
	}
	return ret
}