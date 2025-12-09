package day9

import (
	"sort"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

type Point struct {
	Row int
	Col int
}

type Rectangle struct {
	MinRow, MaxRow int
	MinCol, MaxCol int
	Area           int
}

func GetSolution(url string) utils.Solution {
	reader, resp := utils.FetchInput(url)
	solution := utils.Solution{}

	if resp != nil {
		defer resp.Body.Close()
	}

	var points []Point
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		values := strings.Split(line, ",")
		if len(values) != 2 {
			continue
		}

		col, errCol := strconv.Atoi(values[0])
		row, errRow := strconv.Atoi(values[1])
		if errCol != nil || errRow != nil {
			continue
		}

		points = append(points, Point{
			Row: row,
			Col: col,
		})
	}

	if len(points) < 2 {
		return solution
	}

	n := len(points)

	// Build all rectangles between pairs of points
	rectCount := n * (n - 1) / 2
	rectangles := make([]Rectangle, 0, rectCount)
	maxArea := 0

	for i := 0; i < n; i++ {
		firstCorner := points[i]
		for j := i + 1; j < n; j++ {
			secondCorner := points[j]

			minRow, maxRow := firstCorner.Row, secondCorner.Row
			if minRow > maxRow {
				minRow, maxRow = maxRow, minRow
			}

			minCol, maxCol := firstCorner.Col, secondCorner.Col
			if minCol > maxCol {
				minCol, maxCol = maxCol, minCol
			}

			area := (maxRow - minRow + 1) * (maxCol - minCol + 1)
			if area > maxArea {
				maxArea = area
			}

			rectangles = append(rectangles, Rectangle{
				MinRow: minRow,
				MaxRow: maxRow,
				MinCol: minCol,
				MaxCol: maxCol,
				Area:   area,
			})
		}
	}

	solution.Part1 = maxArea

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Area > rectangles[j].Area
	})

	edgeColsByRow := make(map[int][]int, len(points))
	var edgeRows []int

	addEdgePoint := func(p Point) {
		cols, ok := edgeColsByRow[p.Row]
		if !ok {
			edgeRows = append(edgeRows, p.Row)
		}
		edgeColsByRow[p.Row] = append(cols, p.Col)
	}

	for i, current := range points {
		next := points[(i+1)%len(points)]
		for _, ep := range intermediateEdgePoints(current, next) {
			addEdgePoint(ep)
		}
	}

	for _, p := range points {
		addEdgePoint(p)
	}

	// Sort and deduplicate columns in each row.
	for row, cols := range edgeColsByRow {
		if len(cols) == 0 {
			continue
		}
		sort.Ints(cols)
		write := 1
		for read := 1; read < len(cols); read++ {
			if cols[read] != cols[write-1] {
				cols[write] = cols[read]
				write++
			}
		}
		edgeColsByRow[row] = cols[:write]
	}

	// Sort and deduplicate rows.
	if len(edgeRows) > 0 {
		sort.Ints(edgeRows)
		write := 1
		for read := 1; read < len(edgeRows); read++ {
			if edgeRows[read] != edgeRows[write-1] {
				edgeRows[write] = edgeRows[read]
				write++
			}
		}
		edgeRows = edgeRows[:write]
	}

	// Find the largest valid rectangle

	for _, rect := range rectangles {
		if rect.MaxRow-rect.MinRow <= 1 || rect.MaxCol-rect.MinCol <= 1 {
			solution.Part2 = rect.Area
			break
		}

		if rectangleHasEdgePointInside(rect, edgeRows, edgeColsByRow) {
			continue
		}

		solution.Part2 = rect.Area
		break
	}

	return solution
}

func rectangleHasEdgePointInside(
	rect Rectangle,
	edgeRows []int,
	edgeColsByRow map[int][]int,
) bool {

	startIdx := sort.SearchInts(edgeRows, rect.MinRow+1)
	for i := startIdx; i < len(edgeRows); i++ {
		row := edgeRows[i]
		if row >= rect.MaxRow {
			break
		}

		cols := edgeColsByRow[row]

		firstIdx := sort.SearchInts(cols, rect.MinCol+1)
		if firstIdx < len(cols) && cols[firstIdx] < rect.MaxCol {

			return true
		}
	}
	return false
}

func intermediateEdgePoints(a, b Point) []Point {

	if a.Row == b.Row {
		dist := abs(a.Col - b.Col)
		if dist <= 1 {
			return nil
		}

		minCol, maxCol := a.Col, b.Col
		if minCol > maxCol {
			minCol, maxCol = maxCol, minCol
		}

		points := make([]Point, 0, dist-1)
		for col := minCol + 1; col < maxCol; col++ {
			points = append(points, Point{
				Row: a.Row,
				Col: col,
			})
		}
		return points
	}

	dist := abs(a.Row - b.Row)
	if dist <= 1 {
		return nil
	}

	minRow, maxRow := a.Row, b.Row
	if minRow > maxRow {
		minRow, maxRow = maxRow, minRow
	}

	points := make([]Point, 0, dist-1)
	for row := minRow + 1; row < maxRow; row++ {
		points = append(points, Point{
			Row: row,
			Col: a.Col,
		})
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
