package btree

import (
	googlebtree "github.com/google/btree"
)

type BTree[T any] struct {
	goo googlebtree.BTreeG[T]
}

// func
