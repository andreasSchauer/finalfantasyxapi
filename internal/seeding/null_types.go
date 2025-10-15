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

// used for nullable objects to get their fk ID
func ptrObjIDToNullInt32[T HasID](objPtr *T) sql.NullInt32 {
	if objPtr == nil {
		return sql.NullInt32{}
	}

	obj := *objPtr
	id := obj.GetID()
	return getNullInt32(id)
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
