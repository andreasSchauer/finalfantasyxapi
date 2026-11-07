package api

type SampleParams[P any] struct {
	Description 	string		`json:"description"`
	Params			P			`json:"params"`
}