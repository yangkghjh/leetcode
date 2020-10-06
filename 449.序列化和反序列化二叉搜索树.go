package leetcode

import (
	"bytes"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func NewCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	slice := SerializeBinaryTree(root)

	buf := new(bytes.Buffer)

	for _, i := range slice {
		buf.WriteString(strconv.Itoa(i) + ",")
	}

	if buf.Len() > 0 {
		return buf.String()[:buf.Len()-1]
	}

	return ""
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	r := strings.Split(data, ",")
	nums := []int{}
	for _, i := range r {
		if i != "" {
			n, _ := strconv.Atoi(i)
			nums = append(nums, n)
		}
	}

	return NewBinaryTree(nums...)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end
