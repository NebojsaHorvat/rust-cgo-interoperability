package main

/*
#cgo LDFLAGS: -L. -llib

extern int getStateWrapper(int);
extern int applyStateTransition(int, int(*)(int));

typedef void (*closure)();
*/
import "C"
import "fmt"

//export getState
func getState(index C.int) C.int {
	fmt.Println("Function getStateWrapper from Go called!")
	if index == 3 {
		return 1337
	}
	return 0
}

func main() {
	fmt.Println("Calling Rust function from Go...")
	result := C.applyStateTransition(3, C.closure(C.getStateWrapper))
	fmt.Printf("Result: %d\n", result)
}
