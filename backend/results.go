package main

import (
	"context"
)

type (
	result struct {
		District    string `json:"district"`
		Democrats   uint   `json:"democrats"`
		Republicans uint   `json:"republicans"`
		Invalid     uint   `json:"invalid"`
	}

	repository interface {
		List(ctx context.Context) ([]*result, error)
		Submit(ctx context.Context, district string, r *result) error
	}
)
