package resolver

import (
	"context"

	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase"
)

type queryResolver struct {
	articleUsecase usecase.ArticleUseCase
}

func NewQueryResolver(au usecase.ArticleUseCase) *queryResolver {
	return &queryResolver{
		articleUsecase: au,
	}
}

func (r *queryResolver) Article(ctx context.Context, id *string) (*data.Article, error) {
	res, err := r.articleUsecase.Find(*id)
	return res, err
}
