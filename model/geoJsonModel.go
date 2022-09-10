package model

type GeometryType string

const (
	PointType           GeometryType = "Point"
	MultiPointType      GeometryType = "MultiPoint"
	LineStringType      GeometryType = "LineString"
	MultiLineStringType GeometryType = "MultiLineString"
	PolygonType         GeometryType = "Polygon"
	MultiPolygonType    GeometryType = "MultiPolygon"
)

type (
	PointArray           []float64
	MultiPointArray      [][]float64
	LineStringArray      [][]float64
	MultiLineStringArray [][][]float64
	PolygonArray         [][][]float64
	MultiPolygonArray    [][][][]float64
)

type FeatureCollection struct {
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Crs     Crs       `json:"crs"`
	Feature []Feature `json:"features"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   Geometry               `json:"geometry"`
}

type Geometry struct {
	Type        GeometryType `json:"type"`
	Coordinates interface{}  `json:"coordinates"`
}

/*Can be used later*/
/*type CoordinateType struct {
	PointArray           []float64
	MultiPointArray      [][]float64
	LineStringArray      [][]float64
	MultiLineStringArray [][][]float64
	PolygonArray         [][][]float64
	MultiPolygonArray    [][][][]float64
}*/

type Crs struct {
	Type       string        `json:"type"`
	Properties CrsProperties `json:"properties"`
}
type CrsProperties struct {
	Name string `json:"name"`
}

func GetDefaultCrs() Crs {
	var crs Crs
	crs.Type = "name"
	crs.Properties = CrsProperties{
		Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
	}
	return crs
}

func (geometry Geometry) GetCoordinate(geoType GeometryType) Geometry {
	switch geoType {
	case PointType:
		geometry.Coordinates = PointArray{}
	case MultiPointType:
		geometry.Coordinates = MultiPointArray{}
	case LineStringType:
		geometry.Coordinates = LineStringArray{}
	case MultiLineStringType:
		geometry.Coordinates = MultiLineStringArray{}
	case PolygonType:
		geometry.Coordinates = PolygonArray{}
	case MultiPolygonType:
		geometry.Coordinates = MultiPolygonArray{}
	}
	return geometry
}
