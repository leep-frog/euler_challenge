package y2018

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

type mdNode struct {
	metadata   []int
	childNodes []*mdNode
}

func (md *mdNode) populate(numberList *linkedlist.Node[int], numChildren, numMetadata int) *linkedlist.Node[int] {
	for i := 0; i < numChildren; i++ {
		subChildren := numberList.Value
		numberList = numberList.Next
		numMD := numberList.Value
		numberList = numberList.Next

		newMD := &mdNode{}
		numberList = newMD.populate(numberList, subChildren, numMD)
		md.childNodes = append(md.childNodes, newMD)
	}

	for i := 0; i < numMetadata; i++ {
		md.metadata = append(md.metadata, numberList.Value)
		numberList = numberList.Next
	}

	return numberList
}

func (md *mdNode) mdSum() int {
	sum := bread.Sum(md.metadata)
	for _, cn := range md.childNodes {
		sum += cn.mdSum()
	}
	return sum
}

func (md *mdNode) mdSum2() int {
	if len(md.childNodes) == 0 {
		return bread.Sum(md.metadata)
	}

	var sum int
	for _, idx := range md.metadata {
		idx--
		if idx < 0 || idx >= len(md.childNodes) {
			continue
		}
		sum += md.childNodes[idx].mdSum2()
	}
	return sum
}

func (d *day08) Solve(lines []string, o command.Output) {
	numbers := linkedlist.NewList(parse.AtoiArray(strings.Split(lines[0], " "))...)

	root := &mdNode{}
	numChildren := numbers.Value
	numbers = numbers.Next
	numMD := numbers.Value
	numbers = numbers.Next

	root.populate(numbers, numChildren, numMD)
	o.Stdoutln(root.mdSum(), root.mdSum2())
}

func (d *day08) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"138 66",
			},
		},
		{
			ExpectedOutput: []string{
				"49180 20611",
			},
		},
	}
}
