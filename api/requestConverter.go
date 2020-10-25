package api

import (
	"../dto"
)

const (
	queryParamSrc = "src"
	queryParamDst = "dst"
)

func GetRoutesRequest(queryParams map[string][]string) (dto.RoutesRequest, *dto.RequestError) {
	var routesRequest dto.RoutesRequest

	querySrc := queryParams[queryParamSrc]
	queryDst := queryParams[queryParamDst]

	// check if exactly one source and at least one destination were given
	if len(querySrc) != 1 || len(queryDst) == 0 {
		return routesRequest, dto.InvalidRequestDataError()
	}

	// convert passed src query parameter to Coordinate stuct
	src, error := dto.GetCoordinatesFromString(querySrc[0])
	if error != nil {
		return routesRequest, error
	}

	// convert all passed dst query parameters to Coordinate struct
	var dests []dto.Coordinates
	for _, item := range queryDst {
		if coord, error := dto.GetCoordinatesFromString(item); error == nil {
			dests = append(dests, coord)
		} else {
			return routesRequest, error
		}

	}

	routesRequest.Src = src
	routesRequest.Dests = dests
	return routesRequest, nil
}
