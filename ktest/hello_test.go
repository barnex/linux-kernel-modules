package ktest

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"testing/iotest"
)

const module = "hello"

func init() {
	Rmmod(module)
	if err := Insmod(module); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func TestRmmod(t *testing.T) {
	if err := Rmmod(module); err != nil {
		t.Fatal(err)
	}
	if HaveModule(module) {
		t.Fatalf("rmmod %v: module still in /proc/modules", module)
	}

	if err := Insmod(module); err != nil {
		t.Fatal(err)
	}
	if !HaveModule(module) {
		t.Fatalf("insmod %v: module not in /proc/modules", module)
	}
}

const dev = "/dev/hello"
const want = "Hello from the linux kernel!\n"

// Open and close device several times.
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

// Open while open should fail.
func TestBusy(t *testing.T) {
	f1, err1 := os.Open(dev)
	if err1 != nil {
		t.Fatal(err1)
	}

	f2, err2 := os.Open(dev)
	if err2 != nil {
		t.Fatalf("open %v: %v", dev, err2)
	}
	defer f2.Close()

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

// Read all bytes from device.
func TestReadAll(t *testing.T) {
	f, err := os.Open(dev)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	have := string(bytes)
	if have != want {
		t.Fatalf("read %v: got %q, want %q", dev, have, want)
	}
}

// Read all bytes, one at a time.
func TestReadOneByte(t *testing.T) {
	f, err := os.Open(dev)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(iotest.OneByteReader(f))
	if err != nil {
		t.Fatal(err)
	}
	have := string(bytes)
	if have != want {
		t.Fatalf("read %v: got %q, want %q", dev, have, want)
	}
}

// Read all bytes, in small chunks.
func TestReadHalf(t *testing.T) {
	f, err := os.Open(dev)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(iotest.HalfReader(f))
	if err != nil {
		t.Fatal(err)
	}
	have := string(bytes)
	if have != want {
		t.Fatalf("read %v: got %q, want %q", dev, have, want)
	}
}

// Close before done reading.
func TestEarlyClose(t *testing.T) {
	f, err := os.Open(dev)
	if err != nil {
		t.Fatal(err)
	}
	f.Read(make([]byte, 2))
	f.Close()

	f, err = os.Open(dev)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(iotest.HalfReader(f))
	if err != nil {
		t.Fatal(err)
	}

	have := string(bytes)
	if have != want {
		t.Fatalf("read %v: got %q, want %q", dev, have, want)
	}
}

// Try to corrupt memory by reading more than available.
func TestCorruption(t *testing.T) {
	f, err := os.Open(dev)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	buf := make([]byte, 255)
	for i := range buf {
		buf[i] = byte(i)
	}

	f.Read(buf)
	for i := len(want); i < len(buf); i++ {
		if buf[i] != byte(i) {
			t.Errorf("read %v: userspace data corruption: byte[%v] = %v", dev, i, buf[i])
		}
	}
}
