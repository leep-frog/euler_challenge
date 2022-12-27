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

func FullPath(fileParts ...string) string {
	f := []string{fullPath(fileParts[0], 2)}
	if len(fileParts) > 1 {
		f = append(f, fileParts[1:]...)
	}
	return filepath.Join(f...)
}

func fullPath(filename string, depth int) string {
	_, thisFile, _, ok := runtime.Caller(depth)
	if !ok {
		log.Fatalf("failed to get runtime info")
	}
	return filepath.Join(path.Dir(thisFile), filename)
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

func Exists(f string) bool {
	_, err := os.Stat(fullPath(f, 2))
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("failed to check file info: %v", err)
	}
	return err == nil
}

func Touch(f string) {
	fp := fullPath(f, 2)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if err := ioutil.WriteFile(fp, nil, 0644); err != nil {
			log.Fatalf("failed to touch file %s %s: %v", f, fp, err)
		}
	} else if err != nil {
		log.Fatalf("failed to check file: %v", err)
	}
}

func Write(f, contents string) {
	fp := fullPath(f, 2)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if err := ioutil.WriteFile(fp, []byte(contents), 0644); err != nil {
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

func IntsToStrings(is []int) []string {
	var r []string
	for _, i := range is {
		r = append(r, strconv.Itoa(i))
	}
	return r
}
