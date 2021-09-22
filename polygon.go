package geometry

import "github.com/intdxdt/geom"

type Polygon struct {
	*geom.Polygon
	Id   string
	Meta string
}
