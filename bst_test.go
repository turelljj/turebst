package turebst

import (
	"math/rand"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	tree := &bstree{}
	for i := 0; i < b.N; i++ {
		tree.Insert(rand.Int())
	}
}

func initBST() *bstree {
	t := &bstree{}
	for i := 0; i < 10000; i++ {
		t.Insert(rand.Int())
	}
	return t
}

func BenchmarkSort(b *testing.B) {
	tree := initBST()
	for i := 0; i < b.N; i++ {
		tree.Sort()
	}
}
