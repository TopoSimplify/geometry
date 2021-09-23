package geometry

import "github.com/intdxdt/geom"

type Point struct {
	G    geom.Point
	Id   string
	Meta string
}

func (g *Point) Geometry() geom.Geometry {
	return g.G
}

//CreatePoint construct new polyline
func CreatePoint(id string, coordinates []float64, meta string) Point {
	return Point{geom.CreatePoint(coordinates), id, meta}
}
