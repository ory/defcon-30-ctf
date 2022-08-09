package main

import (
	"context"
)

type (
	result struct {
		DistrictID uint            `json:"district_id"`
		Votes      map[string]uint `json:"votes"`
	}

	repository interface {
		List(ctx context.Context) ([]*result, error)
		Submit(ctx context.Context, r *result) error
	}
)
