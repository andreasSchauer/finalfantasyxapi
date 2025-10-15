package seeding

import (
	"database/sql"
)

// use these for dealing with pointer-fields when inputting values that can be null for a database query
// can also write functions that do the opposite with the sqlc return values (return pointers)

func getNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func getNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

// used to get the id nullable objects for the db query
func ObjPtrToNullInt32ID[T HasID](objPtr *T) sql.NullInt32 {
	if objPtr == nil {
		return sql.NullInt32{}
	}

	obj := *objPtr
	id := obj.GetID()
	return getNullInt32(&id)
}


// used to get the id of nullable objects for the db query where the db expects a not null value
func ObjPtrToInt32ID[T HasID](objPtr *T) int32 {
	if objPtr == nil {
		return 0
	}

	obj := *objPtr
	return obj.GetID()
}


func getNullFloat64(f *float32) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: float64(*f), Valid: true}
}

func stringPtrToString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

/*

func convertNullString(s sql.NullString) *string {
    if !s.Valid {
        return nil
    }

    val := s.String
    return &val
}


func convertNullInt32(i sql.NullInt32) *int32 {
    if !i.Valid {
        return nil
    }

    val := i.Int32
    return &val
}

*/
