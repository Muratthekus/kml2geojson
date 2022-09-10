package parser

import "kml2geojson/model"

func ToGeoJson(folder model.Folder) model.FeatureCollection {
	return model.FeatureCollection{
		Type:    "FeatureCollection",
		Name:    folder.Document.Name,
		Crs:     model.GetDefaultCrs(),
		Feature: getFeatures(folder.Document.Placemark),
	}
}

func getFeatures(placemarks []model.Placemark) []model.Feature {
	var feature []model.Feature
	for _, placemark := range placemarks {
		feature = append(feature, model.Feature{
			Type:       "Feature",
			Properties: setProperties(placemark.ExtendedData),
			Geometry:   getFeatureGeometry(placemark.Polygon),
		})
	}
	return feature
}

func setProperties(extendedData model.ExtendedData) map[string]interface{} {
	var propertyMap = make(map[string]interface{})
	for _, data := range extendedData.Data {
		propertyMap[data.Name] = data.Value
	}
	return propertyMap
}

//  convert to more generic strcuture as geojsonModel
func getFeatureGeometry(polygon model.Polygon) model.Geometry {
	var geometry model.Geometry

	geometry = model.Geometry{
		Type: model.PolygonType,
	}
	geometry.Coordinates = getCoordinates(geometry, polygon)
	return geometry
}

func getCoordinates(geometry model.Geometry, polygon model.Polygon) interface{} {
	geometry.Coordinates = geometry.GetCoordinate(geometry.Type)
	outBoundary := polygon.OuterBoundary.LinearRing.GetAsPointArray()
	innerBoundary := polygon.InnerBoundary.LinearRing.GetAsPointArray()
	polygonArray := make(model.PolygonArray, 0)
	polygonArray = append(polygonArray, outBoundary)
	if innerBoundary != nil && len(innerBoundary) > 0 {
		polygonArray = append(polygonArray, innerBoundary)
	}
	return polygonArray
}
