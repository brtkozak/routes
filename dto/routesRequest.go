package dto

import (
	"fmt"
	"strconv"
	"strings"
)

type RoutesRequest struct {
	Src   Coordinates
	Dests []Coordinates
}

type Coordinates struct {
	Lat  float64
	Long float64
}

// convert coordinates struct to string
func GetCoordinatesAsString(coordinate Coordinates) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%f", coordinate.Lat))
	sb.WriteString(",")
	sb.WriteString(fmt.Sprintf("%f", coordinate.Long))
	return sb.String()
}

//convert coorindates from string to Coordinates struct
func GetCoordinatesFromString(latlong string) (Coordinates, *RequestError) {
	var coordinates Coordinates

	splited := strings.Split(latlong, ",")
	if len(splited) != 2 {
		return coordinates, InvalidRequestDataError()
	}
	latitude, err := strconv.ParseFloat(splited[0], 64)
	if err != nil {
		return coordinates, InvalidRequestDataError()
	}
	longitude, err := strconv.ParseFloat(splited[1], 64)
	if err != nil {
		return coordinates, InvalidRequestDataError()
	}

	coordinates.Lat = latitude
	coordinates.Long = longitude
	return coordinates, nil
}
