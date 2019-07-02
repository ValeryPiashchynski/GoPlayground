package main

import "testing"

func BenchmarkA_FooP(b *testing.B) {
	aa := aa{}
	a := &foos{
		a: "yuv has", b: "a small", c: "dick",
		d: struct {
			a string
			b string
			c string
		}{a: "really", b: "really", c: "small"},
		foo: aa,
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = FooP(a)
	}
}

func BenchmarkA_FooV(b *testing.B) {
	aa := aa{}
	a := foos{
		a: "yuv has", b: "a small", c: "dick",
		d: struct {
			a string
			b string
			c string
		}{a: "really", b: "really", c: "small"},
		foo: aa,
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = FooV(a)
	}
}
