package api

import (
	"encoding/json"
	"net/http"

	"github.com/brtkozak/routes/osrmservice"
)

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// extract routes request dto from query params
	routesRequest, error := GetRoutesRequest(queryParams)
	if error != nil {
		http.Error(w, error.Error.Error(), error.HTTPCode)
	} else {
		// prepare routes response with distances using osrmserivce
		routesResponse, error := osrmservice.GetRoutes(routesRequest)
		if error != nil {
			http.Error(w, error.Error.Error(), error.HTTPCode)
		} else {
			// sort routes and return response
			SortRoutes(routesResponse.Routes)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(routesResponse)
		}
	}
}
