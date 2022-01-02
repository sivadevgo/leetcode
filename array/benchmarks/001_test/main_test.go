package main

import "testing"

func BenchmarkTest1(b *testing.B) {
	a := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		count := 0
		for i := 0; i < len(a); i++ {
			count += len(a)
		}
	}
}

func BenchmarkTest2(b *testing.B) {
	a := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		count := 0
		lena := len(a)
		for i := 0; i < lena; i++ {
			count += lena
		}
	}
}
