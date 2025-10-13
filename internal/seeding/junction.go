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