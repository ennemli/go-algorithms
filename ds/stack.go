package ds
import (
  "sync"
)

type Stack[T any] struct{
  mu sync.Mutex    
  items []T
}
func NewStack[T any](initItems []T) *Stack[T]{
  return &Stack[T]{items:initItems}
}
func (s *Stack[T]) Push(item T){
  s.mu.Lock()
  defer s.mu.Unlock()
  s.items=append(s.items,item)
}

func (s *Stack[T]) Pop() T{
  s.mu.Lock()
  defer s.mu.Unlock()
  l:=len(s.items)
  if l<=0{
    panic("Stack already empty")
  }
  item:=s.items[l-1]
  s.items=s.items[:l-1]
  return item 
}

func (s *Stack[T]) Len() int{
  return len(s.items)
}
