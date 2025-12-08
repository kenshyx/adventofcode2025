package day8

import (
	"sort"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j            int
	distanceSquared int64
}

type DSU struct {
	parent     []int
	size       []int
	components int
}

// NewDSU disjointed set
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size, components: n}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return false
	}
	// union by size
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	d.components--
	return true
}

func distance3dSquared(a, b Point) int64 {
	dx := int64(a.x - b.x)
	dy := int64(a.y - b.y)
	dz := int64(a.z - b.z)
	return dx*dx + dy*dy + dz*dz
}

func GetSolution(url string) utils.Solution {
	reader, resp := utils.FetchInput(url)
	var points []Point
	if resp != nil {
		defer resp.Body.Close()
	}
	for {
		line, e := reader.ReadString('\n')
		if e != nil {
			break
		}
		line = strings.TrimSpace(line)
		values := strings.Split(line, ",")
		if len(values) != 3 {
			continue
		}
		var convertedValues [3]int
		for i, v := range values {
			f, _ := strconv.Atoi(v)
			convertedValues[i] = f
		}
		points = append(points, Point{x: convertedValues[0], y: convertedValues[1], z: convertedValues[2]})
	}
	n := len(points)
	var edges []Edge
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			distance := distance3dSquared(points[i], points[j])
			edges = append(edges, Edge{i: i, j: j, distanceSquared: distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distanceSquared < edges[j].distanceSquared
	})

	dsu := NewDSU(n)
	const K = 1000
	for k := 0; k < K && k < len(edges); k++ {
		edge := edges[k]
		_ = dsu.Union(edge.i, edge.j)
	}

	rootSize := make(map[int]int)
	for i := 0; i < n; i++ {
		r := dsu.Find(i)
		rootSize[r] = dsu.size[r]
	}

	var sizes []int
	seen := make(map[int]bool)
	for root, sz := range rootSize {
		if !seen[root] {
			sizes = append(sizes, sz)
			seen[root] = true
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	// must have at least 3 circuits
	if len(sizes) < 3 {
		return utils.Solution{}
	}

	var extra []Edge
	for k := K; k < len(edges) && dsu.components > 1; k++ {
		e := edges[k]
		if dsu.Union(e.i, e.j) {
			extra = append(extra, e)
		}
	}

	lNode := points[extra[len(extra)-1].i]
	rNode := points[extra[len(extra)-1].j]

	return utils.Solution{Part1: sizes[0] * sizes[1] * sizes[2], Part2: lNode.x * rNode.x}
}
