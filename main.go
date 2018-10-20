package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/adamryman/go_history/dedup"
)

// Purpose: create a history file that maintains most recent command usage, but
// delete duplicates
func main() {
	inputFile := flag.StringP("input", "i", filepath.Join(os.Getenv("HOME"), ".bash_history"), "file to remove leading duplicates from")
	replaceFile := flag.BoolP("replace", "r", false, "input replace file with output")
	outputFileName := flag.StringP("output", "o", "", "output file")
	_ = replaceFile

	flag.Parse()

	file, err := os.Open(*inputFile)
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

	var out io.Writer
	out = os.Stdout
	if *outputFileName != "" {
		out, err := os.Create(*outputFileName)
		if err != nil {
			panic(err)
		}
		defer out.Close()
	}

	// write this structure backwards to get the the correct ordering
	for _, v := range outData {
		_, err := fmt.Fprintf(out, "%s\n", v)
		if err != nil {
			panic(err)
		}
	}
}
