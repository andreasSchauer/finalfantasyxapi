package helpers

import (
	"database/sql"
)

type Zeroable interface {
	IsZero() bool
}

func NilOrPtr[T Zeroable](v T) *T {
	if v.IsZero() {
		return nil
	}
	return &v
}

// use these for dealing with pointer-fields when inputting values that can be null for a database query
// can also write functions that do the opposite with the sqlc return values (return pointers)

func GetNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func GetNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func GetNullFloat64(f *float32) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: float64(*f), Valid: true}
}

// Use derefOrNil for every field that is a pointer

func DerefOrNil[T any](ptr *T) any {
	if ptr == nil {
		return nil
	}
	return *ptr
}

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

func NullStringToPtr(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	val := s.String
	return &val
}

func NullStringToVal(s sql.NullString) string {
	if !s.Valid {
		return ""
	}

	return s.String
}

func NullInt32ToPtr(i sql.NullInt32) *int32 {
	if !i.Valid {
		return nil
	}

	val := i.Int32
	return &val
}

func NullInt32ToVal(i sql.NullInt32) int32 {
	if !i.Valid {
		return 0
	}

	return i.Int32
}

func NullFloat64ToPtr(f sql.NullFloat64) *float32 {
	if !f.Valid {
		return nil
	}

	val := float32(f.Float64)
	return &val
}

func NullFloat64ToVal(f sql.NullFloat64) float32 {
	if !f.Valid {
		return 0
	}

	return float32(f.Float64)
}