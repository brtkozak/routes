package osrmservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"../dto"
)

const (
	baseOsrmAddress = "http://router.project-osrm.org"
)

func GetRoutes(routesRequest dto.RoutesRequest) (dto.RoutesResponse, *dto.RequestError) {
	src := dto.GetCoordinatesAsString(routesRequest.Src)
	routesResponse := dto.RoutesResponse{
		Source: src,
		Routes: make([]dto.Route, 0),
	}

	// Using wait group to get distance from osrm service for every destination
	// Modify routesResponse by passing routes slice as a reference
	// When all request are done return routeResponse
	var wg sync.WaitGroup
	var errorOccured *dto.RequestError
	for _, dest := range routesRequest.Dests {
		wg.Add(1)
		go func(s string, d dto.Coordinates, r *[]dto.Route) {
			err := getDistance(s, d, r)
			if err != nil {
				errorOccured = err
			}
			wg.Done()
		}(src, dest, &routesResponse.Routes)
	}
	wg.Wait()

	if errorOccured != nil {
		return routesResponse, errorOccured
	}

	return routesResponse, nil
}

func getDistance(src string, dst dto.Coordinates, routes *[]dto.Route) *dto.RequestError {
	dstAsString := dto.GetCoordinatesAsString(dst)

	// Build url address
	var sb strings.Builder
	sb.WriteString(baseOsrmAddress)
	sb.WriteString("/route/v1/driving/")
	sb.WriteString(src)
	sb.WriteString(";")
	sb.WriteString(dstAsString)
	sb.WriteString("?overview=false")
	url := sb.String()

	// Make http request to osrm service
	resp, err := http.Get(url)
	if err != nil {
		return dto.ExternalServiceError()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.ExternalServiceError()
	}

	// parse response to osrmResponse dto
	var osrmResponse dto.OsrmResponse
	err = json.Unmarshal(body, &osrmResponse)
	if err != nil {
		return dto.ExternalServiceError()
	}

	// handle osrm service response code
	if osrmResponse.Code != "Ok" {
		return dto.OsrmServiceError(osrmResponse.Code)
	}

	if len(osrmResponse.Routes) == 0 {
		return nil
	}

	// build new Route struct and add it to the slice passed by reference
	newRoute := dto.Route{
		Destination: dstAsString,
		Duration:    osrmResponse.Routes[0].Duration,
		Distance:    osrmResponse.Routes[0].Distance,
	}
	*routes = append(*routes, newRoute)
	return nil
}
