package seeding

type JunctionParams struct {
	DataHashes          []string
	GreatGrandParentIDs []int32 // probably not needed
	GrandParentIDs      []int32
	ParentIDs           []int32
	ChildIDs            []int32
}

func processJunctions[P, C Hashable](l *Lookup, desc string, parents []P, getChildren func(P) ([]C, error)) (JunctionParams, error) {
	params := JunctionParams{
		DataHashes: make([]string, 0),
		ParentIDs:  make([]int32, 0),
		ChildIDs:   make([]int32, 0),
	}

	for _, p := range parents {
		pID, err := l.GetHashID(p)
		if err != nil {
			return JunctionParams{}, err
		}

		children, err := getChildren(p)
		if err != nil {
			return JunctionParams{}, err
		}

		for _, c := range children {
			cID, err := l.GetHashID(c)
			if err != nil {
				return JunctionParams{}, err
			}

			j := StdJunction{
				ParentID: pID,
				ChildID:  cID,
			}
			dataHash := generateJunctionHash(j, desc)

			params.DataHashes = append(params.DataHashes, dataHash)
			params.ParentIDs = append(params.ParentIDs, pID)
			params.ChildIDs = append(params.ChildIDs, cID)
		}
	}

	return params, nil
}

func processThreewayJunctions[GP, P, C Hashable](l *Lookup, desc string, grandParents []GP, getParents func(GP) ([]P, error), getChildren func(P) ([]C, error)) (JunctionParams, error) {
	params := JunctionParams{
		DataHashes:     make([]string, 0),
		GrandParentIDs: make([]int32, 0),
		ParentIDs:      make([]int32, 0),
		ChildIDs:       make([]int32, 0),
	}
	for _, gp := range grandParents {
		gpID, err := l.GetHashID(gp)
		if err != nil {
			return JunctionParams{}, err
		}

		parents, err := getParents(gp)
		if err != nil {
			return JunctionParams{}, err
		}

		for _, p := range parents {
			pID, err := l.GetHashID(p)
			if err != nil {
				return JunctionParams{}, err
			}

			children, err := getChildren(p)
			if err != nil {
				return JunctionParams{}, err
			}

			for _, c := range children {
				cID, err := l.GetHashID(c)
				if err != nil {
					return JunctionParams{}, err
				}

				j := ThreeWayJunction{}
				j.GrandparentID = gpID
				j.ParentID = pID
				j.ChildID = cID
				dataHash := generateJunctionHash(j, desc)

				params.DataHashes = append(params.DataHashes, dataHash)
				params.GrandParentIDs = append(params.GrandParentIDs, gpID)
				params.ParentIDs = append(params.ParentIDs, pID)
				params.ChildIDs = append(params.ChildIDs, cID)
			}
		}
	}

	return params, nil
}
