package main

import (
	"fmt"
)

type Point struct{
	x int
	y int
	obstruction int
	parent *Point
}

// 5. Create an empty "opened" list of Points
// 6. Create an empty "closed" list of Points
var openedList = make([]*Point, 0)
var closedList = make([]*Point, 0)

const(
	defObstruction = 0
	maxNeigbours = 8
	directWeight = 10
	diagonalWeight = 14
)

func newPoint(xi, yj, obstruction int) Point{
	return Point{x: xi, y: yj, obstruction: obstruction}
}

func main() {
	fmt.Println("Start A-STAR")
	// 1. Setup empty field
	size := 10
	field := make([][]Point, size)
	for xi:=0;xi<size;xi++{
		field[xi] = make([]Point, size)
		for yj:=0;yj<size;yj++{
			field[xi][yj] = newPoint(xi, yj, defObstruction)
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
	// 8. Find all neigbourghs for startPoint excluding obstacles and also add them to the openedList
	neighbours := findNeighbours(&startPoint, field)
	// 9. Setup startPoint for neigbourghs as a parent point
	for _, point := range neighbours {
		point.parent = &startPoint
	}
	fmt.Printf("Neighbours: %v\n", neighbours)
	// 10. Remove start point from the openedList and add it to the closedList
	fmt.Printf("openedList: %v\n", openedList)
	openedList = openedList[0:len(openedList)-1]
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
		for i:=(point.x - 1); i <= (point.x + 1); i++{
			for j:=(point.y - 1); j <= (point.y + 1); j++ {
				if (field[i][j].obstruction == 0  && !(point.x == i && point.y == j)) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	if point.x >=1 && point.y == 0 {
		for i:=(point.x - 1); i <= (point.x + 1); i++{
			for j:= point.y; j <= (point.y + 1); j++ {
				if (field[i][j].obstruction == 0  && !(point.x == i && point.y == j)) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	if point.x == 0 && point.y >= 1 {
		for i:= point.x; i <= (point.x + 1); i++{
			for j:=(point.y - 1); j <= (point.y + 1); j++ {
				if (field[i][j].obstruction == 0  && !(point.x == i && point.y == j)) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	if point.x == 0 && point.y == 0 {
		for i:= point.x; i <= (point.x + 1); i++{
			for j:=point.y; j <= (point.y + 1); j++ {
				if (field[i][j].obstruction == 0  && !(point.x == i && point.y == j)) {
					neigbourghs = append(neigbourghs, &field[i][j])
				}
			}
		}
	}
	return neigbourghs
}

func manhattenDistance(start, end *Point) int {
	distance := 0

	return distance
}