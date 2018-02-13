package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

// Purpose: create a history file that maintains most recent command usage, but
// delete duplicates
// Implementation:
// 10:
// pop off bottom line
// store that line on the top of a different file
// delete all lines that match that line from the file
// goto 10
func main() {
	fileName := flag.StringP("file", "f", filepath.Join(os.Getenv("HOME"), ".bash_history"), "")

	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// scan the .bash_history into a string slice
	var fileData []string
	for scanner.Scan() {
		// strip leading and trailing whitespace
		fileData = append(fileData, strings.TrimSpace(scanner.Text()))
	}

	// looping through backwards is sort of like popping off the bottom
	var reversedNewFileData []string
	var check checker
	for i := len(fileData) - 1; i >= 0; i = i - 1 {
		line := fileData[i]
		if check.IsDup(line) {
			continue
		}
		reversedNewFileData = append(reversedNewFileData, line)
	}
	file.Close()

	fmt.Fprintf(os.Stderr, "%d\n", len(reversedNewFileData))

	out, err := os.Create(*fileName + ".go_history")
	if err != nil {
		panic(err)
	}

	// write this structure backwards to get the the correct ordering
	for i := len(reversedNewFileData) - 1; i >= 0; i = i - 1 {
		// need to put those new lines back in
		_, err := fmt.Fprintf(out, "%s\n", reversedNewFileData[i])
		if err != nil {
			panic(err)
		}
	}
	out.Close()
	// print a move command to stdout
	fmt.Fprintln(os.Stdout, "mv "+*fileName+".go_history"+" "+*fileName)
}

type checker struct {
	dups map[string]int
}

func (c *checker) IsDup(s string) bool {
	if c.dups == nil {
		c.dups = make(map[string]int)
	}

	if c.dups[s] == 0 {
		// set this so that it will be a duplicate in the future
		c.dups[s] = 1
		return false
	}

	// counting how often a line happens might be interesting to keep around later
	c.dups[s] = c.dups[s] + 1
	return true
}
