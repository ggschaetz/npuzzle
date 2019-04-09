package main

import (
)

type Node struct {
	brd []int
	f int  // g + h
	g int // cost from root
	pred *Node
	key string
	to_rm bool
}

type PriorityQ []*Node

func (p PriorityQ) Len() int {
	return len(p)
}

func (p PriorityQ) Less(ia int, ib int) bool {
	return p[ia].f < p[ib].f
}

func (p *PriorityQ) Push(elm interface{}) {
	node := elm.(*Node)
	*p = append(*p, node)
}

func (p *PriorityQ) Pop() interface{} {
	old := *p
	n := len(old)
	nodePop := old[n - 1]
	*p = old[:n - 1]
	return nodePop
}

func (p PriorityQ) Swap(ia int, ib int) {
	p[ia], p[ib] = p[ib], p[ia]
}
