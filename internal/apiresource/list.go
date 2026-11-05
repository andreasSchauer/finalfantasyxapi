package apiresource

type ListParams struct {
	Count    int     `json:"count"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}

type APIResourceList[A APIResource] struct {
	ListParams
	Results 	[]A `json:"results"`
}

func (l APIResourceList[A]) GetListParams() ListParams {
	return l.ListParams
}

func (l APIResourceList[A]) GetResults() []A {
	return l.Results
}

type SimpleResourceList struct {
	ListParams
	ParentResource APIResource 			`json:"parent_resource,omitempty"`
	Results        []SimpleResource     `json:"results"`
}

func (l SimpleResourceList) GetListParams() ListParams {
	return l.ListParams
}

func (l SimpleResourceList) GetResults() []SimpleResource {
	return l.Results
}