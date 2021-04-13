package main

import "math"

//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list []int
var prev *TreeNode
var tempMin  = math.MaxInt64

func minDiffInBST(root *TreeNode) int {

	var minVal = math.MaxInt64
	dfs(root)
	if len(list) < 2 {
		return 0
	}

	for i := 0; i < len(list)-1; i++ {
		minVal = min(minVal, list[i + 1] - list[i])
	}
	list = []int{}
	return minVal
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	list = append(list, root.Val)
	dfs(root.Right)
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	} else {
		return v1
	}
}


func minDiffInBST2(root *TreeNode) int {
	dfs2(root)
	result := tempMin
	prev = nil
	tempMin  = math.MaxInt64
	return result
}

func dfs2(root *TreeNode)  {
	if root == nil {
		return
	}

	dfs2(root.Left)

	if prev != nil {
		tempMin = min(tempMin, root.Val - prev.Val)
	}

	prev = root
	dfs2(root.Right)
}