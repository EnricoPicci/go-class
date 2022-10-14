package generics

import "testing"

func TestBinaryTreeFindDoublePointer(t *testing.T) {
	e := Node[int]{nil, nil, 50}
	d := Node[int]{nil, nil, 27}
	c := Node[int]{&d, &e, 30}
	b := Node[int]{nil, nil, 2}
	a := Node[int]{&b, &c, 10}

	bt := BTree[int]{
		&a,
		func(targetVal int, nodeVal int) int {
			t.Logf("%v, %v\n", targetVal, nodeVal)
			return targetVal - nodeVal
		},
	}

	// test the find function
	var expectedHit *Node[int]
	var hit *Node[int]

	expectedHit = &a
	hit = *bt.findDoublePointer(10)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &b
	hit = *bt.findDoublePointer(2)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &c
	hit = *bt.findDoublePointer(30)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &d
	hit = *bt.findDoublePointer(27)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &e
	hit = *bt.findDoublePointer(50)
	checkFindResult(t, expectedHit, hit)

	// test that nil is returned if there is not the value we are searching for in the tree
	expectedHit = nil
	hit = *bt.findDoublePointer(49)
	checkFindResult(t, expectedHit, hit)
}

func TestBinaryTreeFindPointer(t *testing.T) {
	e := Node[int]{nil, nil, 50}
	d := Node[int]{nil, nil, 27}
	c := Node[int]{&d, &e, 30}
	b := Node[int]{nil, nil, 2}
	a := Node[int]{&b, &c, 10}

	bt := BTree[int]{
		&a,
		func(targetVal int, nodeVal int) int {
			t.Logf("%v, %v\n", targetVal, nodeVal)
			return targetVal - nodeVal
		},
	}

	// test the find function
	var expectedHit *Node[int]
	var hit *Node[int]

	expectedHit = &a
	hit = bt.findPointer(10)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &b
	hit = bt.findPointer(2)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &c
	hit = bt.findPointer(30)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &d
	hit = bt.findPointer(27)
	checkFindResult(t, expectedHit, hit)
	expectedHit = &e
	hit = bt.findPointer(50)
	checkFindResult(t, expectedHit, hit)

	// test that nil is returned if there is not the value we are searching for in the tree
	expectedHit = nil
	hit = bt.findPointer(49)
	checkFindResult(t, expectedHit, hit)
}

func checkFindResult(t *testing.T, expectedHit *Node[int], hit *Node[int]) {
	if expectedHit != hit {
		t.Errorf("expected %v, got %v", expectedHit, hit)
	}
}
