package tests

import (
	"testing"

	"github.com/brtkozak/routes/api"
	"github.com/brtkozak/routes/dto"
)

func TestgetRoutesRequest(t *testing.T) {
	var queryParams map[string][]string

	src := []string{
		"33.123456;33.123456",
	}
	dst := []string{
		"34.000000;34.000000",
		"35.000000;35.000000",
		"36.000000;36.000000",
	}

	queryParams["src"] = src
	queryParams["dst"] = dst

	routesRequestDto, error := api.GetRoutesRequest(queryParams)

	expectedStruct := dto.RoutesRequest{
		Src: dto.Coordinates{
			Lat:  33.123456,
			Long: 33.123456,
		},
		Dests: []dto.Coordinates{
			dto.Coordinates{
				Lat:  34.000000,
				Long: 34.000000,
			},
			dto.Coordinates{
				Lat:  35.000000,
				Long: 35.000000,
			},
			dto.Coordinates{
				Lat:  36.000000,
				Long: 36.000000,
			},
		},
	}

	if routesRequestDto.Src != expectedStruct.Src {
		t.Error("Routes request converting test failed")
	}
	if routesRequestDto.Dests[0] != expectedStruct.Dests[0] {
		t.Error("Routes request converting test failed")
	}
	if routesRequestDto.Dests[1] != expectedStruct.Dests[1] {
		t.Error("Routes request converting test failed")
	}
	if routesRequestDto.Dests[2] != expectedStruct.Dests[2] {
		t.Error("Routes request converting test failed")
	}

	//check case when no src is given - expect error
	src = src[:0]
	routesRequestDto, error = api.GetRoutesRequest(queryParams)
	if error == nil {
		t.Error("Routes request converting test failed")
	}

	//src is given again - expext success
	src[0] = "33.123456;33.123456"
	routesRequestDto, error = api.GetRoutesRequest(queryParams)
	if error != nil {
		t.Error("Routes request converting test failed")
	}

	// two srcs are given - expext error
	src[1] = "33.123456;33.123456"
	routesRequestDto, error = api.GetRoutesRequest(queryParams)
	if error == nil {
		t.Error("Routes request converting test failed")
	}

	// one src is given but zero dst - expext error
	src = src[:1]
	dst = dst[:0]
	routesRequestDto, error = api.GetRoutesRequest(queryParams)
	if error == nil {
		t.Error("Routes request converting test failed")
	}

}
