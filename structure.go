package main

type Sale struct {
	Name string
	Lat  float64
	Long float64
}

type Step struct {
	Sale                  *Sale
	PreviousStep          *Step
	DistanceFromLastPoint float64
	TotalCoveredDistance  float64
}

type Route struct {
	Steps []Step
}

// convert types take an int and return a string value.
type CallbackEnd func()
