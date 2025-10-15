package seeding


type Junction struct {
	ParentID 	int32
	ChildID  	int32
}


func (j Junction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
	}
}


func createJunction[T any, P HasID, C HasID](parent P, key T, lookup func(T) (C, error)) (Junction, error) {
	child, err := lookup(key)
	if err != nil {
		return Junction{}, err
	}

	junction := Junction{
		ParentID: 	*parent.GetID(),
		ChildID: 	*child.GetID(),
	}

	return junction, nil
}