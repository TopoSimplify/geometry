package geometry

import "github.com/intdxdt/geom"

type Polygon struct {
	*geom.Polygon
	Id   string
	Meta string
}

//CreatePolygon constructs new Polygon
func CreatePolygon(id string, coordinates []geom.Coords, meta string) Polygon {
	return Polygon{geom.NewPolygon(coordinates...), id, meta}
}

