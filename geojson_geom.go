package geometry

import (
	"encoding/json"
	geojson "github.com/paulmach/go.geojson"
)

func pointsFromFeature(index int, feat *geojson.Feature) []JSONPoint {
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

func lineStringFromFeature(index int, feat *geojson.Feature) []JSONLineString {
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

func polygonFromFeature(index int, feat *geojson.Feature) []JSONPolygon {
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
