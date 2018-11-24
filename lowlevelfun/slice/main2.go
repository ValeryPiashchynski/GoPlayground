package main

import (
	"fmt"
	_ "unsafe"
)

//go:linkname zerobase runtime.zerobase
var zerobase uintptr

//go:linkname now zerobase.now
func now() string

func main() {

	println(now())

	var s struct{}
	var a [42]struct{}

	fmt.Printf("zerobase = %p\n", &zerobase)
	fmt.Printf("       s = %p\n", &s)
	fmt.Printf("       a = %p\n", &a)
}
