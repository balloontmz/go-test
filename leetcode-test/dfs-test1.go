package main

// 一个简单的深度优先遍历练习
// 暂时没实现树的赋值，所以遍历暂时没意义
// 未完成

// Node 单个节点的结构体
type Node struct {
	val  int
	Next []*Node
}

func main() {

}

func dfs(node *Node, target int, mP *map[int]int) bool {
	// 查看当前递归的节点是否为指定目标
	if node.val == target {
		return true
	}
	// 遍历当前节点的所有子节点
	for _, val := range node.Next {
		// m := *mP // 取该地址的值
		if _, ok := (*mP)[val.val]; !ok {
			(*mP)[val.val] = 1
			if dfs(val, target, mP) == true {
				return true
			}
		}
	}
	return false
}
