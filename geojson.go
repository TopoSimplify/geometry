package geometry

import (
	"encoding/json"
	"fmt"
	"github.com/intdxdt/fileutil"
	"github.com/intdxdt/geom"
	geojson "github.com/paulmach/go.geojson"
	"strings"
)

type GeoJSONGeometries struct {
	Points      []Point
	LineStrings []Polyline
	Polygons    []Polygon
}

type JSONPoint struct {
	Id          string
	Coordinates []float64
	Meta        string
}

type JSONLineString struct {
	Id          string
	Coordinates [][]float64
	Meta        string
}

type JSONPolygon struct {
	Id          string
	Coordinates [][][]float64
	Meta        string
}

func ReadInputPolylines(inputJsonFile string) []Polyline {
	return parseInputLinearFeatures(readJsonFile(inputJsonFile))
}

func ReadInputConstraints(inputJsonFile string) GeoJSONGeometries {
	return parseConstraintFeatures(readJsonFile(inputJsonFile))
}

func parseInputLinearFeatures(inputs []string) []Polyline {
	var plns = make([]Polyline, 0, len(inputs))
	for idx, fjson := range inputs {
		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)
		var objs = getLineStringObjects(idx, feat)
		for _, o := range objs {
			var pln = createPolyline(o)
			plns = append(plns, pln)
		}
	}
	return plns
}

func parseConstraintFeatures(inputs []string) GeoJSONGeometries {
	var pts = make([]Point, 0, len(inputs))
	var plns = make([]Polyline, 0, len(inputs))
	var polys = make([]Polygon, 0, len(inputs))

	for idx, fjson := range inputs {
		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)

		var ptObjs = getPointObjects(idx, feat)
		for _, o := range ptObjs {
			pts = append(pts, createPoint(o))
		}

		var lnObjs = getLineStringObjects(idx, feat)
		for _, o := range lnObjs {
			plns = append(plns, createPolyline(o))
		}

		var plyObjs = getPolygonObjects(idx, feat)
		for _, o := range plyObjs {
			polys = append(polys, createPolygon(o))
		}
	}

	return GeoJSONGeometries{
		Points:      pts,
		LineStrings: plns,
		Polygons:    polys,
	}
}

func createPoint(jsonLine JSONPoint) Point {
	return CreatePoint(jsonLine.Id, jsonLine.Coordinates, jsonLine.Meta)
}

func createPolyline(jsonLine JSONLineString) Polyline {
	var coords = geom.AsCoordinates(jsonLine.Coordinates)
	return CreatePolyline(jsonLine.Id, coords, jsonLine.Meta)
}

func createPolygon(jsonLine JSONPolygon) Polygon {
	var coords = make([]geom.Coords, 0, len(jsonLine.Coordinates))
	for _, array := range jsonLine.Coordinates {
		coords = append(coords, geom.AsCoordinates(array))
	}
	return CreatePolygon(jsonLine.Id, coords, jsonLine.Meta)
}

func getPointObjects(index int, feat *geojson.Feature) []JSONPoint {
	var objs = make([]JSONPoint, 0, 1)
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	if feat.Geometry.IsPoint() {
		var id = composeId(index, getFId(feat.Properties), 0)
		objs = append(objs, JSONPoint{id, feat.Geometry.Point, string(meta)})
	} else if feat.Geometry.IsMultiPoint() {
		objs = make([]JSONPoint, 0, len(feat.Geometry.MultiPoint))
		for pos, coords := range feat.Geometry.MultiPoint {
			var id = composeId(index, getFId(feat.Properties), pos)
			objs = append(objs, JSONPoint{id, coords, string(meta)})
		}
	}
	return objs
}

func getLineStringObjects(index int, feat *geojson.Feature) []JSONLineString {
	var objs = make([]JSONLineString, 0, 1)
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	if feat.Geometry.IsLineString() {
		var id = composeId(index, getFId(feat.Properties), 0)
		objs = append(objs, JSONLineString{id, feat.Geometry.LineString, string(meta)})
	} else if feat.Geometry.IsMultiLineString() {
		objs = make([]JSONLineString, 0, len(feat.Geometry.MultiLineString))
		for pos, coords := range feat.Geometry.MultiLineString {
			var id = composeId(index, getFId(feat.Properties), pos)
			objs = append(objs, JSONLineString{id, coords, string(meta)})
		}
	}
	return objs
}

func getPolygonObjects(index int, feat *geojson.Feature) []JSONPolygon {
	var objs = make([]JSONPolygon, 0, 1)
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	if feat.Geometry.IsPolygon() {
		var id = composeId(index, getFId(feat.Properties), 0)
		objs = append(objs, JSONPolygon{id, feat.Geometry.Polygon, string(meta)})
	} else if feat.Geometry.IsMultiPolygon() {
		objs = make([]JSONPolygon, 0, len(feat.Geometry.MultiPolygon))
		for pos, coords := range feat.Geometry.MultiPolygon {
			var id = composeId(index, getFId(feat.Properties), pos)
			objs = append(objs, JSONPolygon{id, coords, string(meta)})
		}
	}
	return objs
}

func getFId(properties map[string]interface{}) string {
	var id = properties["id"]
	if id == nil {
		return "?"
	}
	return fmt.Sprintf("%v", id)
}

func composeId(index int, fid string, pos int) string {
	return fmt.Sprintf("idx:%v-fid:%v-pos:%v", index, fid, pos)
}

func readJsonFile(file string) []string {
	var data, err = fileutil.ReadAllOfFile(file)
	checkError(err)
	var tokens = strings.Split(strings.TrimSpace(data), "\n")
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	return tokens
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
