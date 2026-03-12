package helpers

import "database/sql"

// used to get the id of nullable objects for DerefOrNil
func ObjPtrToID[T HasID](objPtr *T) any {
	if objPtr == nil {
		return nil
	}

	obj := *objPtr
	id := obj.GetID()

	if id == 0 {
		return nil
	}
	return id
}

// used to get the id of nullable objects for the db query
func ObjPtrToNullInt32ID[T HasID](objPtr *T) sql.NullInt32 {
	if objPtr == nil {
		return sql.NullInt32{}
	}

	obj := *objPtr
	id := obj.GetID()
	return GetNullInt32(&id)
}

// used to get the id of nullable objects for the db query where the db expects a not null value
func ObjPtrToInt32ID[T HasID](objPtr *T) int32 {
	if objPtr == nil {
		return 0
	}

	obj := *objPtr
	return obj.GetID()
}