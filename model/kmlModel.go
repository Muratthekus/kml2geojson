package model

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Root struct {
	XMLName xml.Name `xml:"kml"`
	Folder  Folder   `xml:"Folder"`
}

type Folder struct {
	XMLName  xml.Name `xml:"Folder"`
	Name     string   `xml:"name"`
	Document Document `xml:"Document"`
}

type Document struct {
	XMLName   xml.Name    `xml:"Document"`
	Name      string      `xml:"name"`
	Placemark []Placemark `xml:"Placemark"`
}

type Placemark struct {
	XMLName      xml.Name     `xml:"Placemark"`
	Name         string       `xml:"name"`
	ExtendedData ExtendedData `xml:"ExtendedData"`
	Polygon      Polygon      `xml:"Polygon"`
}

type ExtendedData struct {
	XMLName xml.Name            `xml:"ExtendedData"`
	Data    []ExtendedDataArray `xml:"Data"`
}

type ExtendedDataArray struct {
	XMLName xml.Name `xml:"Data"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value"`
}

// make more generic

type Polygon struct {
	XMLName       xml.Name      `xml:"Polygon"`
	OuterBoundary OuterBoundary `xml:"outerBoundaryIs"`
	InnerBoundary InnerBoundary `xml:"innerBoundaryIs"`
}

type OuterBoundary struct {
	XMLName    xml.Name   `xml:"outerBoundaryIs"`
	LinearRing LinearRing `xml:"LinearRing"`
}

type InnerBoundary struct {
	XMLName    xml.Name   `xml:"innerBoundaryIs"`
	LinearRing LinearRing `xml:"LinearRing"`
}

type LinearRing struct {
	XMLName     xml.Name `xml:"LinearRing"`
	Coordinates string   `xml:"coordinates"`
}

func (ring LinearRing) GetAsPointArray() [][]float64 {
	result := make([][]float64, 0)
	coordinates := strings.Replace(strings.TrimSpace(ring.Coordinates), ",0", ",", -1)
	coordinateArray := strings.Split(coordinates, ",")
	for i := 0; i < len(coordinateArray); i++ {
		coordinateArray[i] = strings.TrimSpace(coordinateArray[i])
	}
	for i := 0; i < len(coordinateArray)-1; i += 2 {
		temp := make([]float64, 0)
		temp = append(temp, getAsFloat64(coordinateArray[i]))
		temp = append(temp, getAsFloat64(coordinateArray[i+1]))
		result = append(result, [][]float64{temp}...)
	}

	return result
}

func getAsFloat64(str string) float64 {
	s, _ := strconv.ParseFloat(str, 64)
	return s
}
