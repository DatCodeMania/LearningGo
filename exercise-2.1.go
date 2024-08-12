package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroK Fahrenheit = -459.67
	FreezingK     Fahrenheit = 32
	BoilingK      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
