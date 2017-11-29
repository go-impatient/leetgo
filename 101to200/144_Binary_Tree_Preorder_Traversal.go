package main

import (
	. "leetgo/structure/btree"
)

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr1 := preorderTraversal(root.Left)
	arr2 := preorderTraversal(root.Right)

	res := []int{root.Val}
	return append(append(res, arr1...), arr2...)
}


func main() {

}