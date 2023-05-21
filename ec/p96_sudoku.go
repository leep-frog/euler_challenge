package eulerchallenge

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P96() *problem {
	return fileInputNode(96, func(lines []string, o command.Output) {
		n := 9
		var boards []*sudokuBoard
		for len(lines) > 0 {
			lines = lines[1:]

			var board [][]*sudokuCell
			for i := 0; i < n; i++ {
				var row []*sudokuCell
				for j := 0; j < n; j++ {
					row = append(row, newCell(i, j, parse.Atoi(lines[i][j:j+1]), n))
				}
				board = append(board, row)
			}
			boards = append(boards, &sudokuBoard{n, maths.Sqrt(n), board, n * n, false, 0})
			lines = lines[n:]
		}

		sum := 0

		for idx, board := range boards {
			board.Solve()
			path, _ := bfs.Search[string]([]*sudokuBoard{board})
			if len(path) == 0 {
				o.Stderrln("unsolved board ", idx)
				o.Stderrln(board)
				return
			}
			finalBoard := path[len(path)-1].board
			sum += *(finalBoard[0][0].value) * 100
			sum += *(finalBoard[0][1].value) * 10
			sum += *(finalBoard[0][2].value)
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args:     []string{"p96.txt"},
			want:     "24702",
			estimate: 0.5,
		},
		{
			args: []string{"p96_example.txt"},
			want: "483",
		},
	})
}

type sudokuBoard struct {
	n         int
	sn        int
	board     [][]*sudokuCell
	remaining int
	broken    bool

	guesses int
}

func (sb *sudokuBoard) String() string {
	var all []string
	for i := 0; i < sb.n; i++ {
		var row []string
		for j := 0; j < sb.n; j++ {
			row = append(row, fmt.Sprintf("%v", sb.board[i][j]))
			if (j+1)%sb.sn == 0 && j+1 != sb.n {
				row = append(row, "|")
			}
		}
		all = append(all, strings.Join(row, " "))
		if (i+1)%sb.sn == 0 && i+1 != sb.n {
			all = append(all, strings.Repeat("-", 2*(sb.n+sb.sn-1)-1))
		}
	}
	return strings.Join(all, "\n")
}

func (sb *sudokuBoard) sudokuSquare(i, j int) int {
	return (i/sb.sn)*sb.sn + (j / sb.sn)
}

func (sb *sudokuBoard) sudokuSquareIndices(i, j int) [][]int {
	var indices [][]int
	for row := (i / sb.sn) * sb.sn; row < (i/sb.sn)*sb.sn+sb.sn; row++ {
		for col := (j / sb.sn) * sb.sn; col < (j/sb.sn)*sb.sn+sb.sn; col++ {
			indices = append(indices, []int{row, col})
		}
	}
	return indices
}

func sudokuSquareIndicesFromNumber(square, n, sn int) [][]int {
	var indices [][]int
	rowStart := (square / sn) * sn
	colStart := (square % sn) * sn
	for row := rowStart; row < rowStart+sn; row++ {
		for col := colStart; col < colStart+sn; col++ {
			indices = append(indices, []int{row, col})
		}
	}
	return indices
}

type sudokuCell struct {
	row     int
	col     int
	options map[int]bool
	value   *int
}

func (sc *sudokuCell) copy() *sudokuCell {
	var k *int
	if sc.value != nil {
		t := *sc.value
		k = &t
	}
	mc := map[int]bool{}
	for k, v := range sc.options {
		mc[k] = v
	}
	return &sudokuCell{sc.row, sc.col, mc, k}
}

func (sc *sudokuCell) String() string {
	if sc.value == nil {
		return "0"
	}
	return fmt.Sprintf("%d", *sc.value)
}

func newCell(row, col, val, n int) *sudokuCell {
	if val != 0 {
		// Keep value as nil so we can update the board after all cells are set
		return &sudokuCell{row, col, map[int]bool{val: true}, nil}
	}
	opts := map[int]bool{}
	for i := 1; i <= n; i++ {
		opts[i] = true
	}
	return &sudokuCell{row, col, opts, nil}
}

func (sc *sudokuCell) done() bool {
	return sc.value != nil
}

