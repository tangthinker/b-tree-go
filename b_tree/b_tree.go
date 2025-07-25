package b_tree

type Node[T comparable] struct {
	kvs      []KV[T]
	children []*Node[T]
}

type KV[T comparable] struct {
	K T
	V any
}

type BTree[T comparable] struct {
	root *Node[T]
	m    int
}

func New[T comparable](m int) *BTree[T] {
	return &BTree[T]{
		m: m,
	}
}

func (t *BTree[T]) Insert(k T, v any) {
	if t.root == nil {
		t.root = &Node[T]{
			kvs: []KV[T]{
				{k, v},
			},
		}
		return
	}
	if t.root.children == nil {

	}

}

func (t *BTree[T]) Delete(k T) {

}

func (t *BTree[T]) Search(k T) any {
	return nil
}
