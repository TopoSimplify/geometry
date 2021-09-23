package geometry

import (
	"github.com/TopoSimplify/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

//Polyline Type
type Polyline struct {
	G *geom.LineString
	Id       string
	Meta     string
}

func (g *Polyline) Geometry() geom.Geometry {
	return g.G
}


//CreatePolyline construct new polyline
func CreatePolyline(id string, coordinates geom.Coords, meta string) Polyline {
	return Polyline{geom.NewLineString(coordinates), id, meta}
}

//SegmentBounds segment bounds
func (ln *Polyline) SegmentBounds() []mono.MBR {
	var I, J int
	var n = ln.Len() - 1
	var a, b *geom.Point
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		a, b = ln.G.Coordinates.Pt(i), ln.G.Coordinates.Pt(i+1)
		I, J = ln.G.Coordinates.Idxs[i], ln.G.Coordinates.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(a[geom.X], a[geom.Y], b[geom.X], b[geom.Y]),
			I:   I, J: J,
		})
	}
	return items
}

//Range of entire polyline
func (ln *Polyline) Range() rng.Rng {
	return rng.Range(ln.G.Coordinates.FirstIndex(), ln.G.Coordinates.LastIndex())
}

//Segment given range
func (ln *Polyline) Segment(i, j int) *geom.Segment {
	return geom.NewSegment(ln.G.Coordinates, i, j)
}

//SubPolyline - generates sub polyline from generator indices
func (ln *Polyline) SubPolyline(rng rng.Rng) Polyline {
	return CreatePolyline(ln.Id, ln.SubCoordinates(rng), ln.Meta)
}

//SubCoordinates - generates sub polyline from generator indices
func (ln *Polyline) SubCoordinates(rng rng.Rng) geom.Coords {
	var coords = ln.G.Coordinates
	coords.Idxs = make([]int, 0, rng.J-rng.I+1)
	for i := rng.I; i <= rng.J; i++ {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}

//Len - Length of coordinates in polyline
func (ln *Polyline) Len() int {
	return ln.G.Coordinates.Len()
}
