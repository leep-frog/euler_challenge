package parse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/leep-frog/euler_challenge/functional"
)

var (
	numberRegex = regexp.MustCompile("^[0-9]+$")
)

func IsNumberFormat(s string) bool {
	return numberRegex.MatchString(s)
}

func Atoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(fmt.Sprintf("failed to convert string to int: %v", err))
	}
	return i
}

func AtoiArray(s []string) []int {
	return functional.Map(s, func(v string) int {
		return Atoi(v)
	})
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
	return rs
}

func MapToGrid[T any](lines []string, m map[rune]T) [][]T {
	var grid [][]T
	for _, line := range lines {
		var row []T
		for _, c := range line {
			v, ok := m[c]
			if !ok {
				log.Fatalf("Unknown grid key: %c", c)
			}
			row = append(row, v)
		}
		grid = append(grid, row)
	}
	return grid
}

func PrintAOCGrid[T comparable](grid [][]T, m map[T]rune) {
	for _, row := range grid {
		for _, t := range row {
			fmt.Printf("%c", m[t])
		}
		fmt.Println()
	}
}

func AOCGrid[T any](lines []string, dotValue, hashtagValue T) [][]T {
	return MapToGrid(lines, map[rune]T{
		'.': dotValue,
		'#': hashtagValue,
	})
}

func ToGrid(lines []string, separator string) [][]int {
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, str := range strings.Split(line, separator) {
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

func Split(lines []string, delimiter string) [][]string {
	var r [][]string
	var cur []string
	for _, line := range lines {
		if line == delimiter {
			r = append(r, cur)
			cur = nil
		} else {
			cur = append(cur, line)
		}
	}
	return append(r, cur)
}

func ToCharArray(s string) []rune {
	var r []rune
	for _, c := range s {
		r = append(r, c)
	}
	return r
}
