package algorithm

import (

	"battlecity_test/game"
	"battlecity_test/solver/algorithm/graph"

	"gopkg.in/karalabe/cookiejar.v2/collections/queue"
	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

// Breadth-first-search algorithm data structure.
type Bfs struct {
	graph  *graph.Graph
	source game.Point

	visited map[game.Point]bool
	parents map[game.Point]game.Point
	order   []game.Point
	paths   map[game.Point][]game.Point

	pending *queue.Queue
	builder *stack.Stack
}

// Creates a new random-order bfs structure.
func New(g *graph.Graph, src game.Point) *Bfs {
	d := new(Bfs)

	d.graph = g
	d.source = src

	d.visited = make(map[game.Point]bool, g.Vertices())
	d.visited[src] = true
	d.parents = make(map[game.Point]game.Point, g.Vertices())
	d.order = make([]game.Point, 1, g.Vertices())
	d.order[0] = src
	d.paths = make(map[game.Point][]game.Point)

	d.pending = queue.New()
	d.pending.Push(src)
	d.builder = stack.New()

	return d
}

// Generates the path from the source node to the destination.
func (d *Bfs) Path(dst game.Point) []game.Point {
	// Return nil if not reachable
	if !d.Reachable(dst) {
		return nil
	}
	// If reachable, but path not yet generated, create and cache
	if cached, ok := d.paths[dst]; !ok {
		for cur := dst; cur != d.source; {
			d.builder.Push(cur)
			cur = d.parents[cur]
		}
		d.builder.Push(d.source)

		path := make([]game.Point, d.builder.Size())
		for i := 0; i < len(path); i++ {
			path[i] = d.builder.Pop().(game.Point)
		}
		d.paths[dst] = path
		return path
	} else {
		return cached
	}
}

// Checks whether a given vertex is reachable from the source.
func (d *Bfs) Reachable(dst game.Point) bool {
	if !d.visited[dst] && !d.pending.Empty() {
		d.search(dst)
	}
	return d.visited[dst]
}

// Generates the full order in which nodes were traversed.
func (d *Bfs) Order() []game.Point {
	// Force bfs termination
	if !d.pending.Empty() {
		d.search(game.Point{X: 32, Y: 32})
	}
	return d.order
}

// Continues the bfs search from the last yield point, looking for dst.
func (d *Bfs) search(dst game.Point) {
	for !d.pending.Empty() {
		// Fetch the next node, and visit if new
		src := d.pending.Pop().(game.Point)
		d.graph.Do(src, func(peer interface{}) {
			if p := peer.(game.Point); !d.visited[p] {
				d.visited[p] = true
				d.order = append(d.order, p)
				d.parents[p] = src
				d.pending.Push(p)
			}
		})
		// If we found the destination, yield
		if dst == src {
			return
		}
	}
}