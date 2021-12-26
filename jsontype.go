package geometry

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
