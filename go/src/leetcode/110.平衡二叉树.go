/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (51.50%)
 * Likes:    308
 * Dislikes: 0
 * Total Accepted:    72.6K
 * Total Submissions: 140.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 *
 * 本题中，一棵高度平衡二叉树定义为：
 *
 *
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
 *
 *
 * 示例 1:
 *
 * 给定二叉树 [3,9,20,null,null,15,7]
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回 true 。
 *
 * 示例 2:
 *
 * 给定二叉树 [1,2,2,3,3,null,null,4,4]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * 返回 false 。
 *
 *
 *
 */
package main

import (
	"math"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

	var depth int
	return isBalancedTree(root, &depth)
}

func isBalancedTree(node *TreeNode, depth *int) bool {
	if node == nil {
		*depth = 0
		return true
	}
	var leftDepth int
	var rightDepth int
	l := isBalancedTree(node.Left, &leftDepth)
	r := isBalancedTree(node.Right, &rightDepth)
	if l && r {
		diff := leftDepth - rightDepth
		if diff <= 1 && diff >= -1 {
			*depth = int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
			return true
		}
	}
	return false
}

// @lc code=end