// Returns whether or not a cell has been updated
func (sc *sudokuCell) updateBoard(sb *sudokuBoard) bool {
	if sc.done() {
		return false
	}

	if len(sc.options) == 0 {
		sb.broken = true
		return false
		//panic("cell has no options left")
	}

	if len(sc.options) > 1 {
		return false
	}

	// Only one option left and hasn't been set yet
	for v := range sc.options {
		k := v
		sc.value = &k
	}

	sb.remaining--

	// Update rows and columns
	for col := 0; col < sb.n; col++ {
		if cell := sb.board[sc.row][col]; !cell.done() {
			delete(cell.options, *sc.value)
		}
	}

	for row := 0; row < sb.n; row++ {
		if cell := sb.board[row][sc.col]; !cell.done() {
			delete(cell.options, *sc.value)
		}
	}

	for _, point := range sb.sudokuSquareIndices(sc.row, sc.col) {
		if cell := sb.board[point[0]][point[1]]; !cell.done() {
			delete(cell.options, *sc.value)
		}
	}
	return true
}

var (
	cachedNumberSets = map[int]map[int][][][]int{}
)

func numberSets(n, sn int) [][][]int {
	if cachedNumberSets[n][sn] != nil {
		return cachedNumberSets[n][sn]
	}
	sets := [][][]int{}

	for fixed := 0; fixed < n; fixed++ {
		var rowSet, colSet [][]int
		for moving := 0; moving < n; moving++ {
			rowSet = append(rowSet, []int{fixed, moving})
			colSet = append(colSet, []int{moving, fixed})
		}
		sets = append(sets, rowSet, colSet, sudokuSquareIndicesFromNumber(fixed, n, sn))
	}
	maths.Insert(cachedNumberSets, n, sn, sets)
	return sets
}

func (sb *sudokuBoard) Solve() bool {
	for changed := true; changed; {
		changed = false

		// See if any cell has only one option
		for row := 0; row < sb.n; row++ {
			for col := 0; col < sb.n; col++ {
				if sb.board[row][col].updateBoard(sb) {
					changed = true
				}
				if sb.broken {
					return false
				}
			}
		}

		// See if a cell in a row is the only one in the cell that can be a value
		for _, set := range numberSets(sb.n, sb.sn) {
			// map from number to idx in set.
			m := map[int][][]int{}
			for _, indices := range set {
				if cell := sb.board[indices[0]][indices[1]]; !cell.done() {
					for opt := range cell.options {
						m[opt] = append(m[opt], indices)
					}
				}
			}

			for v, possibilities := range m {
				if len(possibilities) == 1 {
					cell := sb.board[possibilities[0][0]][possibilities[0][1]]
					cell.options = map[int]bool{v: true}
					changed = true
				}
			}
		}
	}

	for row := 0; row < sb.n; row++ {
		for col := 0; col < sb.n; col++ {
			if !sb.board[row][col].done() {
				return false
			}
		}
	}
	return true
}

func (sb *sudokuBoard) Code() string {
	return sb.String()
}

func (sb *sudokuBoard) Done() bool {
	sb.Solve()
	return !sb.broken && sb.remaining == 0
}

func (sb *sudokuBoard) Distance() bfs.Int {
	return bfs.Int(sb.remaining)
}

func (sb *sudokuBoard) copy() *sudokuBoard {
	var newBoard [][]*sudokuCell
	for i := 0; i < sb.n; i++ {
		var row []*sudokuCell
		for j := 0; j < sb.n; j++ {
			row = append(row, sb.board[i][j].copy())
		}
		newBoard = append(newBoard, row)
	}
	return &sudokuBoard{sb.n, sb.sn, newBoard, sb.remaining, sb.broken, sb.guesses}
}

func (sb *sudokuBoard) AdjacentStates() []*sudokuBoard {
	// don't do more than 3 guesses
	if sb.guesses > 3 {
		return nil
	}

	var neighbors []*sudokuBoard
	for row := 0; row < sb.n; row++ {
		for col := 0; col < sb.n; col++ {
			cell := sb.board[row][col]
			if cell.done() {
				continue
			}
			if len(cell.options) == 2 {
				for k := range cell.options {
					new := sb.copy()
					new.board[row][col].options = map[int]bool{k: true}
					new.board[row][col].updateBoard(new)
					new.guesses++
					if !new.broken {
						neighbors = append(neighbors, new)
					}
				}
			}
		}
	}
	return neighbors
}
