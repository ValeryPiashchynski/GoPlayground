// Code generated by cmd/cgo; DO NOT EDIT.

//line main.go:1:1
package main

import _ "unsafe"

import (
	"sync"
)

var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int {
	return a + b
}

func main() {

}
