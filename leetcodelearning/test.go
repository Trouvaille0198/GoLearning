package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		if len(nums) == 0 {
			break
		}
		res := 0
		for _, num := range nums {
			tmp, _ := strconv.Atoi(num)
			res += tmp
		}
		fmt.Println(res)
	}
}
