package geometry

import "github.com/intdxdt/geom"

type Point struct {
	geom.Point
	Id   string
	Meta string
}


//CreatePoint construct new polyline
func CreatePoint(id string, coordinates []float64, meta string) Point {
	return Point{geom.CreatePoint(coordinates), id, meta}
}

