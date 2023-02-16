package main

import "fmt"

func main() {

	strs := []string{"ciaio", "ciaione", "ciain", "ciaios"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {

	first_string_len := len(strs[0])
	col_len := 0
	row_len := len(strs)
	z_ext := 0

	if len(strs) > 0 {
		col_len = first_string_len
	}

	for i := 1; i < len(strs); i++ {
		ln := len(strs[i])
		if ln < col_len {
			col_len = ln
		}
	}

	for ; z_ext < col_len; z_ext++ {

		ch := strs[0][z_ext]

		for i := 1; i < row_len; i++ {

			if strs[i][z_ext] != ch {
				return strs[0][:z_ext]
			}
		}
	}

	return strs[0][:z_ext]
}
