package common

type Topic string

const (
	DijkstraErrors Topic = "errors.dijkstra"
	NamiErrors     Topic = "errors.nami"
	SuezErrors     Topic = "errors.suez"
)
