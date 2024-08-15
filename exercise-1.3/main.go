package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	// Took 0.000022 seconds to execute.
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	// Took 0.000002 seconds to execute.
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	// Took 0.000002 seconds to execute.
	fmt.Println((strings.Join(os.Args[0:], " ")))
}

func execEcho(function func()) {
	start := time.Now()
	function()
	fmt.Printf("Took %f seconds to execute.\n", time.Since(start).Seconds())
}

func main() {
	execEcho(echo1)
	execEcho(echo2)
	execEcho(echo3)
}
