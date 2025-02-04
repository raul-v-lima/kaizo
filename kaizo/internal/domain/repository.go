package domain

import (
	"context"
)

type MonRepositoryImpl interface {
	GetMon(ctx context.Context, id string) (*Mon, error)
	GetAllMons(ctx context.Context) ([]*Mon, error)
	CreateMon(ctx context.Context, mon *Mon) error
	UpdateMon(ctx context.Context, mon *Mon) error
	DeleteMon(ctx context.Context, id string) error
}
