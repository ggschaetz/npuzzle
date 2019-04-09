package main

import (
		"os"
		"flag"
		"fmt"
		"strings"
		"strconv"
		"io/ioutil"
       )

func (e *npuzzle) parseArguments() bool {
	filename := flag.String("f", "", "file that contains the N puzzle (required)")
	flag.BoolVar(&e.debug, "d", false, "prints the debug")
	heuristic := flag.String("h", "m", `choose from 5 heuristics:
	- m   => manhattan (default)
	- dd  => diagonal distance
	- mlc => manhattan + linear conflict
	- eu  => euclidian distance
	- mt  => misplaced tiles`)
	flag.BoolVar(&e.greedy, "g", false, "greedy search")
	flag.BoolVar(&e.hide, "z", false, "hide solution")
	flag.Parse()

	if *filename == "" {
		putUsage()
	}
	if (*heuristic != "m" && *heuristic != "dd" && *heuristic != "mlc" && *heuristic != "mt" && *heuristic != "eu") {
		putError("invalid heuristic [ m | dd | mlc | eu | mt ]")
	}
	switch *heuristic {
		case "m":
			e.heuristic = e.h_manhattan
		case "dd":
			e.heuristic = e.h_diagonal_distance
		case "mlc":
			e.heuristic = e.h_mlc
		case "mt":
			e.heuristic = e.h_misplaced_tiles
		case "eu":
			e.heuristic = e.h_euclidian_distance
		default:
			e.heuristic = e.h_manhattan
	}
	if (e.debug) {
		fmt.Println("heuristic", *heuristic)
	}
	_, err := os.Open(*filename)
	if err != nil {
		putError("cannot open the file")
	}
	file, err := ioutil.ReadFile(*filename)
	if err != nil {
		putError("Not a valid file or flag -f is missing")
	}
	return e.checkFileContent(strings.Trim(string(file), "\n"))
}

func getUserHeight(str string) (bool, int) {
	height := -1
	line := strings.Split(strings.TrimSpace(str), " ")

      	if len(line) != 1 {
	      return false, height
	}
	height, err := strconv.Atoi(line[0])
	if err != nil {
		return false, height
	}
	return true, height
}

func filterComments(vs []string) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if len(v) > 0 && v[0] != '#' {
			str := strings.Split(v, "#")
			if len(str) > 0 && len(str[0]) > 0 {
				vsf = append(vsf, strings.TrimSpace(str[0]))
			}
		}
	}
	return vsf
}

func boardAtoi(lines []string, height int) [][]int {
	brd := make([][]int, height)
	for y, line := range lines {
		brd[y] = make([]int, height)
		ns := filter(strings.Split(line, " "), func (s string) bool { return s != ""})
		if len(ns) != height {
			putError("this is not a square")
		}
		for x, n := range ns {
			val, err := strconv.Atoi(n)
			if err != nil {
				putError(fmt.Sprintf("not a number at [%d:%d]", y, x))
			}
			brd[y][x] = val
		}
	}
	return brd
}

func (e *npuzzle) checkFileContent(str string) bool {
	lines := filterComments(strings.Split(str, "\n"))
	e.putDebug("after filtering comment", lines)
	height := len(lines)
	if height == 0 {
		putError("the file looks empty")
	}
	if height == 1 {
		putError("there's only one line ?!")
	}
	valid, userHeight := getUserHeight(lines[0])
	if valid == false {
		putError("the first line must be the size")	
	}
	if height - 1 != userHeight {
		putError("the height on the first line doesnt match the actual height")
	}
	if height < 4 {
		putError("the size must be at least 3")
	}
	e.size = userHeight
	e.board = boardAtoi(lines[1:], userHeight)
	return true
}
