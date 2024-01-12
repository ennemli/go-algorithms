package traversal

import (
	"go-algorithms/ds"
	"go-algorithms/tree"
)

func DFS(node *tree.Node) []any {
	result := []any{}
	visited := map[*tree.Node]bool{}
	stack := ds.NewStack([]*tree.Node{node})
	for stack.Len() > 0 {
		v := stack.Pop()
		if _, ok := visited[v]; !ok {
			visited[v] = true
			result = append(result, v.Value)
		}
		for _, n := range v.Neighbors {
			if _, ok := visited[n]; !ok {
				stack.Push(n)
			}
		}
	}
	return result
}
