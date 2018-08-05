// Copyright (C) 2018 Yoeri Apts - All Rights Reserved
// You may use, distribute and modify this code under the terms
// of the MIT license, see https://opensource.org/licenses/MIT

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const X int = -1

type Row []int
type Puzzle []Row

type Cell struct {
	fixed bool
	value int
}

type Board struct {
	cells [][]Cell
	dim_y int
	dim_x int
}

// Set next value to test on cell[x,y]
func (board *Board) setNextValue(x, y int) bool {
	if board.cells[y][x].fixed {
		return true
	}
	switch board.cells[y][x].value {
	case X:
		board.cells[y][x].value = 0
		return true
	case 0:
		board.cells[y][x].value = 1
		return true
	case 1:
		board.cells[y][x].value = X
		return false
	default:
		return false
	}
}

// Check that current board is still valid (='possible'), board is checked left to right, top to bottom
func (board *Board) isBoardStillValid(x, y int) bool {

	// look left 2 places, should never see 3 identical digits next to each other
	if x >= 2 && (board.cells[y][x-2].value == board.cells[y][x-1].value && board.cells[y][x-1].value == board.cells[y][x].value) {
		return false
	}
	// look 1 place left and 1 place right, should never see 3 identical digits next to each other
	if x >= 1 && x < board.dim_x-1 &&
		(board.cells[y][x-1].value == board.cells[y][x].value) &&
		(board.cells[y][x].value == board.cells[y][x+1].value) {
		return false
	}
	// look 2 places to the right, should never see 3 identical digits next to each other
	if x >= 0 && x < board.dim_x-2 &&
		(board.cells[y][x].value == board.cells[y][x+1].value) &&
		(board.cells[y][x+1].value == board.cells[y][x+2].value) {
		return false
	}

	// look up 2 places, should never see 3 identical digits above each other
	if y >= 2 && (board.cells[y-2][x].value == board.cells[y-1][x].value && board.cells[y-1][x].value == board.cells[y][x].value) {
		return false
	}
	// look up 1 place, down 1 place, should never see 3 identical digits above each other
	if y >= 1 && y < board.dim_y-1 &&
		(board.cells[y-1][x].value == board.cells[y][x].value) &&
		(board.cells[y][x].value == board.cells[y+1][x].value) {
		return false
	}
	// look down 2 places, should never see 3 identical digits above each other
	if y >= 0 && y < board.dim_y-2 &&
		(board.cells[y][x].value == board.cells[y+1][x].value) &&
		(board.cells[y+1][x].value == board.cells[y+2][x].value) {
		return false
	}

	// Count number of 0s and 1s in this row so far
	nmbr_of_1 := 0
	nmbr_of_0 := 0
	for i := 0; i <= x; i++ {
		if board.cells[y][i].value == 0 {
			nmbr_of_0 += 1
		} else if board.cells[y][i].value == 1 {
			nmbr_of_1 += 1
		} else {
			// this should NEVER happen
			return false
		}
	}
	// can never be more 0s or 1s than halve of row
	if (nmbr_of_0 > board.dim_x/2) || (nmbr_of_1 > board.dim_x/2) {
		return false
	}
	// no 2 rows can be identical (check only when row complete)
	if (x == board.dim_x-1) && (y >= 1) {
		identical := true
		for row1 := 0; row1 < y; row1++ {
			for row2 := row1 + 1; row2 <= y; row2++ {
				for x_ := 0; x_ < board.dim_x; x_++ {
					if board.cells[row1][x_].value == board.cells[row2][x_].value {
						continue
					} else {
						identical = false
						break
					}
				}
				if identical {
					break
				}
			}
			if identical {
				break
			}
		}
		if identical {
			return false
		}
	}

	// Count number of 0s and 1s in this column so far
	nmbr_of_1 = 0
	nmbr_of_0 = 0
	for i := 0; i <= y; i++ {
		if board.cells[i][x].value == 0 {
			nmbr_of_0 += 1
		} else if board.cells[i][x].value == 1 {
			nmbr_of_1 += 1
		} else {
			return false
		}
	}
	// can never be more 0s or 1s than halve of columns
	if (nmbr_of_0 > board.dim_y/2) || (nmbr_of_1 > board.dim_y/2) {
		return false
	}

	// no 2 columns can be identical (check only when column is complete)
	if (y == board.dim_y-1) && (x > 0) {
		identical := true
		for col1 := 0; col1 < x; col1++ {
			for col2 := col1 + 1; col2 <= x; col2++ {
				for y_ := 0; y_ < board.dim_y; y_++ {
					if board.cells[y_][col1].value == board.cells[y_][col2].value {
						continue
					} else {
						identical = false
						break
					}
				}
				if identical {
					break
				}
			}
			if identical {
				break
			}
		}
		if identical {
			return false
		}
	}

	return true
}

