package resolver

import (
	"context"
	"strconv"

	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
)

type articleResolver struct{}

func (r *articleResolver) ID(ctx context.Context, obj *data.Article) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}
