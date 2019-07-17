package resolver

import (
	"context"
	"strconv"

	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
)

type articleResolver struct{}

func (r *articleResolver) ID(ctx context.Context, obj *data.Article) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

func (r *articleResolver) PublishedAt(ctx context.Context, obj *data.Article) (*string, error) {
	if obj.PublishedAt == nil {
		return nil, nil
	}

	d := obj.PublishedAt.String()

	return &d, nil
}

func (r *articleResolver) CreatedAt(ctx context.Context, obj *data.Article) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r *articleResolver) UpdatedAt(ctx context.Context, obj *data.Article) (string, error) {
	return obj.UpdatedAt.String(), nil
}

func (r *articleResolver) DeletedAt(ctx context.Context, obj *data.Article) (*string, error) {
	if obj.DeletedAt == nil {
		return nil, nil
	}

	d := obj.DeletedAt.String()

	return &d, nil
}
