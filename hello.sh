#! /bin/bash
module=hello
set -e
astyle --indent=tab --style=java $module.c
make clean
make
rm -f ./ktest/*.test
(cd ktest && go test -c)
sudo ./ktest/ktest.test
echo OK
