# linux-kernel-modules

Fun with loadable kernel modules

## hellodev

hellodev.c implements the classical /dev/hello, that outputs `Hello from the linux kernel!` each time it is read.

test/hellodev_test.go has fairly extensive tests illustrating how many things can go wrong, even in such simple module. E.g.:
  - TestBusy: make sure the same device can't be opened twice. Failing to do so will cause concurrent reads to step onto each other's toes.
  - TestReadAll: make sure we read the expected message.
  - TestReadOneByte: read the message, but one byte at a time. One failing case returned `HHHHHHHHHHHHHHHHHHHHHHHHHHH`
  - TestReadHalf: read the message, but in chunks at most half of the total message size.
  - TestEarlyClose: read 2 bytes of the message, close the file, open and read again. One failing case returend `llo from the linux kernel!`
  - TestCorruption: read into a buffer with extra space at the end, and make sure that space does not get overwritten. Intentionally causing the device to copy too many bytes to userspace causes this test to fail (and also causes most other tests to segfault).
  
