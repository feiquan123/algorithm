#! /bin/bash

echo "main.c out:"
gcc main.c && ./a.out && rm ./a.out

echo  "main.go out:"
go run main.go
