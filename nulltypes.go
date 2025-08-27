package main

import (
	"database/sql"
)

// use these for dealing with pointer-fields when inputting values that can be null for a database query

func getNullString(s *string) sql.NullString {
    if s == nil {
        return sql.NullString{}
    }
    return sql.NullString{String: *s, Valid: true}
}

func getNullInt64(i *int) sql.NullInt64 {
    if i == nil {
        return sql.NullInt64{}
    }
    return sql.NullInt64{Int64: int64(*i), Valid: true}
}

func getNullFloat64(f *float64) sql.NullFloat64 {
    if f == nil {
        return sql.NullFloat64{}
    }
    return sql.NullFloat64{Float64: *f, Valid: true}
}

func getNullBool(b *bool) sql.NullBool {
    if b == nil {
        return sql.NullBool{}
    }
    return sql.NullBool{Bool: *b, Valid: true}
}