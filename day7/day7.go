package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"io"
)

type Node struct {
	name     string
	weight   int
	children []*Node
	parent   *Node
}

func (n *Node) TowerWeight() int {
	weight := n.weight
	for _, child := range n.children {
		weight += child.TowerWeight()
	}
	return weight
}

func (n *Node) Print() string {
	children := make([]string, 0)
	for _, c := range n.children {
		children = append(children, c.name)
	}

	childrenAsStr := strings.Join(children, ",")
	if len(childrenAsStr) == 0 {
		childrenAsStr = "[]"
	}

	return fmt.Sprintf("%s[%d, %s]", n.name, n.weight, childrenAsStr)
}

func (n *Node) FindUnbalancedNode() *Node {

	fmt.Println("Finding unbalanced for node", n.Print())

	for _, child := range n.children {
		unbalanced := child.FindUnbalancedNode()

		if unbalanced != nil {
			return unbalanced
		}
	}

	odd := FindOddOneOut(n.children)
	return odd
}

func FindOddOneOut(nodes []*Node) *Node {
	odd := func(a *Node, b *Node, c *Node) *Node {
		if a.weight == b.weight && b.weight != c.weight {
			return c
		}
		if a.weight == c.weight && c.weight != b.weight {
			return b
		}
		if b.weight == c.weight && c.weight != a.weight {
			return a
		}
		return nil
	}

	for i := 0; i < len(nodes)-2; i++ {
		oddOne := odd(nodes[i], nodes[i+1], nodes[i+2])
		if oddOne != nil {
			return oddOne
		}
	}
	return nil
}

func (n *Node) HasBalancedChildren() bool {
	weight := -1
	for _, child := range n.children {
		w := child.TowerWeight()
		if weight == -1 {
			weight = w
		} else {
			if w != weight {
				return false
			}
		}
	}

	return true
}

func NewNode(name string, weight int) *Node {
	return &Node{
		name:     name,
		weight:   weight,
		children: make([]*Node, 0),
	}
}

func BuildGraph(root string, graph map[string][]string, weights map[string]int) *Node {
	parent := NewNode(root, weights[root])

	children, _ := graph[root]
	for _, child := range children {
		childNode := BuildGraph(child, graph, weights)
		childNode.parent = parent
		parent.children = append(parent.children, childNode)
	}

	return parent
}

func FindUnbalancedNode(node *Node) (*Node, *Node) {
	isBalanced := node.HasBalancedChildren()
	if !isBalanced {

		for _, child := range node.children {
			unbalanced, _ := FindUnbalancedNode(child)

			if unbalanced == nil {
				//	this is the first unbalanced
				return child, node
			}
		}

	}

	return nil, node
}

func GetNodeDepth(key string, graph map[string][]string) int {
	depth := 1
	children, ok := graph[key]
	if !ok {
		panic(fmt.Sprintf("%s %s", key, graph))
	}
	max := 0
	for _, child := range children {
		d := GetNodeDepth(child, graph)
		if d > max {
			max = d
		}
	}

	return depth + max
}

func FindRoot(graph map[string][]string) string {
	var maxKey string
	maxDepth := 0
	for key, _ := range graph {
		depth := GetNodeDepth(key, graph)
		if depth > maxDepth {
			maxDepth = depth
			maxKey = key
		}
	}

	return maxKey
}

func ParseRow(row string) (string, int, []string) {
	tokens := strings.Fields(row)
	name := tokens[0]
	weight, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(tokens[1], "("), ")"))
	if err != nil {
		panic(err)
	}

	children := make([]string, 0)
	if len(tokens) > 3 {
		rawChildren := tokens[3:]
		for _, child := range rawChildren {
			children = append(children, strings.TrimSuffix(child, ","))
		}
	}

	return name, weight, children
}

func ParseInput(reader io.Reader) *Node {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	children := make(map[string][]string)
	weights := make(map[string]int)

	for scanner.Scan() {
		row := scanner.Text()
		name, weight, childs := ParseRow(row)
		children[name] = childs
		weights[name] = weight
	}

	root := FindRoot(children)
	return BuildGraph(root, children, weights)
}

func main() {
	graph := ParseInput(os.Stdin)

	fmt.Println("Root node:", graph.name)
}
