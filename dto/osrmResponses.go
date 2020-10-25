package dto

type OsrmResponse struct {
	Code   string
	Routes []OsrmRoute
}

type OsrmRoute struct {
	Distance float64
	Duration float64
}
