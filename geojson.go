package geometry

import (
	"strings"
)

const (
	pointType           = "Point"
	multiPointType      = "MultiPoint"
	lineStringType      = "LineString"
	multiLineStringType = "MultiLineString"
	polygonType         = "Polygon"
	multiPolygonType    = "MultiPolygon"
)

type GeoJSON struct {
	Type string `json:"type"`
}

func (gjson *GeoJSON) IsPoint() bool {
	return strings.ToLower(gjson.Type) == strings.ToLower(pointType)
}

func (gjson *GeoJSON) IsMultiPoint() bool {
	return strings.ToLower(gjson.Type) == strings.ToLower(multiPointType)
}

func (gjson *GeoJSON) IsLineString() bool {
	return strings.ToLower(gjson.Type) == strings.ToLower(lineStringType)
}

func (gjson *GeoJSON) IsMultiLineString() bool {
	return strings.ToLower(gjson.Type) == strings.ToLower(multiLineStringType)
}

func (gjson *GeoJSON) IsPolygon() bool {
	return strings.ToLower(gjson.Type) == strings.ToLower(polygonType)
}

func (gjson *GeoJSON) IsMultiPolygon() bool {
	return strings.ToLower(gjson.Type) == strings.ToLower(multiPolygonType)
}

type GeoJSONGeometries struct {
	Points      []Point
	LineStrings []Polyline
	Polygons    []Polygon
}
