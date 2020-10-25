package tests

import (
	"testing"

	"github.com/brtkozak/routes/api"
	"github.com/brtkozak/routes/dto"
)

func TestSort(t *testing.T) {
	routes := []dto.Route{
		dto.Route{
			Destination: "1",
			Duration:    200,
			Distance:    100,
		},
		dto.Route{
			Destination: "2",
			Duration:    300,
			Distance:    100,
		},
		dto.Route{
			Destination: "0",
			Duration:    100,
			Distance:    100,
		},
	}
	api.SortRoutes(routes)
	if routes[0].Destination != "0" || routes[1].Destination != "1" || routes[2].Destination != "2" {
		t.Error("Sorting test failed")
	}

	routes = []dto.Route{
		dto.Route{
			Destination: "1",
			Duration:    200,
			Distance:    80,
		},
		dto.Route{
			Destination: "2",
			Duration:    200,
			Distance:    100,
		},
		dto.Route{
			Destination: "0",
			Duration:    200,
			Distance:    50,
		},
	}
	api.SortRoutes(routes)
	if routes[0].Destination != "0" || routes[1].Destination != "1" || routes[2].Destination != "2" {
		t.Error("Sorting test failed")
	}

	api.SortRoutes(routes)
	if routes[0].Destination != "0" || routes[1].Destination != "1" || routes[2].Destination != "2" {
		t.Error("Sorting test failed")
	}

	routes = []dto.Route{
		dto.Route{
			Destination: "2",
			Duration:    200,
			Distance:    11,
		},
		dto.Route{
			Destination: "1",
			Duration:    200,
			Distance:    10,
		},
		dto.Route{
			Destination: "0",
			Duration:    100,
			Distance:    1000,
		},
	}
	api.SortRoutes(routes)
	if routes[0].Destination != "0" || routes[1].Destination != "1" || routes[2].Destination != "2" {
		t.Error("Sorting test failed")
	}

	routes = []dto.Route{
		dto.Route{
			Destination: "0",
			Duration:    1,
			Distance:    1000,
		},
		dto.Route{
			Destination: "1",
			Duration:    2,
			Distance:    100,
		},
		dto.Route{
			Destination: "2",
			Duration:    3,
			Distance:    10,
		},
	}
	api.SortRoutes(routes)
	if routes[0].Destination != "0" || routes[1].Destination != "1" || routes[2].Destination != "2" {
		t.Error("Sorting test failed")
	}
}
