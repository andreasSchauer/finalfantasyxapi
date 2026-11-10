package helpers

import "database/sql"


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


func ObjPtrToNullInt32ID[T HasID](objPtr *T) sql.NullInt32 {
	if objPtr == nil {
		return sql.NullInt32{}
	}

	obj := *objPtr
	id := obj.GetID()
	return GetNullInt32(&id)
}


func ObjPtrToInt32ID[T HasID](objPtr *T) int32 {
	if objPtr == nil {
		return 0
	}

	obj := *objPtr
	return obj.GetID()
}
