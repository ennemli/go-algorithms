package tree
type nodeData interface{}

type Node struct{
  Value nodeData
  Neighbors []*Node
}

