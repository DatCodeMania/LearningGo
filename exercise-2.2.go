package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	metersToFeet    float64 = 3.28084
	KGsToLBS        float64 = 2.20462
	litersToGallons float64 = 0.264172
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			inp, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "converter: %v\n", err)
				os.Exit(1)
			}
			converter(inp)
		}
	} else {
		fmt.Println("No CLI arguments provided, reading from stdin:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			inp, err := strconv.ParseFloat(line, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "converter: %v\n", err)
				os.Exit(1)
			}
			converter(inp)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "err reading from stdin: %v\n", err)
			os.Exit(1)
		}
	}
}

func converter(num float64) {
	meters := num
	feet := meters * metersToFeet

	kgs := num
	lbs := kgs * KGsToLBS

	liters := num
	gallons := liters * litersToGallons

	fmt.Printf("Meters: %.2f\tFeet: %.2f\nKilograms: %.2f\tPounds: %.2f\nLiters: %.2f\tGallons: %.2f\n", meters, feet, kgs, lbs, liters, gallons)
}
