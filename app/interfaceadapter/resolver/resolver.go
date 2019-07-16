package resolver

import (
	gqlgen_clean_architecture "github.com/ymmooot/gqlgen-clean-architecture"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase/repository"
)

type Resolver struct {
	articleRepository repository.ArticleRepository
}

func NewResolver(r repository.ArticleRepository) *Resolver {
	return &Resolver{
		articleRepository: r,
	}
}

func (r *Resolver) Article() gqlgen_clean_architecture.ArticleResolver {
	return &articleResolver{}
}

func (r *Resolver) Query() gqlgen_clean_architecture.QueryResolver {
	return NewQueryResolver(r.articleRepository)
}
