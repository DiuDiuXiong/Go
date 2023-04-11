package main

import (
	"container/list"
	"fmt"
	"os"
)

// we use 2D 0/1 for maze representation

type point struct {
	row, col int
}

type maze [][]int

func (p *point) getNextMove(alreadyChecked *map[point]struct{}, row int, col int, m *maze) []*point {
	res := make([]*point, 4)
	if _, ok := (*alreadyChecked)[point{row: p.row - 1, col: p.col}]; p.row-1 >= 0 && (*m)[p.row-1][p.col] != 1 && !ok { // U
		res[0] = &point{row: p.row - 1, col: p.col}
	}

	if _, ok := (*alreadyChecked)[point{row: p.row + 1, col: p.col}]; p.row+1 < row && (*m)[p.row+1][p.col] != 1 && !ok { // D
		res[1] = &point{row: p.row + 1, col: p.col}
	}

	if _, ok := (*alreadyChecked)[point{row: p.row, col: p.col - 1}]; p.col-1 >= 0 && (*m)[p.row][p.col-1] != 1 && !ok { // L
		res[2] = &point{row: p.row, col: p.col - 1}
	}

	if _, ok := (*alreadyChecked)[point{row: p.row, col: p.col + 1}]; p.col+1 < col && (*m)[p.row][p.col+1] != 1 && !ok { // R
		res[3] = &point{row: p.row, col: p.col + 1}
	}
	return res
}

func getRoute(m *maze, row int, col int) { // assume start is (0,0) end is (row - 1, col - 1)
	alreadyChecked := make(map[point]struct{})
	parent := make(map[point]*point)
	remainPath := list.New()

	start := point{row: 0, col: 0}
	remainPath.PushBack(&start)
	for remainPath.Len() > 0 {
		curr := remainPath.Remove(remainPath.Front()).(*point)
		alreadyChecked[*curr] = struct{}{}

		// check if destination, then print route
		if curr.row == row-1 && curr.col == col-1 {
			fmt.Println("Trace found!")
			for curr.row != 0 || curr.col != 0 {
				fmt.Printf("Point pos: row-%d col-%d\n", curr.row, curr.col)
				curr = parent[*curr]
			}
			fmt.Printf("Point pos: row-%d col-%d\n", curr.row, curr.col)
			return
		}

		// get next steps
		for _, n := range curr.getNextMove(&alreadyChecked, row, col, m) {
			if n != nil {
				parent[*n] = curr
				remainPath.PushBack(n)
			}
		}
	}
	fmt.Println("Route Not found.")
}
func readMaze(mazeName string) maze {
	// 01 read file
	file, err := os.Open(mazeName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	// 02 read row & column
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	// 03 maze initialize & read maze
	resMaze := make(maze, row)
	for r := range resMaze {
		resMaze[r] = make([]int, col)
		for c := range resMaze[r] {
			fmt.Fscanf(file, "%d", &resMaze[r][c])
		}
	}

	return resMaze
}

func main() {
	// fmt.Println(readMaze("12_maze_algorithm/maze.in"))  // verify that maze if correctly read in

	// m := readMaze("12_maze_algorithm/maze.in") // verify that getNextMove is correct
	// fmt.Println((&point{row: 0, col: 0}).getNextMove(&map[point]struct{}{}, 6, 5, &m))
	// fmt.Println((&point{row: 0, col: 1}).getNextMove(&map[point]struct{}{}, 6, 5, &m))
	// fmt.Println((&point{row: 1, col: 1}).getNextMove(&map[point]struct{}{}, 6, 5, &m))

	// m := readMaze("12_maze_algorithm/maze.in")
	// m := readMaze("12_maze_algorithm/maze1.in")
	// m := readMaze("12_maze_algorithm/maze2.in")
	// m := readMaze("12_maze_algorithm/maze3.in")
	m := readMaze("12_maze_algorithm/maze4.in")

	row, col := len(m), len(m[0])
	getRoute(&m, row, col)

}
