// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	CreatePeriod(ctx context.Context, arg CreatePeriodParams) (Period, error)
}

var _ Querier = (*Queries)(nil)
