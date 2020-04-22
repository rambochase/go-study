package benchmark

import (
	"bytes"
	"testing"
)

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "rambo", "chase", "love", "peace", ",", "you know!"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "rambo", "chase", "love", "peace", ",", "you know!"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
