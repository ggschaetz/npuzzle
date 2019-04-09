NAME	= npuzzle
SRC	= parser.go \
	  tools.go \
	  boardChecker.go \
	  astar.go \
	  heuristics.go \
	  priorityQ.go \
	  main.go

all: $(NAME)

$(NAME): $(SRC)
	go build -o $(NAME)

clean:
	/bin/rm -f $(NAME)

re: clean all

test:
	./npuzzle -f ./tests/nonExistingFile
	./npuzzle -f ./tests/emptyFile
	./npuzzle -f ./tests/onlysize
	./npuzzle -f ./tests/2x2
	./npuzzle -f ./tests/wrongSize
	./npuzzle -f ./tests/notSquare1
	./npuzzle -f ./tests/notSquare2
	./npuzzle -f ./tests/notSquare3
	./npuzzle -f ./tests/notNumber
	./npuzzle -f ./tests/wrongValue1
	./npuzzle -f ./tests/wrongValue2
	./npuzzle -f ./tests/3x3unsolvable
	cat ./tests/withComments1
	./npuzzle -f ./tests/withComments1 -z
	cat ./tests/withComments2
	./npuzzle -f ./tests/withComments2 -z
	cat ./tests/3x3solvable
	./npuzzle -f ./tests/3x3solvable -z
	cat ./tests/greedy_vs_uniform
	./npuzzle -f ./tests/greedy_vs_uniform -g -z
	cat ./tests/4x4solvable
	./npuzzle -f ./tests/4x4solvable -z
	cat ./tests/5x5solvable
	./npuzzle -f ./tests/5x5solvable -z
