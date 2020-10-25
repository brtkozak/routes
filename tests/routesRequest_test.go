package tests

import (
	"encoding/json"
	"testing"

	"../dto"
)

func TestConvertStringToCoordinates(t *testing.T) {
	s := "00:000000,00:000000"
	coordinates, error := dto.GetCoordinatesFromString(s)
	if coordinates.Lat != 0 || coordinates.Long != 0 {
		t.Errorf("String to Coordinates test failed for %s", s)
	}

	s = "1,2"
	coordinates, error = dto.GetCoordinatesFromString(s)
	if coordinates.Lat != 1 || coordinates.Long != 2 {
		t.Errorf("String to Coordinates test failed for %s", s)
	}

	s = "99.2383285,0"
	coordinates, error = dto.GetCoordinatesFromString(s)
	if coordinates.Lat != 99.2383285 || coordinates.Long != 0 {
		t.Errorf("String to Coordinates test failed for %s", s)
	}

	s = "99.0fail,0"
	coordinates, error = dto.GetCoordinatesFromString(s)
	if error == nil {
		t.Errorf("String to Coordinates test failed for %s", s)
	}

	s = "99.9,0f"
	coordinates, error = dto.GetCoordinatesFromString(s)
	if error == nil {
		t.Errorf("String to Coordinates test failed for %s", s)
	}
}

func TestConvertCoordinatesToString(t *testing.T) {
	coordinates := dto.Coordinates{
		Lat:  50.999999,
		Long: 50.999999,
	}
	s := dto.GetCoordinatesAsString(coordinates)
	if s != "50.999999,50.999999" {
		json, _ := json.Marshal(coordinates)
		t.Errorf("Coordinates to String test failed for %s", json)
	}

	coordinates = dto.Coordinates{
		Lat:  1,
		Long: 0,
	}
	s = dto.GetCoordinatesAsString(coordinates)
	if s != "1.000000,0.000000" {
		json, _ := json.Marshal(coordinates)
		t.Errorf("Coordinates to String test failed for %s", json)
	}

	coordinates = dto.Coordinates{
		Lat:  0.001,
		Long: 1.0001,
	}
	s = dto.GetCoordinatesAsString(coordinates)
	if s != "0.001000,1.000100" {
		json, _ := json.Marshal(coordinates)
		t.Errorf("Coordinates to String test failed for %s", json)
	}
}
