package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
  "math"
  "sort"
)

type Point struct {
  x, y float64
}

type PointDist struct {
  point Point
  dist  float64
}

func point_from_str(str_point string) Point {
  arr := strings.Split(strings.Trim(str_point, "()"), ",")
  x, _ := strconv.ParseFloat(arr[0], 64)
  y, _ := strconv.ParseFloat(arr[1], 64)
  return Point{x,y}
}

func dist(p1 Point, p2 Point) float64 {
  return math.Pow(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2), 0.5)
}

type ByDist []PointDist

func (a ByDist) Len() int           { return len(a) }
func (a ByDist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDist) Less(i, j int) bool { return a[i].dist < a[j].dist }

func main() {
  // read in the points from stdin formatted like so:
  // (1,2) (3,4) (5,6) (6,7)
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  str_points := strings.Split(strings.Trim(input, "\n"), " ")
  var points []Point
  for _, str_point := range str_points {
    points = append(points, point_from_str(str_point))
  }
  fmt.Println(points)

  // find the k closest for every point in the list
  var result [][3]Point
  for i, p1 := range points {
    closest3 := make([]PointDist, 0)
    for j, p2 := range points {
      if (i == j) {
        continue
      }
      closestN := append(closest3, PointDist{p2,dist(p1,p2)})
      if (len(closestN) > 3) {
        sort.Sort(ByDist(closestN))
        closest3 = closestN[:3]
      } else {
        closest3 = closestN
      }
    }

    var closest3_points [3]Point
    for k, l := range closest3 { closest3_points[k] = l.point }
    result = append(result, closest3_points)
  }

  // print the results to stdout formatted like so:
  // [[(3,4) (5,6) (6,7)], ...]
  fmt.Println(result)
}
