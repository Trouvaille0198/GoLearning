package oslearning

import (
	"fmt"
	"testing"
)

func TestAlphaReader(t *testing.T) {
	virtualMemorySize := 32
	addressNum := 255
	fmt.Printf(
		"%-10s %-12s %-10s %-10s %-10s \n",
		"PageSize", "MemorySize", "OPT", "LRU", "FIFO")
	for i := 0; i < 7; i++ {
		internalSize := (i + 1) * 4 // 每次翻四倍
		for j := 0; j < 8; j++ {
			pageSize := j + 1
			r := NewPageReplacer(
				addressNum, virtualMemorySize, pageSize, internalSize)
			r.DisplayResult()
		}
		fmt.Println()
	}
}
