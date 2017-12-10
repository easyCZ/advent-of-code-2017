package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRow(t *testing.T) {
	scenarios := []struct {
		in       string
		name     string
		weight   int
		children []string
	}{
		{"pbga (66)", "pbga", 66, []string{}},
		{"xhth (57)", "xhth", 57, []string{}},
		{"ebii (61)", "ebii", 61, []string{}},
		{"havc (66)", "havc", 66, []string{}},
		{"ktlj (57)", "ktlj", 57, []string{}},
		{"fwft (72) -> ktlj, cntj, xhth", "fwft", 72, []string{"ktlj", "cntj", "xhth"}},
		{"qoyq (66)", "qoyq", 66, []string{}},
		{"padx (45) -> pbga, havc, qoyq", "padx", 45, []string{"pbga", "havc", "qoyq"}},
		{"tknk (41) -> ugml, padx, fwft", "tknk", 41, []string{"ugml", "padx", "fwft"}},
		{"jptl (61)", "jptl", 61, []string{}},
		{"ugml (68) -> gyxo, ebii, jptl", "ugml", 68, []string{"gyxo", "ebii", "jptl"}},
		{"gyxo (61)", "gyxo", 61, []string{}},
		{"cntj (57)", "cntj", 57, []string{}},
	}

	for _, s := range scenarios {
		name, weight, children := ParseRow(s.in)
		assert.Equal(t, s.name, name)
		assert.Equal(t, s.weight, weight)
		assert.Equal(t, s.children, children)

	}

}

func TestFindRoot(t *testing.T) {
	children, _ := GetChildrenAndWeights()
	root := FindRoot(children)

	assert.Equal(t, "tknk", root)
}

func TestBuildGraph(t *testing.T) {
	children, weights := GetChildrenAndWeights()

	root := FindRoot(children)
	graph := BuildGraph(root, children, weights)

	assert.Equal(t, "tknk", graph.name)
	assert.Len(t, graph.children, 3)

	firstChildren := graph.children
	ugml := firstChildren[0]
	padx := firstChildren[1]
	fwft := firstChildren[2]

	assert.Equal(t, ugml.name, "ugml")
	assert.Equal(t, padx.name, "padx")
	assert.Equal(t, fwft.name, "fwft")

	assert.Len(t, ugml.children, 3)
	gyxo := ugml.children[0]
	ebii := ugml.children[1]
	jptl := ugml.children[2]

	assert.Equal(t, gyxo.name, "gyxo")
	assert.Equal(t, ebii.name, "ebii")
	assert.Equal(t, jptl.name, "jptl")

	assert.Len(t, padx.children, 3)
	pbga := padx.children[0]
	havc := padx.children[1]
	qoyq := padx.children[2]

	assert.Equal(t, pbga.name, "pbga")
	assert.Equal(t, havc.name, "havc")
	assert.Equal(t, qoyq.name, "qoyq")

	assert.Len(t, fwft.children, 3)
	ktlj := fwft.children[0]
	cntj := fwft.children[1]
	xhth := fwft.children[2]

	assert.Equal(t, ktlj.name, "ktlj")
	assert.Equal(t, cntj.name, "cntj")
	assert.Equal(t, xhth.name, "xhth")
}

func TestFindUnbalancedNode(t *testing.T) {
	graph := GetGraph()
	unbalanced := graph.FindUnbalancedNode()
	assert.NotNil(t, unbalanced)
	assert.Equal(t, unbalanced.name, "ugml")
}

//func TestFindCorruptNode(t *testing.T) {
//	graph := GetGraph()
//	unbalanced, _ := FindUnbalancedNode(graph)

//
//	assert.Equal(t, "ugml", unbalanced.name)
//}

func GetChildrenAndWeights() (map[string][]string, map[string]int) {
	return map[string][]string{
		"pbga": {},
		"xhth": {},
		"ebii": {},
		"havc": {},
		"ktlj": {},
		"fwft": {"ktlj", "cntj", "xhth"},
		"qoyq": {},
		"padx": {"pbga", "havc", "qoyq"},
		"tknk": {"ugml", "padx", "fwft"},
		"jptl": {},
		"ugml": {"gyxo", "ebii", "jptl"},
		"gyxo": {},
		"cntj": {},
	}, map[string]int{
		"pbga": 66,
		"xhth": 57,
		"ebii": 61,
		"havc": 66,
		"ktlj": 57,
		"fwft": 72,
		"qoyq": 66,
		"padx": 45,
		"tknk": 41,
		"jptl": 61,
		"ugml": 68,
		"gyxo": 61,
		"cntj": 57,
	}

}

func GetGraph() *Node {
	children, weights := GetChildrenAndWeights()
	root := FindRoot(children)
	graph := BuildGraph(root, children, weights)
	return graph
}
