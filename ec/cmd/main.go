package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/command/sourcerer"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
	"github.com/leep-frog/euler_challenge/rgx"
)

func main() {
	os.Exit(sourcerer.Source([]sourcerer.CLI{&ecCLI{}}))
}

type ecCLI struct{}

func (ecCLI) Name() string    { return "ec" }
func (ecCLI) Changed() bool   { return false }
func (ecCLI) Setup() []string { return nil }

func (ecCLI) Node() command.Node {
	return &commander.BranchNode{
		Branches:          eulerchallenge.Branches(),
		DefaultCompletion: true,
		Default:           n(),
		// Default:           eulerchallenge.FileGenerator(),
	}
}

func n() command.Node {

	r := rgx.New("^(p[0-9]+)_(.*).go$")
	return commander.SerialNodes(
		commander.SimpleProcessor(func(i *command.Input, o command.Output, d *command.Data, ed *command.ExecuteData) error {

			_, thisFile, _, ok := runtime.Caller(0)
			if !ok {
				return o.Stderr("failed to run runtime.Caller")
			}
			ecDir := filepath.Dir(filepath.Dir(thisFile))

			return o.Err(filepath.WalkDir(ecDir, func(path string, d fs.DirEntry, err error) error {
				fmt.Println(path)
				if d.IsDir() {
					return nil
				}

				m, ok := r.Match(filepath.Base(path))
				if !ok {
					return nil
				}
				pNumber := m[0]

				b, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("failed to read file: %v", err)
				}
				contents := strings.Split(string(b), "\n")
				if contents[0] != "package eulerchallenge" {
					return nil
				}
				contents[0] = fmt.Sprintf("package %s", pNumber)
				if err := os.WriteFile(path, []byte(strings.Join(contents, "\n")), 0644); err != nil {
					return fmt.Errorf("failed to write file: %v, err")
				}

				fmt.Println("COMPLETED", pNumber)

				return nil
			}))
		}, nil),
	)
}
