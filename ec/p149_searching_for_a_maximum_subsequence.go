package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func bigSquare() [][]int {
	var s []int
	for i := 0; i < 2000*2000; i++ {
		if i < 55 {
			v := 100003 - 200003*(i+1) + 300007*(maths.Pow(i+1, 3))
			s = append(s, (v%1000000)-500000)
		} else {
			v := s[i-24] + s[i-55] + 1000000
			s = append(s, (v%1000000)-500000)
		}
	}
	var square [][]int
	for i := 0; i < 2000; i++ {
		square = append(square, s[i*2000:i*2000+2000])
	}
	return square
}

func smallSquare() [][]int {
	return [][]int{
		{-2, 5, 3, 2},
		{9, -6, 5, 1},
		{3, 2, 7, 3},
		{-1, 8, -4, 8},
	}
}

func P149() *problem {
	return intInputNode(149, func(o command.Output, n int) {

		var square [][]int
		if n > 1 {
			square = bigSquare()
		} else {
			square = smallSquare()
		}

		best := maths.Largest[int, int]()
		LN := len(square)

		for i, curRow := range square {
			var curCol, upRightSeq, downRightSeq, seq3, seq4 []int
			for j := range square {
				curCol = append(curCol, square[j][i])
			}
			/*for row, col := i, 0; row < i+LN; row, col = row+1, col+1 {
				upRightSeq = append(upRightSeq, square[row%LN][col%LN])
				downRightSeq = append(downRightSeq, square[(2*LN-1-row)%LN][col%LN])
				seq3 = append(seq3, square[(2*LN-1-row)%LN][(2*LN-1-col)%LN])
				seq4 = append(seq4, square[row%LN][(2*LN-1-col)%LN])
			}*/
			for row, col := i, 0; row >= 0; row, col = row-1, col+1 {
				upRightSeq = append(upRightSeq, square[row][col])
				downRightSeq = append(downRightSeq, square[LN-1-row][col])
				seq3 = append(seq3, square[LN-1-row][LN-1-col])
				seq4 = append(seq4, square[row][LN-1-col])
			}
			best.Check(linearLargestSequence(curRow))
			best.Check(linearLargestSequence(curCol))
			best.Check(linearLargestSequence(upRightSeq))
			best.Check(linearLargestSequence(downRightSeq))
			best.Check(linearLargestSequence(seq3))
			best.Check(linearLargestSequence(seq4))
		}
		o.Stdoutln(best.Best())
	}, []*execution{
		{
			args:     []string{"2"},
			want:     "52852124",
			estimate: 1,
		},
		{
			args: []string{"1"},
			want: "16",
		},
	})
}

func linearLargestSequence(sequence []int) int {
	var sum, max int
	for _, v := range sequence {
		sum = maths.Max(0, sum+v)
		max = maths.Max(max, sum)
	}
	return max
}

/*func bruteLargestSequence(sequence []int, best *maths.Bester[int, int]) {
	for i := range sequence {
		var sum int
		for j := i; j < len(sequence); j++ {
			sum += sequence[j]
			best.Check(sum)
		}
	}
}*/
