package graph

import (
	"slices"
)

type Graph[T comparable] struct {
	nodes     []T
	adjacency map[T][]T
}

func New[T comparable](nodes ...T) *Graph[T] {
	return &Graph[T]{
		nodes:     nodes,
		adjacency: make(map[T][]T),
	}
}

func (g *Graph[T]) AddNode(node T) bool {
	if slices.Contains(g.nodes, node) {
		return false
	}
	g.nodes = append(g.nodes, node)
	return true
}

func (g *Graph[T]) AddEdge(u, v T) {
	g.adjacency[u] = append(g.adjacency[u], v)
}

func (g *Graph[T]) isCircular(node T, visited, recursionStack map[T]bool) bool {
	visited[node] = true
	recursionStack[node] = true

	for _, neighbor := range g.adjacency[node] {
		if !visited[neighbor] {
			if g.isCircular(neighbor, visited, recursionStack) {
				return true
			}
		} else if recursionStack[neighbor] {
			return true
		}
	}

	recursionStack[node] = false
	return false
}

func (g *Graph[T]) IsCircular() bool {
	visited := make(map[T]bool)
	recursionStack := make(map[T]bool)

	for _, node := range g.nodes {
		if !visited[node] {
			if g.isCircular(node, visited, recursionStack) {
				return true
			}
		}
	}
	return false
}
