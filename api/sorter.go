package api

import (
	"sort"

	"github.com/brtkozak/routes/dto"
)

//sort routes by duration - if equal then compare distance
func SortRoutes(routes []dto.Route) {
	sort.SliceStable(routes, func(i, j int) bool {
		if routes[i].Duration == routes[j].Duration {
			return routes[i].Distance < routes[j].Distance
		}
		return routes[i].Duration < routes[j].Duration
	})
}
