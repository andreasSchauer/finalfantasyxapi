package api

import (
	"context"
	"database/sql"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func queryMany(query func(context.Context, sql.NullInt32) ([]int32, error)) func(context.Context, int32) ([]int32, error) {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return query(ctx, h.GetNullInt32(&id))
	}
}

func queryNoInput(query func(context.Context) ([]int32, error)) func(context.Context, int32) ([]int32, error) {
	return func(ctx context.Context, _ int32) ([]int32, error) {
		return query(ctx)
	}
}

func queryOne(query func(context.Context, sql.NullInt32) (int32, error)) func(context.Context, int32) (int32, error) {
	return func(ctx context.Context, id int32) (int32, error) {
		return query(ctx, h.GetNullInt32(&id))
	}
}