package api

import (
	"context"
	"database/sql"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type IDFilter func(context.Context) ([]int32, error)
type QueryWrapBasic func(*Config, *http.Request, context.Context, string, QueryParam) ([]int32, error)
type QueryWrapBool func(*Config, *http.Request, context.Context, bool) ([]int32, error)
type QueryWrapEnum[E any] func(*Config, *http.Request, context.Context, E) ([]int32, error)
type QueryWrapIntSlice func(*Config, *http.Request, context.Context, []int32) ([]int32, error)
type QueryWrapInt func(*Config, *http.Request, context.Context, int32) ([]int32, error)
type QueryWrapJoined func(*Config, *http.Request, context.Context) ([]int32, error)

type DBQueryNoInput func(context.Context) ([]int32, error)
type DBQueryIntOne func(context.Context, int32) (int32, error)
type DBQueryIntMany func(context.Context, int32) ([]int32, error)
type DBQueryIntList func(context.Context, []int32) ([]int32, error)
type DBQueryNullIntOne func(context.Context, sql.NullInt32) (int32, error)
type DBQueryNullIntMany func(context.Context, sql.NullInt32) ([]int32, error)
type DBQueryBool func(context.Context, bool) ([]int32, error)
type DBQueryEnum[E any] func(context.Context, E) ([]int32, error)
type DBQueryEnumList[E any] func(context.Context, []E) ([]int32, error)
type DBQueryNullEnum[N any] func(context.Context, N) ([]int32, error)
type DBQueryNullEnumList[N any] func(context.Context, []N) ([]int32, error)
type DBQueryStringMany func(context.Context, string) ([]int32, error)
type DBQueryStringList func(context.Context, []string) ([]int32, error)

func ToIntManyNull(q DBQueryNullIntMany) DBQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return q(ctx, h.GetNullInt32(&id))
	}
}

func ToIntManyNoInput(q DBQueryNoInput) DBQueryIntMany {
	return func(ctx context.Context, _ int32) ([]int32, error) {
		return q(ctx)
	}
}

func ToIntOneNull(q DBQueryNullIntOne) DBQueryIntOne {
	return func(ctx context.Context, id int32) (int32, error) {
		return q(ctx, h.GetNullInt32(&id))
	}
}

func ToEnumQuery[E, N any](et EnumType[E, N], q DBQueryNullEnum[N]) DBQueryEnum[E] {
	return func(ctx context.Context, enum E) ([]int32, error) {
		return q(ctx, et.getNullEnum(&enum))
	}
}

func ToEnumListQuery[E, N any](et EnumType[E, N], q DBQueryNullEnumList[N]) DBQueryEnumList[E] {
	return func(ctx context.Context, enums []E) ([]int32, error) {
		nullEnums := []N{}

		for _, enum := range enums {
			nullEnum := et.getNullEnum(&enum)
			nullEnums = append(nullEnums, nullEnum)
		}

		return q(ctx, nullEnums)
	}
}
