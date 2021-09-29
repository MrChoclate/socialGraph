package main

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"
)

func buildDummyGraph() SocialGraph {
	graph := SocialGraph{[]Node{Node{phoneNumber: "p"}, Node{phoneNumber: "x"}, Node{phoneNumber: "y"}, Node{phoneNumber: "z"}}} // must be sorted
	p, x, y, z := &graph.nodes[0], &graph.nodes[1], &graph.nodes[2], &graph.nodes[3]

	addEdge(x, y)
	addEdge(x, z)
	addEdge(y, x)
	addEdge(y, z)
	addEdge(p, x)
	return graph
}


func TestLookup(t *testing.T) {
	graph := buildDummyGraph()
	expected := []string{"z", "y"}
	lookup := graph.lookup("x")
	var actual []string
	for _, node := range lookup {
		actual = append(actual, string(node.phoneNumber))
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v %v", expected, actual)
	}
}

func TestRlookup(t *testing.T) {
	graph := buildDummyGraph()
	expected := []string{"y", "x"}
	lookup := graph.rlookup("z")
	var actual []string
	for _, node := range lookup {
		actual = append(actual, string(node.phoneNumber))
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v %v", expected, actual)
	}
}

func TestSuggest(t *testing.T) {
	graph := buildDummyGraph()
	expected := []string{"z", "y"}
	lookup := graph.suggest("p")
	var actual []string
	for _, suggestion := range lookup {
		actual = append(actual, string(suggestion.node.phoneNumber))
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v %v", expected, actual)
	}
}


var result interface{}

func benchmarkLookup(size int, b *testing.B) {
	rand.Seed(0)
	graph := buildSocialGraph(size, 50)
	runtime.GC()
	b.ResetTimer()
	var lookup []*Node
	for n := 0; n < b.N; n++ {
		lookup = graph.lookup(graph.nodes[rand.Intn(len(graph.nodes))].phoneNumber)
	}
	result = lookup
}

func benchmarkRlookup(size int, b *testing.B) {
	rand.Seed(0)
	graph := buildSocialGraph(size, 50)
	runtime.GC()
	b.ResetTimer()
	var rlookup []*Node
	for n := 0; n < b.N; n++ {
		rlookup = graph.rlookup(graph.nodes[rand.Intn(len(graph.nodes))].phoneNumber)
	}
	result = rlookup
}

func benchmarkFindNode(size int, b *testing.B) {
	rand.Seed(0)
	graph := buildSocialGraph(size, 0)
	runtime.GC()
	b.ResetTimer()
	var findNode *Node
	for n := 0; n < b.N; n++ {
		findNode = graph.findNode(graph.nodes[rand.Intn(len(graph.nodes))].phoneNumber)
	}

	result = findNode
}

func benchmarkSuggestions(size int, b *testing.B) {
	rand.Seed(0)
	graph := buildSocialGraph(size, 50)
	runtime.GC()
	b.ResetTimer()
	var suggest PairList
	for n := 0; n < b.N; n++ {
		 suggest = graph.suggest(graph.nodes[rand.Intn(len(graph.nodes))].phoneNumber)
	}
	result = suggest
}

func BenchmarkLookupSmall(b *testing.B) {
	benchmarkLookup(100_000, b)
}

func BenchmarkLookupBig(b *testing.B) {
	benchmarkLookup(1_000_000, b)
}

func BenchmarkLookupBig2(b *testing.B) {
	benchmarkLookup(2_000_000, b)
}

func BenchmarkRlookupSmall(b *testing.B) {
	benchmarkRlookup(100_000, b)
}

func BenchmarkRlookupBig(b *testing.B) {
	benchmarkRlookup(1_000_000, b)
}

func BenchmarkRlookupBig2(b *testing.B) {
	benchmarkRlookup(2_000_000, b)
}

func BenchmarkSuggestionsSmall(b *testing.B) {
	benchmarkSuggestions(100_000, b)
}

func BenchmarkSuggestionsBig(b *testing.B) {
	benchmarkSuggestions(1_000_000, b)
}

func BenchmarkSuggestionsBig2(b *testing.B) {
	benchmarkSuggestions(2_000_000, b)
}

func BenchmarkFindNodeSmall(b *testing.B) {
	benchmarkFindNode(100_000, b)
}

func BenchmarkFindNodeBig(b *testing.B) {
	benchmarkFindNode(1_000_000, b)
}

func BenchmarkFindNodeBig2(b *testing.B) {
	benchmarkFindNode(2_000_000, b)
}


