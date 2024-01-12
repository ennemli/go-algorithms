package main

import (
	"fmt"
  "go-algorithms/tree"
  "go-algorithms/traversal"
)

func main(){
  a:=&tree.Node{Value:"A"}
  b:=&tree.Node{Value:"B"}
  c:=&tree.Node{Value:"C"}
  d:=&tree.Node{Value:"D"}
  e:=&tree.Node{Value:"E"}
  a.Neighbors=[]*tree.Node{d,c}
  b.Neighbors=[]*tree.Node{d,c,e}
  c.Neighbors=[]*tree.Node{a,b}
  d.Neighbors=[]*tree.Node{a,e}
  e.Neighbors=[]*tree.Node{b,d}
  dfs_result:=traversal.DFS(a)
  bfs_result:=traversal.BFS(a)
  fmt.Println(dfs_result)
  fmt.Println(bfs_result)
}
