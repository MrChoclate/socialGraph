package main

import "testing"

func benchmarkLookup(size int, b *testing.B) {
	b.StopTimer()
	graph := buildSocialGraph(size, 50)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		graph.lookup(graph.nodes[0].phoneNumber)
	}
}

func benchmarkRlookup(size int, b *testing.B) {
	b.StopTimer()
	graph := buildSocialGraph(size, 50)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		graph.rlookup(graph.nodes[0].phoneNumber)
	}
}

func benchmarkSuggestions(size int, b *testing.B) {
	b.StopTimer()
	graph := buildSocialGraph(size, 50)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		graph.suggest(graph.nodes[0].phoneNumber)
	}
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
