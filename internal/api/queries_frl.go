package api

type filteredResList[A APIResource] struct {
	resources []A
	err       error
}

func frl[A APIResource](res []A, err error) filteredResList[A] {
	return filteredResList[A]{
		resources: res,
		err:       err,
	}
}