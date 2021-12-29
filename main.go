package main

import (
	"goLearning/oslearning"
)

func main() {
	fileSystem, _ := oslearning.NewFileSystem()
	fileSystem.ShowUserName()
}
