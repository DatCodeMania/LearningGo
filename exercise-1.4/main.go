package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	dupFiles := make([]map[string][]string, 0)
	for line, n := range counts {
		if n > 1 {
			filesContainingLine := make([]string, 0)
			for _, arg := range files {
				f, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
					continue
				}
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					if scanner.Text() == line {
						filesContainingLine = append(filesContainingLine, arg)
					}
				}
			}
			dupFiles = append(dupFiles, map[string][]string{line: filesContainingLine})
		}
	}
	for _, dict := range dupFiles {
		for line, files := range dict {
			fmt.Printf("Line: '%s' is duplicated in files: %v\n", line, files)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
