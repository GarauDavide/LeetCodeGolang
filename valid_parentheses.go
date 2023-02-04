package main

import (
	"container/list"
	"fmt"
)

func main() {

	fmt.Println(isValid("()"))
}

func isValid(s string) bool {

	strLen := len(s)
	associations := map[string]string{")": "(", "]": "[", "}": "{"}
	openList := list.New()

	for i := 0; i < strLen; i++ {
		str := string(s[i])

		if str == "(" || str == "[" || str == "{" {
			openList.PushFront(str)
		} else {
			v := openList.Front()

			if v == nil || associations[str] != v.Value {
				return false
			} else {
				openList.Remove(openList.Front())
			}
		}
	}

	return openList.Len() == 0
}
