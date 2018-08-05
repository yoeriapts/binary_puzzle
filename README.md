# Binary puzzle / Binario / Binero solver

You may like solving a single _binary puzzle_ every day, but this golang program
will solve every binary puzzle in a single day (kinda ;-) )

Binary puzzles are also called 'binario' or 'binero' puzzles.

Rules:
- Every puzzle has an even number of rows and columns eg: an 8 by 6 puzzle, most puzzles are square (n by n)
- Each cell can contain only a 0 or a 1
- There can be no more than two similar numbers below or next to each other
- Each row and each column is unique
- Each row and each columns contains as many zeros as ones

Some links concerning these binary puzzles:
- http://www.binarypuzzle.com/
- https://lovattspuzzles.com/2014/02/17/binary-puzzle-video-tutorial/

## Program execution
clone the repo, then
```sh
$ go build bp.go
$ ./bp
Usage: ./bp puzzle-filename, or
       ./bp -p puzzle-number
```
Solve one of the built-in puzzles
```sh
$ ./bp -p 4
Reading predefined puzzle nmbr '4'
Puzzle to solve:
X  X  X  1  X  X  X  X  1  X  X  X  X  X  
1  X  X  X  0  X  X  0  X  X  1  X  X  X  
X  X  X  1  X  1  X  X  1  X  X  X  X  0  
X  X  1  X  X  X  X  X  X  X  X  X  1  X  
0  X  1  X  X  X  X  X  X  X  X  0  X  X  
0  X  X  0  X  X  X  1  1  X  X  X  X  X  
X  X  0  X  X  X  X  X  X  X  1  X  1  X  
X  X  X  X  0  X  1  X  X  0  X  0  1  X  
X  0  X  X  X  X  X  X  X  X  X  X  X  X  
0  0  X  X  X  X  X  X  0  X  X  X  X  X  
X  X  X  X  X  X  X  X  X  X  X  X  X  X  
X  0  X  X  X  1  X  X  1  X  0  X  X  X  
1  X  X  X  X  X  X  X  X  X  1  1  X  X  
X  X  X  X  1  1  X  X  X  X  X  X  X  X  

Dimensions: 14 by 14
Started at:  2018-08-03 22:02:51.588859736 +0200 CEST
Solved!
counter=438258047, solved=true
Stopped at:  2018-08-03 22:05:48.155709596 +0200 CEST
Duration:  2m56.56684986s
Solution:
1  1  0  1* 1  0  0  1  1* 0  0  1  0  0  
1* 0  1  0  0* 1  1  0* 0  1  1* 0  0  1  
0  1  0  1* 0  1* 0  1  1* 0  0  1  1  0*
1  0  1* 0  1  0  1  0  0  1  0  1  1* 0  
0* 0  1* 1  0  0  1  1  0  1  1  0* 0  1  
0* 1  0  0* 1  1  0  1* 1* 0  0  1  0  1  
1  0  0* 1  1  0  1  0  0  1  1* 0  1* 0  
0  1  1  0  0* 1  1* 0  1  0* 1  0* 1* 0  
1  0* 0  1  0  1  0  1  1  0  0  1  0  1  
0* 0* 1  0  1  0  0  1  0* 1  1  0  1  1  
1  1  0  1  0  0  1  0  0  1  1  0  1  0  
0  0* 1  1  0  1* 1  0  1* 0  0* 1  0  1  
1* 1  0  0  1  0  0  1  0  0  1* 1* 0  1  
0  1  1  0  1* 1* 0  0  1  1  0  0  1  0  

```
Solve a puzzle defined in a file:
```sh
$ cat puzzles/puzzle0.txt
# Example puzzle
#  8 x 6 cells
#--------------
X X X X X X X 1
0 X 1 X X X 0 X
X X 0 X X 0 X X
0 X 0 X X X X X	# OXO ?
X 0 X X 1 X 0 X
X X X X X X X 1
#--------------

$ ./bp puzzles/puzzle0.txt
Reading puzzle from file 'puzzles/puzzle0.txt'
Puzzle to solve:
X  X  X  X  X  X  X  1  
0  X  1  X  X  X  0  X  
X  X  0  X  X  0  X  X  
0  X  0  X  X  X  X  X  
X  0  X  X  1  X  0  X  
X  X  X  X  X  X  X  1  

Dimensions: 8 by 6
Started at:  2018-08-03 22:00:55.435818424 +0200 CEST
Solved!
counter=3123, solved=true
Stopped at:  2018-08-03 22:00:55.43653303 +0200 CEST
Duration:  714.606µs
Solution:
0  0  1  1  0  0  1  1*
0* 1  1* 0  1  1  0* 0  
1  0  0* 1  0  0* 1  1  
0* 1  0* 1  0  1  1  0  
1  0* 1  0  1* 1  0* 0  
1  1  0  0  1  0  0  1*
```
to solve your own puzzle, simply type it into a text file and feed that to the solver
```sh
$ ./bp puzzles/your_puzzle.txt
```

## Logic
The solver uses deep learning AI techniques. NO, scratch that! The solver uses bruto force to run over all possible (valid) board combinations starting from the top left, over all the rows (horizontal) and down (vertical), It backtracks when the board becomes invalid.

## Shortcomings
- The solver applies no heuristics
  - example: when 2 adjacent cells contain zeroes, the cells to the left and right of those can only be ones
- The solver stops after finding 1 solution
- If the puzzle input is illegal, for example a puzzle with 3 zeroes in one row, the
solver will only notice this after exhausting all possibilities, which can take a long time

## Disclaimer
I wrote this little program to learn some Go. There are better/faster binary puzzle
solvers out there!

## License

[MIT](LICENSE)
