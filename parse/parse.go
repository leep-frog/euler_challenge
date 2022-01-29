package parse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to convert string to int: %v", err))
	}
	return i
}

func Itos(i int) string {
	return fmt.Sprintf("%d", i)
}

func fullPath(filename string, depth int) string {
	_, thisFile, _, ok := runtime.Caller(depth)
	if !ok {
		log.Fatalf("failed to get runtime info")
	}
	return filepath.Join(path.Dir(thisFile), "input", filename)
}

func ReadFileInput(f string) string {
	return readFileInput(f)
}

func readFileInput(f string) string {
	fp := fullPath(f, 3)
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
	return string(b)
}

func Touch(f string) {
	fp := fullPath(f, 2)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if err := ioutil.WriteFile(fp, nil, 0644); err != nil {
			log.Fatalf("failed to write to file: %v", err)
		}
	} else if err != nil {
		log.Fatalf("failed to check file: %v", err)
	}
}

func ReadFileLines(f string) []string {
	rs := strings.Split(readFileInput(f), "\n")
	for idx := range rs {
		rs[idx] = strings.TrimSpace(rs[idx])
	}
	return rs
}

func ToGrid(lines []string) [][]int {
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, str := range strings.Split(line, ",") {
			row = append(row, Atoi(str))
		}
		grid = append(grid, row)
	}
	return grid
}
