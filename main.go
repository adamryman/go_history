package main

import (
	"bufio"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	// TODO: Read file
	fileName := flag.StringP("file", "f", "", "")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var fileData []string
	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}

	var revNewFileData []string
	dups := make(map[string]int)
	for i := len(fileData) - 1; i >= 0; i = i - 1 {
		line := fileData[i]
		if dups[line] != 0 {
			dups[line] = dups[line] + 1
			//fmt.Println(line)
			continue
		}
		dups[line] = 1
		revNewFileData = append(revNewFileData, line)
	}
	fmt.Printf("%d\n", len(revNewFileData))

	file.Close()

	out, err := os.Create(*fileName + ".go_history")
	if err != nil {
		panic(err)
	}

	for i := len(revNewFileData) - 1; i >= 0; i = i - 1 {
		//for _, line := range revNewFileData {
		_, err := fmt.Fprintf(out, "%s\n", revNewFileData[i])
		if err != nil {
			panic(err)
		}
	}
	out.Close()
	fmt.Println("done")

	// 10
	// pop off bottom line
	// store that line on the top of a different file
	// delete all lines that match that line from the file
	// goto 10

}
