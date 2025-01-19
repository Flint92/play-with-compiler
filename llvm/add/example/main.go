package main

/*
#cgo LDFLAGS: -L. -ladd_fn
extern int addFn(int a, int b);
*/
import "C"
import "fmt"

func main() {
	result := C.addFn(1, 2)
	fmt.Println("Result from addFn:", result)
}
