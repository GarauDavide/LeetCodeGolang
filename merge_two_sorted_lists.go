package main

import "fmt"

type numeri struct {
	num  int
	next *numeri
}

func main() {

	var l1 *numeri = nil
	var l2 *numeri = nil

	addTail(&l1, 1)
	addTail(&l1, 4)
	addTail(&l1, 7)
	addTail(&l1, 7)
	addTail(&l1, 9)

	addTail(&l2, 1)
	addTail(&l2, 3)
	addTail(&l2, 5)
	addTail(&l2, 6)
	addTail(&l2, 8)

	newList := mergeLists(&l1, &l2)

	fmt.Println("New list")
	for index := newList; index != nil; index = index.next {

		fmt.Println("Val:", index.num)
	}
}

func addFront(l **numeri, n int) {
	newElem := numeri{num: n}
	newElem.next = (*l)
	(*l) = &newElem
}

func addTail(l *(*numeri), n int) {

	if (*l) == nil {
		addFront(&(*l), n)
		return
	}

	newElem := numeri{num: n}

	index := (*l)
	for ; index.next != nil; index = index.next {
	}

	newElem.next = index.next
	index.next = &newElem
}

func mergeLists(l1 **numeri, l2 **numeri) *numeri {

	list1 := (*l1)
	list2 := (*l2)

	var newList *numeri = nil

	idx1 := list1
	idx2 := list2
	idxN := newList

	for idx1 != nil || idx2 != nil {

		nE := numeri{next: nil}

		if newList == nil {
			newList = &nE
			idxN = &nE
		} else {
			idxN.next = &nE
			idxN = idxN.next
		}

		if idx1 == nil && idx2 != nil {
			nE.num = idx2.num
			idx2 = idx2.next
		} else if idx1 != nil && idx2 == nil {
			nE.num = idx1.num
			idx1 = idx1.next
		} else if idx1.num < idx2.num {
			nE.num = idx1.num
			idx1 = idx1.next
		} else if idx1.num > idx2.num {
			nE.num = idx2.num
			idx2 = idx2.next
		} else {
			nE.num = idx1.num
			/* create the second equal element */
			nEequ := numeri{num: idx2.num, next: nil}
			idxN.next = &nEequ
			idxN = idxN.next

			idx1 = idx1.next
			idx2 = idx2.next
		}
	}

	return newList
}

func mergeKSortedList(kList []*numeri) *numeri {
	var nlist int = len(kList)
	var allList *numeri = nil

	if nlist == 0 {
		return allList
	}

	if nlist == 1 {
		return kList[0]
	}

	firts := 0
	other := 1
	allList = kList[firts]

	for other < nlist {

		allList = mergeLists(&allList, &kList[other])
		other++
	}

	return allList
}
