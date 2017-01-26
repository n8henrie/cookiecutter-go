package main

import (
	"testing"
)

func TestPackage(t *testing.T) {

	// Structure for single data point, e.g.
	// type testpair struct {
	// 	num int
	// 	str string
	// 	}
	type test struct {
	}

	// Array of test, e.g.
	// var tests = []{test{
	// 	{12, "twelve"},
	// 	{9, "four"},
	// }
	var tests = []test{}

	for _, point := range tests {
		// n := testMyFunction(point.num)
		// if n != pair.dsum {
		// 	t.Error("For", pair.num, "expected", pair.dsum, "got", n)
		// }
	}
}

func BenchmarkPackage(b *testing.B) {
	// run the Foo function b.N times
	// NB: do not use n or b.N in the loop
	for n := 0; n < b.N; n++ {
		//	Foo(10)
	}
}
