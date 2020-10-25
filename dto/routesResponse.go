package dto

type RoutesResponse struct {
	Source string
	Routes []Route
}

type Route struct {
	Destination string
	Duration    float64
	Distance    float64
}
