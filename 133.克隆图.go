package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
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

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	hash := map[int]*Node{}
	res := &Node{
		Neighbors: []*Node{},
	}

	cloneGraphDFS(node, res, hash)

	return res.Neighbors[0]
}

func cloneGraphDFS(node, target *Node, hash map[int]*Node) {
	n, ok := hash[node.Val]

	if !ok {
		n = &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
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
