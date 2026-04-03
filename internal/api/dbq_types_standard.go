package api

import (
	"context"
	"database/sql"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type DbQueryNoInput func(context.Context) ([]int32, error)
type DbQueryIntOne func(context.Context, int32) (int32, error)
type DbQueryIntMany func(context.Context, int32) ([]int32, error)
type DbQueryIntList func(context.Context, []int32) ([]int32, error)
type DbQueryNullIntOne func(context.Context, sql.NullInt32) (int32, error)
type DbQueryNullIntMany func(context.Context, sql.NullInt32) ([]int32, error)
type DbQueryBool func(context.Context, bool) ([]int32, error)
type DbQueryEnum[E any] func(context.Context, E) ([]int32, error)
type DbQueryEnumList[E any] func(context.Context, []E) ([]int32, error)
type DbQueryNullEnum[N any] func(context.Context, N) ([]int32, error)
type DbQueryNullEnumList[N any] func(context.Context, []N) ([]int32, error)
type DbQueryAny func(context.Context, any) ([]int32, error)


func ToIntManyNull(q DbQueryNullIntMany) DbQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return q(ctx, h.GetNullInt32(&id))
	}
}

func ToIntManyAny(q DbQueryAny) DbQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return q(ctx, id)
	}
}

func ToIntManyNoInput(q DbQueryNoInput) DbQueryIntMany {
	return func(ctx context.Context, _ int32) ([]int32, error) {
		return q(ctx)
	}
}

func ToIntOneNull(q DbQueryNullIntOne) DbQueryIntOne {
	return func(ctx context.Context, id int32) (int32, error) {
		return q(ctx, h.GetNullInt32(&id))
	}
}

func ToEnumQuery[E, N any](et EnumType[E, N], q DbQueryNullEnum[N]) DbQueryEnum[E] {
	return func(ctx context.Context, enum E) ([]int32, error) {
		return q(ctx, et.getNullEnum(&enum))
	}
}

func ToEnumListQuery[E, N any](et EnumType[E, N], q DbQueryNullEnumList[N]) DbQueryEnumList[E] {
	return func(ctx context.Context, enums []E) ([]int32, error) {
		nullEnums := []N{}

		for _, enum := range enums {
			nullEnum := et.getNullEnum(&enum)
			nullEnums = append(nullEnums, nullEnum)
		}

		return q(ctx, nullEnums)
	}
}