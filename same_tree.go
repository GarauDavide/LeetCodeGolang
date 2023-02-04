package main

type tree struct {
	num   int
	left  *tree
	right *tree
}

func main() {

	var root *tree = nil

	insertElement(&root, 10)
}

func insertElement(t *(*tree), n int) {

	elem := tree{num: n}

	if (*t) == nil {
		elem.left = nil
		elem.right = nil
		(*t) = &elem
	}
}

func stampTree(t *(*tree)) {

	if (*t) == nil {
		return
	}
}