// Return string representing a board, the fixed cells are suffixed with *
func (board Board) Show() string {
	s := ""
	//for y := 0; y < len(p); y++ {
	for _, row := range board.cells {
		//for x := 0; x < len(p[y]); x++ {
		for _, cell := range row {
			switch cell.value {
			case X:
				s += "X"
			case 0:
				s += "0"
			case 1:
				s += "1"
			default: // invalid!
				s += "?"
			}
			if cell.fixed {
				s += "* "
			} else {
				s += "  "
			}
		}
		s += "\n"
	}
	return s
}

// return coordinates of previous, not fixed cell (for backtracking)
func (board *Board) getPreviousXY(x, y int) (x_, y_ int, err error) {
	x_ = x - 1
	y_ = y
	if x_ < 0 {
		x_ = board.dim_x - 1
		y_ -= 1
	}
	if y_ < 0 {
		return -1, -1, errors.New("Already at first cell")
	} else {
		if board.cells[y_][x_].fixed {
			return board.getPreviousXY(x_, y_)
		}
		return x_, y_, nil
	}
}

// return coordinates of next cell to try
func (board *Board) getNextXY(x, y int) (x_, y_ int, err error) {
	x_ = x + 1
	y_ = y
	if x_ >= board.dim_x {
		x_ = 0
		y_ += 1
	}
	if y_ >= board.dim_y {
		return -1, -1, errors.New("Already at last cell")
	} else {
		return x_, y_, nil
	}
}

// Solve this thing already!
func (board *Board) Solve() (solved bool, err error) {
	counter := 0
	x, y := 0, 0
	solved = false
	for solved == false {
		if counter%10000000 == 0 {
			fmt.Printf("(%v), At [%v,%v]\n", counter, x, y)
			fmt.Print(board.Show())
		}
		counter++
		valid := board.setNextValue(x, y)
		// valid is true when cell fixed or cell not fixed but set to 0 or 1
		if valid && board.isBoardStillValid(x, y) {
			// fmt.Println("isBoardStillValid is true")
			x, y, err = board.getNextXY(x, y)
			if err != nil {
				// Fell of the end of the board... solved !!!
				fmt.Println("Solved!")
				solved = true
				break
			}
		} else {
			//fmt.Println("isBoardStillValid is false")
			if valid && !board.cells[y][x].fixed {
				continue
			}
			x, y, err = board.getPreviousXY(x, y)
			if err != nil {
				// Fell of the beginning of the board... NO solution
				solved = false
				fmt.Println("No solution found")
				break
			}
		}
	}

	fmt.Printf("counter=%v, solved=%v \n", counter, solved)
	return solved, err
}

// Stringer for a Puzzle
func (puzzle Puzzle) String() string {
	s := ""
	for _, row := range puzzle {
		for _, x := range row {
			switch x {
			case X:
				s += "X  "
			case 0:
				s += "0  "
			case 1:
				s += "1  "
			default: // invalid!
				s += "?  "
			}
		}
		s += "\n"
	}
	return s
}

// load a Puzzle into a Board
func puzzle_to_board(puzzle Puzzle) (Board, error) {
	// The board to return
	var board Board

	// allocate composed 2d array
	rows := len(puzzle)
	cols := len(puzzle[0])

	if rows%2 != 0 || cols%2 != 0 {
		return board, errors.New(fmt.Sprintf("rows[%v] or cols[%v] not a multiple of 2", rows, cols))
	}
	a := make([][]Cell, rows)
	e := make([]Cell, rows*cols)
	for i := range a {
		a[i] = e[i*cols : (i+1)*cols]
	}

	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[y]); x++ {
			// put puzzle in board
			switch puzzle[y][x] {
			case X:
				a[y][x] = Cell{false, X}
			case 0:
				a[y][x] = Cell{true, 0}
			case 1:
				a[y][x] = Cell{true, 1}
			default: // invalid!
				return board, errors.New(fmt.Sprintf("Cell [%v,%v] is not 0, 1 or X", y, x))
			}
		}
	}

	board.dim_y = len(puzzle)
	board.dim_x = len(puzzle[0])
	board.cells = a
	return board, nil
}

