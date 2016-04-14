package hellodev

import (
	"os"
	"testing"
)

const dev = "/dev/hello"

func TestOpenClose(t *testing.T) {
	for i := 0; i < 10; i++ {
		f, err := os.Open(dev)
		if err != nil {
			t.Fatal(err)
		}
		err2 := f.Close()
		if err2 != nil {
			t.Fatal(err2)
		}
	}
}
