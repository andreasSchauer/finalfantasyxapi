package api

// might potentially need more than that
// (thinking about slice elements, but I don't think they'll need it?)
type ServiceParams interface {
	GetDoc(*Config) ParamsDoc
}

type ServiceResponse interface {
	HasURL
}