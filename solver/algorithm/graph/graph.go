package graph

import (
	"battlecity_test/game"

	_ "gopkg.in/karalabe/cookiejar.v2/collections/bag"
)

// Data structure for representing a graph.
type Graph struct {
	nodes int
	infos map[int]interface{}
	edges map[game.Point][]game.Point
}

// Creates a new undirected graph.
func New(vertices int) *Graph {
	g := &Graph{
		nodes: vertices,
		infos: make(map[int]interface{}),
		edges: make(map[game.Point][]game.Point),
	}
	return g
}

// Returns the number of vertices in the graph.
func (g *Graph) Vertices() int {
	return g.nodes
}

// Assigns some data to a graph node.
func (g *Graph) Assign(id int, data interface{}) {
	g.infos[id] = data
}

// Retrieves the data associated with a graph node.
func (g *Graph) Retrieve(id int) interface{} {
	return g.infos[id]
}

// Connects two vertices of a graph (may be a loopback).
func (g *Graph) Connect(a, b game.Point) {
	// fmt.Println()
	g.edges[a] = append(g.edges[a], b)
	if a != b {
		g.edges[b] = append(g.edges[b], a)
	}
}

func RemoveConnection(a game.Point, b []game.Point)[]game.Point {
	var i int
	for in, v := range b{
		if a == v{
			i = in
		}
	}
	return append(b[:i], b[i+1:]...)
}

func (g *Graph) Disconnect(a, b game.Point) {

	g.edges[a] = RemoveConnection(b, g.edges[a])
	if a != b {
		g.edges[b] = RemoveConnection(a, g.edges[b])
	}
}

func (g *Graph) Show()map[game.Point][]game.Point{
	return g.edges
}

// // Executes a function for every neighbor of a vertex.
func (g *Graph) Do(vert game.Point, f func(interface{})) {
	for _, v := range g.edges[vert]{
		f(v)
	}
}