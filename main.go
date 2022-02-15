package main

import (
	"fmt"
	"goLearning/oslearning/file_system"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	fileSystem, err := file_system.NewFileSystem(cwd + "\\oslearning\\file_system\\file_example.json")
	if err != nil {
		fmt.Println(err)
	}
	fileSystem.Run()
}
