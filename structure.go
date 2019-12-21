package main

type Sale struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Step struct {
	Sale                 *Sale
	PreviousStep          *Step
	DistanceFromLastPoint float64
	TotalCoveredDistance  float64
}

type Route struct {
	Steps []Step
}
