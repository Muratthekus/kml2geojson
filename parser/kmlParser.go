package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"kml2geojson/model"
	"os"
)

func UnMarshallKml(filePath string) model.Folder {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var root model.Root

	err = xml.Unmarshal(byteValue, &root)
	if err != nil {
		fmt.Println(err, "--", filePath)
		panic(err)
	}

	return root.Folder
}