// Load some predefined puzzles (used for interactive monkey testing)
func load_predef_puzzles() []Puzzle {
	var puzzles []Puzzle

	// 6 x 6, all undefined
	puzzles = append(puzzles, Puzzle{
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
	})

	// 14 x 12, all undefined
	// http: //www.binairepuzzel.net/puzzels.php?size=10&level=3&nr=53
	puzzles = append(puzzles, Puzzle{
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
	})

	// 8 x 6 puzzle
	puzzles = append(puzzles, Puzzle{
		{1, X, 1, 0, X, X, X, X},
		{1, X, X, 1, X, 0, X, 1},
		{0, 0, X, X, X, X, X, X},
		{1, 0, X, X, X, X, X, X},
		{0, 1, X, X, 1, X, X, 1},
		{0, 1, X, X, X, X, X, X},
	})

	// http://www.binairepuzzel.net/dagpuzzel.php?id=2970
	// 12x12http://www.binairepuzzel.net/puzzels.php?size=10&level=3&nr=53
	puzzles = append(puzzles, Puzzle{
		{X, X, X, X, 1, X, X, 0, 0, X, 0, X},
		{X, X, X, 0, X, 1, X, X, X, X, X, X},
		{X, 1, X, X, X, X, 1, X, X, X, X, X},
		{X, X, 0, X, X, X, X, 1, X, X, X, X},
		{X, X, 0, 1, X, X, X, X, X, 1, 1, X},
		{X, X, X, X, X, X, X, X, X, X, X, X},
		{X, X, 1, X, X, X, X, 0, X, X, X, X},
		{1, 1, X, X, X, X, X, X, 1, X, X, 1},
		{X, 1, 1, X, X, 0, 1, X, 1, X, X, X},
		{X, X, 1, X, X, X, X, X, X, X, X, X},
		{1, X, X, X, X, X, X, 1, X, X, 1, X},
		{X, 0, X, 0, X, X, 0, X, X, X, X, X},
	})

	//14x14 Very Difficult 80http://www.binairepuzzel.net/puzzels.php?size=10&level=3&nr=53
	//http://www.binairepuzzel.net/puzzels.php?size=14&level=4&nr=80
	puzzles = append(puzzles, Puzzle{
		{X, X, X, 1, X, X, X, X, 1, X, X, X, X, X},
		{1, X, X, X, 0, X, X, 0, X, X, 1, X, X, X},
		{X, X, X, 1, X, 1, X, X, 1, X, X, X, X, 0},
		{X, X, 1, X, X, X, X, X, X, X, X, X, 1, X},
		{0, X, 1, X, X, X, X, X, X, X, X, 0, X, X},
		{0, X, X, 0, X, X, X, 1, 1, X, X, X, X, X},
		{X, X, 0, X, X, X, X, X, X, X, 1, X, 1, X},
		{X, X, X, X, 0, X, 1, X, X, 0, X, 0, 1, X},
		{X, 0, X, X, X, X, X, X, X, X, X, X, X, X},
		{0, 0, X, X, X, X, X, X, 0, X, X, X, X, X},
		{X, X, X, X, X, X, X, X, X, X, X, X, X, X},
		{X, 0, X, X, X, 1, X, X, 1, X, 0, X, X, X},
		{1, X, X, X, X, X, X, X, X, X, 1, 1, X, X},
		{X, X, X, X, 1, 1, X, X, X, X, X, X, X, X},
	})
	// Solved!
	// Duration:  27.837238418s
	// 	1  1  0  1* 1  0  0  1  1* 0  0  1  0  0
	// 	1* 0  1  0  0* 1  1  0* 0  1  1* 0  0  1
	// 	0  1  0  1* 0  1* 0  1  1* 0  0  1  1  0*
	// 	1  0  1* 0  1  0  1  0  0  1  0  1  1* 0
	// 	0* 0  1* 1  0  0  1  1  0  1  1  0* 0  1
	// 	0* 1  0  0* 1  1  0  1* 1* 0  0  1  0  1
	// 	1  0  0* 1  1  0  1  0  0  1  1* 0  1* 0
	// 	0  1  1  0  0* 1  1* 0  1  0* 1  0* 1* 0
	// 	1  0* 0  1  0  1  0  1  1  0  0  1  0  1
	// 	0* 0* 1  0  1  0  0  1  0* 1  1  0  1  1
	// 	1  1  0  1  0  0  1  0  0  1  1  0  1  0
	// 	0  0* 1  1  0  1* 1  0  1* 0  0* 1  0  1
	// 	1* 1  0  0  1  0  0  1  0  0  1* 1* 0  1
	// 	0  1  1  0  1* 1* 0  0  1  1  0  0  1  0

	// 6 x 6, invalid puzzle because it has 3 ones on the last line
	// The solver now only finds that the puzzle has no solution
	// after it exhausts all possible combinations; should add feature
	// that checks for a 'valid' puzzle before starting
	puzzles = append(puzzles, Puzzle{
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, X, X, X},
		{X, X, X, 1, 1, 1},
	})

	return puzzles
}

