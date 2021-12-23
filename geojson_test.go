package geometry

import (
	"encoding/json"
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestGeoJSON_IO(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("GeoJSON", func() {
		g.It("geojson io", func() {
			g.Timeout(1 * time.Minute)
			var obj JSONType
			var dat = `{ "type": "Point", "coordinates": [30.0, 10.0] }`

			var err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsTrue()

			dat = `{ "type": "FeatureCollection", "features": [ { "type": "Feature", "geometry": { "type": "Point", "coordinates": [102.0, 0.5] }, "properties": { "prop0": "value0" } }, { "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }, { "type": "Feature", "geometry": { "type": "Polygon", "coordinates": [ [ [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0] ] ] }, "properties": { "prop0": "value0", "prop1": {"this": "that"} } } ] }`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsTrue()

			dat = `{"type": "Feature", "geometry": {"type": "Polygon", "coordinates": [[[100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0]]]}, "properties": {"prop0": "value0", "prop1": {"this": "that"}}}`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsFalse()
			g.Assert(obj.isFeatureType()).IsTrue()

			dat = `{ "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsFalse()
			g.Assert(obj.isFeatureType()).IsTrue()

			dat = `{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [102.0, 0.5] }, "properties": { "prop0": "value0" } }`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsFalse()
			g.Assert(obj.isFeatureType()).IsTrue()

		})
	})
}
