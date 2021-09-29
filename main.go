package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

type PhoneNumber string // TODO: Better memory usage than string ?

type Node struct {
	phoneNumber PhoneNumber
}

type Edge struct {
	in, out *Node
}

type SocialGraph struct {
	nodes []Node
	edges []Edge
}

func (node *Node) String() string {
	return fmt.Sprintf("Node{phoneNumber:%s}", node.phoneNumber)
}

func buildSocialGraph(size, meanEdges int) SocialGraph {
	graph := SocialGraph{
		nodes: make([]Node, 0, size),
		edges: make([]Edge, 0, meanEdges*size),
	}
	buildNodes(size, &graph)
	buildEdges(size, &graph)
	return graph
}

func buildEdges(size int, graph *SocialGraph) {
	for i := 0; i < size; i++ {
		edgesCount := randomEdgesCount()
		for j := 0; j < edgesCount; j++ {
			nodeIndex := rand.Intn(size-1) // Ignore possible duplicates
			graph.edges = append(graph.edges, Edge{in: &graph.nodes[i], out: &graph.nodes[nodeIndex]})
		}
	}
}

func randomEdgesCount() int {
	min := 0
	max := 100
	return rand.Intn(max-min+1) + min
}

func buildNodes(size int, graph *SocialGraph) {
	for i := 0; i < size; i++ {
		graph.nodes = append(graph.nodes, Node{phoneNumber: generatePhoneNumber(i)})
	}
	sort.Slice(graph.nodes, func(i, j int) bool {
		return isLessPhoneNumber(graph.nodes[i].phoneNumber, graph.nodes[j].phoneNumber)
	})
}

func isLessPhoneNumber(p1 PhoneNumber, p2 PhoneNumber) bool {
	return p1 < p2
}

func generatePhoneNumber(i int) PhoneNumber {
	return PhoneNumber(fmt.Sprintf("+33%010d", i))
}

func (graph *SocialGraph) lookup(phoneNumber PhoneNumber) []*Node {
	children := make([]*Node, 0)
	node := graph.findNode(phoneNumber)
	if node == nil {
		return []*Node{}
	}

	for _, edge := range graph.edges {
		if edge.in == node {
			children = append(children, edge.out)
		}
	}
	return children
}

func (graph *SocialGraph) rlookup(phoneNumber PhoneNumber) []*Node {
	children := make([]*Node, 0)
	node := graph.findNode(phoneNumber)
	if node == nil {
		return []*Node{}
	}

	for _, edge := range graph.edges {
		if edge.out == node {
			children = append(children, edge.in)
		}
	}
	return children
}

type Suggestions map[*Node]float64

func (graph *SocialGraph) suggest(phoneNumber PhoneNumber) PairList {
	suggestions := make(Suggestions)
	node := graph.findNode(phoneNumber)
	friends := graph.lookup(node.phoneNumber)
	const depth = 3
	var rec func (currentNode *Node, depth int, scoreRatio float64)
	rec = func (currentNode *Node, depth int, scoreRatio float64) {
		if depth <= 0 {
			return
		}
		children := graph.lookup(currentNode.phoneNumber)
		for _, child := range children {
			isAlreadyFriends := isIn(friends, child)
			var newRatio float64
			if isAlreadyFriends {
				newRatio = scoreRatio * 2
			} else {
				suggestions[child] += scoreRatio
				newRatio = scoreRatio / 2
			}
			rec(child, depth-1, newRatio)
		}
	}
	rec(node, depth, 1.0)

	return firstSuggestions(suggestions, 10)
}

func isIn(friends []*Node, child *Node) bool {
	isIn := false
	for _, friend := range friends {
		if friend == child {
			isIn = true
			break
		}
	}
	return isIn
}

type Suggestion struct {
	node  *Node
	score float64
}

func (graph *SocialGraph) findNode(phoneNumber PhoneNumber) *Node {
	i := sort.Search(len(graph.nodes), func(i int) bool { return !isLessPhoneNumber(graph.nodes[i].phoneNumber, phoneNumber) })
	if i < len(graph.nodes) && graph.nodes[i].phoneNumber == phoneNumber {
		return &graph.nodes[i]
	} else {
		return nil
	}
}

func main() {
	meanEdges := 50
	rand.Seed(time.Now().UnixNano())

	graph := buildSocialGraph(100_000, meanEdges)
	PrintMemUsage()

	fmt.Println(graph.lookup(graph.nodes[0].phoneNumber))
	fmt.Println(graph.rlookup(graph.nodes[0].phoneNumber))
	fmt.Println(graph.suggest(graph.nodes[0].phoneNumber))

	runtime.GC()
	graph = buildSocialGraph(1_000_000, meanEdges)
	PrintMemUsage()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory usage of 100M graph should be %v MiB\n", m.Alloc*100/1024/1024)
}
