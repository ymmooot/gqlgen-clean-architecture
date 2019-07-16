package resolver

import (
	"context"

	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase/repository"
)

type queryResolver struct {
	articleRepository repository.ArticleRepository
}

func NewQueryResolver(r repository.ArticleRepository) *queryResolver {
	return &queryResolver{
		articleRepository: r,
	}
}

func (r *queryResolver) Article(ctx context.Context, id *string) (*data.Article, error) {
	res, err := r.articleRepository.Find(*id)
	return res, err
}
