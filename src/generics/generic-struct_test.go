package generics

import (
	"testing"
)

func TestNodeStruct(t *testing.T) {
	left := node[int]{5, nil, nil}
	right := node[int]{20, nil, nil}
	val := node[int]{10, &left, &right}

	// test the node value
	if val.letf != &left {
		t.Errorf("the left node %v is not what is expected which is %v", val.letf, &left)
	}
	if *val.letf != left {
		t.Errorf("the left node %v is not what is expected which is %v", *val.letf, left)
	}
}
