package main

import (
	"fmt"
	"goLearning/oslearning"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	fileSystem, err := oslearning.NewFileSystem(cwd + "\\oslearning\\file_example.json")
	if err != nil {
		fmt.Println(err)
	}
	fileSystem.Run()
}
