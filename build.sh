#! /bin/bash
set -e
astyle --indent=tab --style=java hello.c
make
rm -f ./ktest/*.test
(cd ktest && go test -c)
sudo ./ktest/ktest.test
echo OK
