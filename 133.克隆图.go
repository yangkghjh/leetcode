package leetcode

type CloneGraphNode struct {
	Val       int
	Neighbors []*CloneGraphNode
}

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *CloneGraphNode) *CloneGraphNode {
	if node == nil {
		return nil
	}

	hash := map[int]*CloneGraphNode{}
	res := &CloneGraphNode{
		Neighbors: []*CloneGraphNode{},
	}

	cloneGraphDFS(node, res, hash)

	return res.Neighbors[0]
}

func cloneGraphDFS(node, target *CloneGraphNode, hash map[int]*CloneGraphNode) {
	n, ok := hash[node.Val]

	if !ok {
		n = &CloneGraphNode{
			Val:       node.Val,
			Neighbors: []*CloneGraphNode{},
		}
		hash[node.Val] = n
	}

	target.Neighbors = append(target.Neighbors, n)

	if !ok {
		for _, son := range node.Neighbors {
			cloneGraphDFS(son, n, hash)
		}
	}
}

// @lc code=end
