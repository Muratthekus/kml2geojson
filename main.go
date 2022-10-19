package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kml2geojson/parser"
	"strings"
)

const DIR = "./filesv2/" //
const OUTPUT_DIR = "./outputs/"

func main() {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		panic(err)
	}

	for index, f := range files {
		filePath := DIR + f.Name()
		fileName := OUTPUT_DIR + strings.TrimSuffix(f.Name(), ".kml") + ".geojson"
		var kml = parser.UnMarshallKml(filePath)
		var geoJson = parser.ToGeoJson(kml)
		file, _ := json.MarshalIndent(geoJson, "", " ")
		_ = ioutil.WriteFile(fileName, file, 0644)
		fmt.Println(index)
	}
}
