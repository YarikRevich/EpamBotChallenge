package bfs

import (
	"battlecity_test/game"
	"battlecity_test/third-party/graph"

	"gopkg.in/karalabe/cookiejar.v2/collections/queue"
	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

// Breadth-first-search algorithm data structure.
type Bfs struct {
	graph  *graph.Graph
	source game.Point

	visited map[game.Point]bool
	parents map[game.Point]game.Point
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
	d.paths = make(map[game.Point][]game.Point)

	d.pending = queue.New()
	d.pending.Push(src)
	
	d.builder = stack.New()

	return d
}

// Generates the path from the source node to the destination.
func (d *Bfs) Path(dsts ...game.Point) []game.Point {

	for _, dst := range dsts {
		if !d.Reachable(dst) {
			return nil
		}
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
	return nil
}

// Checks whether a given vertex is reachable from the source.
func (d *Bfs) Reachable(dst game.Point) bool {
	if !d.visited[dst] && !d.pending.Empty() {
		if f := d.search(dst); !f{
			return false
		}
	}
	return d.visited[dst]
}

// Continues the bfs search from the last yield point, looking for dst.
func (d *Bfs) search(dst game.Point) bool {
	for !d.pending.Empty() {
		// Fetch the next node, and visit if new
		src := d.pending.Pop().(game.Point)
		d.graph.Do(src, func(peer interface{}) {
			if p := peer.(game.Point); !d.visited[p] {
				d.visited[p] = true
				d.parents[p] = src
				d.pending.Push(p)
			}
		})
		if dst == src {
			return true
		}
	}
	return false
}
