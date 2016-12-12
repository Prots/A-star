package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x           int
	y           int
	obstruction int
	parent      *Point
	g	    int
	h           int
}

// 5. Create an empty "opened" list of Points
// 6. Create an empty "closed" list of Points
var openedList = make([]*Point, 0)
var closedList = make([]*Point, 0)

const (
	defObstruction = 0
	maxNeigbours   = 8
	directWeight   = 10
	diagonalWeight = 14
)

func newPoint(xi, yj, obstruction int, parent *Point) Point {
	return Point{x: xi, y: yj, obstruction: obstruction, parent: parent, g: 0, h: 0}
}

func main() {
	fmt.Println("Start A-STAR")
	// 1. Setup empty field
	size := 10
	field := make([][]Point, size)
	for xi := 0; xi < size; xi++ {
		field[xi] = make([]Point, size)
		for yj := 0; yj < size; yj++ {
			field[xi][yj] = newPoint(xi, yj, defObstruction, nil)
		}
	}
	fmt.Printf("New empty field: %v\n", field)
	// 2. Setup obstacles (points with transparency < 1)
	field[6][3].obstruction = 1
	field[6][4].obstruction = 1
	field[6][5].obstruction = 1
	field[6][6].obstruction = 1
	fmt.Printf("Field with obstacles: %v\n", field)
	// 3. Choose a start point
	startPoint := field[0][0]
	fmt.Printf("Start point: %v\n", startPoint)
	// 4. Choose an end point
	endPoint := field[9][6]
	fmt.Printf("End point: %v\n", endPoint)
	// 7. Add startPoint to the openedList
	openedList = append(openedList, &startPoint)
// ===========================================================================


	// 8. Find all neigbourghs for startPoint excluding obstacles and also add them to the openedList
	neighbours := findNeighbours(&startPoint, field)
	// 9. Setup startPoint for neigbourghs as a parent point
	for _, point := range neighbours {
		point.parent = &startPoint
	}
	fmt.Printf("Neighbours: %v\n", neighbours)
	// 10. Remove start point from the openedList and add it to the closedList
	fmt.Printf("openedList: %v\n", openedList)
	openedList = openedList[0 : len(openedList)-1]
	fmt.Printf("openedList: %v\n", openedList)
	closedList = append(closedList, &startPoint)
	fmt.Printf("closedList: %v\n", closedList)
}

// 1. Setup empty field
// 2. Setup obstacles (points with transparency < 1)
// 3. Choose a start point
// 4. Choose an end point
// 5. Create an empty "opened" list of Points
// 6. Create an empty "closed" list of Points
// 7. Add startPoint to the openedList
// 8. Find all neigbourghs for startPoint excluding obstacles and also add them to the openedList
// 9. Setup startPoint for neigbourghs as a parent point
// 10. Remove start point from the openedList and add it to the closedList

func findNeighbours(point *Point, field [][]Point) []*Point {
	neigbourghs := make([]*Point, 0)
	fmt.Println("point: ", &point)
	if point.x >= 1 && point.y >= 1 {
		for i := (point.x - 1); i <= (point.x + 1); i++ {
			for j := (point.y - 1); j <= (point.y + 1); j++ {
				if field[i][j].obstruction == 0 && !(point.x == i && point.y == j) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	if point.x >= 1 && point.y == 0 {
		for i := (point.x - 1); i <= (point.x + 1); i++ {
			for j := point.y; j <= (point.y + 1); j++ {
				if field[i][j].obstruction == 0 && !(point.x == i && point.y == j) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	if point.x == 0 && point.y >= 1 {
		for i := point.x; i <= (point.x + 1); i++ {
			for j := (point.y - 1); j <= (point.y + 1); j++ {
				if field[i][j].obstruction == 0 && !(point.x == i && point.y == j) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	if point.x == 0 && point.y == 0 {
		for i := point.x; i <= (point.x + 1); i++ {
			for j := point.y; j <= (point.y + 1); j++ {
				if field[i][j].obstruction == 0 && !(point.x == i && point.y == j) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	return neigbourghs
}

func manhattenDistance(start, end *Point) float64 {
	return math.Abs((float64)(end.x-start.x)) + math.Abs((float64)(end.y-start.y))
}

func evklidDistance(start, end *Point) float64 {
	dx := math.Abs((float64)(end.x - start.x))
	dy := math.Abs((float64)(end.y - start.y))
	return math.Hypot(dx, dy)
}

//func findNextPoint(start *Point) {
//	for _, neighbour := range openedList {
//		findWeight(start, neighbour)
//	}
//}

func findWeight(start, finish, current *Point) float64 {
	hWeight := manhattenDistance(current, finish)
	gWeight := getBackRoute(start, current)
	return gWeight + hWeight
}

func getBackRoute(start, end *Point) float64 {
	route := 0
	if start == end {
		return 0
	}
	if end.parent == nil {
		return 0
	}
	if manhattenDistance(end, end.parent) == 1 {
		route = 10
	}
	if manhattenDistance(end, end.parent) == 2 {
		route = 14
	}
	return (float64)(route + getBackRoute(start, end.parent))
}

func (points []*Point) Len() int {
	return len(points)
}
func (points []*Point) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}
func (points []*Point) Less(i, j int) bool {
	return (points[i].g + points[i].h) < (points[j].g + points[j].h)
}

func addToOpenedList(point *Point) {
	openedList = append(openedList, point)
	sort.Sort(openedList)
}