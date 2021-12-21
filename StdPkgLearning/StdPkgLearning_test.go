package StdPkgLearning

import (
	"fmt"
	"io"
	"testing"
)

func TestAlphaReader(t *testing.T) {
	reader := NewAlphaReader("hhHello Go U r so f**king weired!")
	p := make([]byte, 4)
	for {
		length, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:length]))
	}

}
