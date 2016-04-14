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

func TestBusy(t *testing.T) {
	f1, err1 := os.Open(dev)
	if err1 != nil {
		t.Fatal(err1)
	}

	f2, err2 := os.Open(dev)
	if f2 != nil || err2 == nil {
		t.Fatalf("open %v: expected nil, EBUSY, got: %v, %v", dev, f2, err2)
	}
	if err := f1.Close(); err != nil {
		t.Fatal(err)
	}

	f3, err3 := os.Open(dev)
	if err3 != nil {
		t.Fatal(err3)
	}
	if err := f3.Close(); err != nil {
		t.Fatal(err)
	}
}
