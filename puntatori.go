package main

import "fmt"

type numeri struct {
	num  int
	next *numeri
}

func main() {

	/* prova lista go */
	var l *numeri = nil

	addTail(&l, 12)
	addTail(&l, 13)
	addTail(&l, 14)

	addHead(&l, 1)
	addHead(&l, 2)
	addHead(&l, 3)

	for index := l; index != nil; index = index.next {
		fmt.Println(index.num)
	}
}

func addHead(l *(*numeri), n int) {

	newElem := numeri{num: n}

	newElem.next = (*l)
	(*l) = &newElem
}

func addTail(l *(*numeri), n int) {

	if (*l) == nil {
		addHead(&(*l), n)
		return
	}

	newElem := numeri{num: n}

	index := (*l)
	for ; index.next != nil; index = index.next {
	}

	newElem.next = index.next
	index.next = &newElem
}
