package oslearning

import (
	"testing"
)

func TestFileSystem(t *testing.T) {
	fileSystem, _ := NewFileSystem()
	fileSystem.ShowUserName()
}
