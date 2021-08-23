package utils

import "testing"

var sli = []string{"hello", "hello", "golang", "goland", "hello", "goland"}

func TestSliceUnique(t *testing.T) {
	res := sliceUnique(sli)
	t.Log(res)
}

func BenchmarkSliceUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceUnique(sli)
	}
}