// read puzzle from given file
func puzzle_from_file(fname string) Puzzle {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	row := make(Row, 0)
	puzzle := make(Puzzle, 0)

	buffer := make([]byte, 1)
	in_comment := false

	for err == nil {
		//var n1 int
		//n1, err = f.Read(buffer)
		_, err = f.Read(buffer)
		if err != nil {
			// eof reached
			break
		}

		c := buffer[0]
		//fmt.Printf("n1:%v, c:%c\n", n1, c)

		if in_comment {
			// processing a comment, eat all chars till eol
			if c == '\n' {
				in_comment = false
			} else {
				continue
			}
		}

		switch c {
		case '0':
			//fmt.Println("0 found")
			row = append(row, 0)
		case '1':
			row = append(row, 1)
		case 'X':
			row = append(row, -1)
		case '#':
			//start processing comment
			in_comment = true
		case '\n':
			if len(row) > 1 {
				puzzle = append(puzzle, row)
				row = make(Row, 0)
			}
		default:
			//fmt.Println("default: char %v", int(c))
		}
	}
	f.Close()

	return puzzle
}

func usage(cmd string) {
	fmt.Printf("Usage: %s puzzle-filename, or\n", os.Args[0])
	fmt.Printf("       %s -p puzzle-number\n", os.Args[0])
}

func main() {
	puzzle := Puzzle{}

	// fmt.Println("board:")
	// fmt.Println(board.Show())

	// Process command line arguments
	if len(os.Args) != 2 && len(os.Args) != 3 {
		usage(os.Args[0])
		os.Exit(-1)
	}

	if len(os.Args) == 2 {
		fmt.Printf("Reading puzzle from file '%s'\n", os.Args[1])
		puzzle = puzzle_from_file(os.Args[1])
	}

	if len(os.Args) == 3 {
		if os.Args[1] != "-p" {
			usage(os.Args[0])
			os.Exit(-1)
		}

		fmt.Printf("Reading predefined puzzle nmbr '%s'\n", os.Args[2])
		puzzles := load_predef_puzzles()
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("%s is not an integer", os.Args[2])
			panic(err)
		}

		if index < 0 || index >= len(puzzles) {
			fmt.Printf("The predefined puzzles are numbered %d to %d\n", 0, len(puzzles)-1)
			os.Exit(-1)
		}
		puzzle = puzzles[index]
	}

	// Solve chosen puzzle
	fmt.Println("Puzzle to solve:")
	fmt.Println(puzzle)

	var board Board
	board, err := puzzle_to_board(puzzle)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Printf("Dimensions: %d by %d \n", board.dim_x, board.dim_y)

	time_start := time.Now()
	fmt.Println("Started at: ", time_start)

	solved, err := board.Solve()

	time_stop := time.Now()
	fmt.Println("Stopped at: ", time_stop)

	duration := time_stop.Sub(time_start)
	fmt.Println("Duration: ", duration)

	if solved {
		fmt.Println("Solution:")
		fmt.Print(board.Show())
	}
}
