package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/leep-frog/command"
)

const (
	N = "N"
)

func main() {
	command.RunNodes(node())
}

func node() *command.Node {
	return command.BranchNode(map[string]*command.Node{
		"1":  p1(),
		"2":  p2(),
		"3":  p3(),
		"4":  p4(),
		"5":  p5(),
		"6":  p6(),
		"7":  p7(),
		"8":  p8(),
		"9":  p9(),
		"10": p10(),
		"11": p11(),
	}, nil, true)
}

func ReadInput(problem int) string {
	b, err := ioutil.ReadFile(filepath.Join("input", fmt.Sprintf("p%d.txt", problem)))
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
	return string(b)
}

func ReadLines(problem int) []string {
	rs := strings.Split(ReadInput(problem), "\n")
	for idx := range rs {
		rs[idx] = strings.TrimSpace(rs[idx])
	}
	return rs
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to convert string to int: %v", err)
	}
	return i
}

func IsSquare(i int) bool {
	rt := int(math.Sqrt(float64(i)))
	return rt*rt == i
}
