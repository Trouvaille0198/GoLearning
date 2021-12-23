package main

import (
	"fmt"
	"goLearning/osLearning"
)

func main() {
	add1 := osLearning.GetRandomAddresses(255, 4)
	add2 := osLearning.GetAddressPages(add1, 2)
	fmt.Println(add1[0], add2[0])
}
