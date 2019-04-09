package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type npuzzle struct {
	debug	bool
	size	int
	board	[][]int
	brd		[]int
	fbrd	[]int
	heuristic func(brd []int) int
	greedy	bool
	hide	bool
	cpltime	int
	cplsize	int
}

func main() {
	var npzzl npuzzle
	var greedy_npzzl npuzzle
	gchan := make(chan []*Node)
	
	npzzl.parseArguments()
	npzzl.checkBoard()

	greedy_npzzl = npzzl
	greedy_npzzl.heuristic = func (brd []int) int { return 0 }
	go greedy_npzzl.greedy_search(gchan) //threading the greedy search
	path := npzzl.solver(true)
	npzzl.print_info(path, &greedy_npzzl, gchan)
	npzzl.get_input_n_print_brd(path)
}

func (e * npuzzle) greedy_search(gchan chan []*Node) {
	if (e.greedy) {
		gchan <- e.solver(true)
	}
}

func (e *npuzzle) print_info (path []*Node, greedy_npzzl* npuzzle, gchan chan []*Node) {
	if (e.greedy) {
		greedy_path := <- gchan
		greedy_npzzl.print_info_greedy(greedy_path)
	} 
	fmt.Printf("\033[32m ----------------------------\033[39m\n")
	fmt.Printf("\033[32m|\033[39m       GREEDY SEARCH        \033[32m|\033[39m\n")
	fmt.Printf("\033[32m ----------------------------\033[39m\n")
	fmt.Printf("\033[32m|\033[39mcomplexity in time: %7d \033[32m|\033[39m\n", e.cpltime)
	fmt.Printf("\033[32m|\033[39mcomplexity in size: %7d \033[32m|\033[39m\n", e.cplsize)
	fmt.Printf("\033[32m|\033[39mnumber of moves   : %7d \033[32m|\033[39m\n", len(path) - 1)
	fmt.Printf("\033[32m ----------------------------\033[39m\n")
}

func (e *npuzzle) print_info_greedy (path []*Node) {

	fmt.Printf("\033[33m ----------------------------\033[39m\n")
	fmt.Printf("\033[33m|\033[39m    UNIFORM COST SEARCH     \033[33m|\033[39m\n")
	fmt.Printf("\033[33m ----------------------------\033[39m\n")
	fmt.Printf("\033[33m|\033[39mcomplexity in time: %7d \033[33m|\033[39m\n", e.cpltime)
	fmt.Printf("\033[33m|\033[39mcomplexity in size: %7d \033[33m|\033[39m\n", e.cplsize)
	fmt.Printf("\033[33m|\033[39mnumber of moves   : %7d \033[33m|\033[39m\n", len(path) - 1)
	fmt.Printf("\033[33m ----------------------------\033[39m\n")
	fmt.Printf("              -              \n")

}


func (e *npuzzle) get_input_n_print_brd (path []*Node) {

	reader := bufio.NewReader(os.Stdin)
	done := false
	step := 0
	
	if (e.hide) {
		return
	}
	fmt.Printf("\033[32m> Do you want to print the solution [Y/n/s] ?\033[32m|\033[39m\n")
	for ;done == false; {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, " ", "", -1)
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "y", "Y":
			for _, n := range path{
				e.putBoard(n.brd)
			}
			done = true
		case "n":
			done = true
		case "s", "":
			if (step < len(path)) {
				e.putBoard(path[step].brd)
				step++;
			}
			if (step >= len(path)) {
				done = true
			}
		default:
				done = false
		}

	}



	
	
}