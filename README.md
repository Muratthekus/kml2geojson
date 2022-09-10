## Kml to GeoJson Converter -In Progress- ##

Basic Kml to GeoJson convert app. It reads all the Kml file 
under the `inputs` directory, converts each to GeoJson and write in `outputs` directory.



```go
package main
func main() {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		filePath := DIR + f.Name()
		fileName := OUTPUT_DIR + strings.TrimSuffix(f.Name(), ".kml") + ".geojson"
		var kml = parser.UnMarshallKml(filePath)
		var geoJson = parser.ToGeoJson(kml)
		file, _ := json.MarshalIndent(geoJson, "", " ")
		
		_ = ioutil.WriteFile(fileName, file, 0644)
	}
}
```


