#! /bin/bash
set -e
astyle --indent=tab --style=java hello.c
make
echo OK
