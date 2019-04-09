package main

import (
	"fmt"
	"container/heap"
	"sort"
	"os"
)

func (e *npuzzle) solver(greedy bool) []*Node {

	openQ := PriorityQ{}
	openMap := map[string]*Node{}
	closeMap := map[string]*Node{}
	nextChan := make(chan *Node)
	root := &Node {
		brd:	e.brd,
		f:		e.heuristic(e.brd),
		g:		0,
		pred:	nil,
		key:	brd_to_key(e.brd),
		to_rm:	false,
	}
	e.cplsize = 0;
	e.cpltime = 0;
	if (greedy) {
		root.f = root.g
	}
	heap.Push(&openQ, root)
	openMap[root.key] = root
	ln := openQ.Len()

	for ln > 0 {
		e.cpltime++
		if (ln > e.cplsize) {
			e.cplsize = ln
		}
		if (e.debug) {
			fmt.Println(openQ.Len(), len(closeMap), ">")
			for i:=0; i < openQ.Len();i++ {
				putNode(openQ[i])
			}
		}
		current := heap.Pop(&openQ).(*Node)
		closeMap[current.key] = current
		if (e.debug) {
			fmt.Println("Current")
			putNode(current)
			fmt.Println("--")
		}
		if(comp_board(current.brd, e.fbrd)) {
			return e.build_path(current)
		}
		e.get_nodes(current, nextChan, greedy)
		for i:=0;i<4;i++ {
			next := <- nextChan
			if (next.brd == nil) {
				continue
			}
			_, okclose := closeMap[next.key]
			if (okclose == true) {
				continue
			}
			_, okopen := openMap[next.key]
			if (okopen == true) {
				continue
			} else {
				heap.Push(&openQ, next)
				openMap[next.key] = next
			}
			ln = openQ.Len()
			sort.Sort(openQ)
			if (e.debug) {
				for i:=0; i < openQ.Len();i++ {
					putNode(openQ[i])
				}
				if ln >= 10 {
					os.Exit(1);
				}
			}
		}
	}
	return []*Node{}
}