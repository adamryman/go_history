package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/adamryman/go_history/dedup"
)

// Purpose: create a history file that maintains most recent command usage, but
// delete duplicates
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
	file.Close()

	outData := dedup.Leading(fileData)

	outFile, err := os.Create(*fileName + ".go_history")
	if err != nil {
		panic(err)
	}

	// write this structure backwards to get the the correct ordering
	for _, v := range outData {
		_, err := fmt.Fprintf(outFile, "%s\n", v)
		if err != nil {
			panic(err)
		}
	}
	outFile.Close()
	// print a move command to stdout
	fmt.Fprintln(os.Stdout, "mv "+*fileName+".go_history"+" "+*fileName)
}
