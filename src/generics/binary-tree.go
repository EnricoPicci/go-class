package generics

type Node[T any] struct {
	left, right *Node[T]
	data        T
}

type BTree[T any] struct {
	root    *Node[T]
	compare func(T, T) int
}

func (bt *BTree[T]) findDoublePointer(v T) **Node[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmpRes := bt.compare(v, (*pl).data); {
		case cmpRes < 0:
			pl = &(*pl).left
		case cmpRes > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

func (bt *BTree[T]) findPointer(v T) *Node[T] {
	pl := bt.root
	for pl != nil {
		switch cmpRes := bt.compare(v, (*pl).data); {
		case cmpRes < 0:
			pl = (*pl).left
		case cmpRes > 0:
			pl = (*pl).right
		default:
			return pl
		}
	}
	return nil
}
