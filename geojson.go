package geometry

import (
	"encoding/json"
	"fmt"
	"github.com/intdxdt/fileutil"
	"github.com/intdxdt/geom"
	geojson "github.com/paulmach/go.geojson"
	"strings"
)

const (
	TypePoint             string = "Point"
	TypeLineString        string = "LineString"
	TypePolygon           string = "Polygon"
	TypeMultiPoint        string = "MultiPoint"
	TypeMultiLineString   string = "MultiLineString"
	TypeMultiPolygon      string = "MultiPolygon"
	TypeFeature           string = "Feature"
	TypeFeatureCollection string = "FeatureCollection"
)

type IGeometry interface {
	Geometry() geom.Geometry
}

type JSONType struct {
	Type string `json:"type"`
}

func (o *JSONType) isGeometryType() bool {
	return o.Type == TypePoint || o.Type == TypeLineString || o.Type == TypePolygon ||
		o.Type == TypeMultiPoint || o.Type == TypeMultiLineString || o.Type == TypeMultiPolygon
}

func (o *JSONType) isFeatureType() bool {
	return o.Type == TypeFeature
}

func (o *JSONType) isFeatureCollectionType() bool {
	return o.Type == TypeFeatureCollection
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
	var tokens = readJsonFile(inputJsonFile)
	return parseInputLinearFeatures(tokens)
}

func ReadInputConstraints(inputJsonFile string) []IGeometry {
	var tokens = readJsonFile(inputJsonFile)
	return parseConstraintFeatures(tokens)
}

func parseInputLinearFeatures(inputs []string) []Polyline {
	var plns = make([]Polyline, 0, len(inputs))
	for idx, fjson := range inputs {
		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)
		var objs = lineStringFromFeature(idx, feat)
		for _, o := range objs {
			var pln = createPolyline(o)
			plns = append(plns, pln)
		}
	}
	return plns
}

func parseConstraintFeatures(inputs []string) []IGeometry {
	var geometries = make([]IGeometry, 0, len(inputs))

	for idx, fjson := range inputs {
		var jtype JSONType
		var err = json.Unmarshal([]byte(fjson), &jtype)

		if jtype.isGeometryType() {
			g, err := geojson.UnmarshalGeometry([]byte(fjson))
			checkError(err)
			println(g)
		} else if jtype.isFeatureType() {

		}

		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)

		var ptObjs = pointsFromFeature(idx, feat)
		for _, o := range ptObjs {
			geometries = append(geometries, createPoint(o))
		}

		var lnObjs = lineStringFromFeature(idx, feat)
		for _, o := range lnObjs {
			geometries = append(geometries, createPolyline(o))
		}

		var plyObjs = polygonFromFeature(idx, feat)
		for _, o := range plyObjs {
			geometries = append(geometries, createPolygon(o))
		}
	}

	return geometries
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

func getFId(properties map[string]interface{}) string {
	var id = properties["mn id"]
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
