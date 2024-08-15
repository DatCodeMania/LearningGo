package main

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

/*
What happens here is the 64-bit unsigned integer(x) is broken down into 8 bytes, which makes sense because 8 * 8 is 64 and a byte is 8 bits
The way this works is the 8 bytes are shifted by n*8, e.g. shifted by 0 for n=0 etc.
Then, due to the way the byte function works, the rightmost byte is extracted.
This basically goes over every byte out of 8, and extracts the rightmost one every time after shifting.

For a deeper explanation read the PopCountShift comment
*/
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var bytes int
	for i := 0; i < 8; i++ {
		bytes += int(pc[byte(x>>(i*8))])
	}
	return bytes
}

/*
Basically, whats going on here is we loop through the whole 64-bit unsigned integer(x)
With each iteration, we check if the right most bit (e.g. 1 in 0001) is == 1
If so, we increment the count integer by 1
After this, we shift the bits to the right, meaning that what was the pre-rightmost bit (e.g. 1 in 0010) becomes the rightmost bit.
The current rightmost bit (e.g. 1 in 0001) is discarded, and the leftmost bit is set to 0.

This way our function is able to check every single bit out of the 64 passed to it.

Bitwise operator explanation:
& : AND - this operation (in our case) only cares about the rightmost bit, due to the 1 in x&1, this means that the 0th bit is checked; the 0th bit is the rightmost. It compares e.g. 0001 (mask) and 0111 (x), checking all of the bits, but in our case only the last one because that is what we specify in our mask.

0001 - mask
0111 - x
----
0001 - comparison returns 1, which for us is true - the last bits are identical.

x >>= 1 : shorthand for x = x >> 1, all this does is shift the bits to the right, this is further explained above.
*/
func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ { // loop 64 times (to loop thru x)
		if x&1 == 1 { // check if the rightmost bit == 1
			count++
		}
		x >>= 1 // shift the bits to the right
	}
	return count
}

/*
x&(x-1) clears the rightmost non-zero bit of x.
Using this information, we are able to count the number of set bits in x by counting the number of times we can clear the rightmost bit.
*/
func PopCountClear(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(0x1234567890ABCDEF)
	}
}
