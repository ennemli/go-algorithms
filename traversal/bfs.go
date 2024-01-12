
package traversal
import (
  "go-algorithms/tree"
  "go-algorithms/ds"
)
func BFS(node *tree.Node)  []any{
  result:=[]any{}
  visited :=map[*tree.Node]bool{}
  queue:=ds.NewQueue([]*tree.Node{node})
  for queue.Len()>0{
    v:=queue.Dequeue()
    if _,ok:=visited[v];!ok{
      visited[v]=true
      result=append(result,v.Value)
    }
    for _,n:=range v.Neighbors{
      if _,ok:=visited[n];!ok{
        queue.Enqueue(n)
    }
 
    }
  }
  return result
}
