package main

/*
#cgo CFLAGS: -g -Wall
extern void inCFile();
*/
import "C"
import "fmt"

func main() {
	fmt.Println("I am in Go code now!")
	C.inCFile()
}
