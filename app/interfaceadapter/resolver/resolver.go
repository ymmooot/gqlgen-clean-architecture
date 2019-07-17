package resolver

import (
	gqlgen_clean_architecture "github.com/ymmooot/gqlgen-clean-architecture"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase"
)

type Resolver struct {
	articleUseCase usecase.ArticleUseCase
}

func NewResolver(au usecase.ArticleUseCase) *Resolver {
	return &Resolver{
		articleUseCase: au,
	}
}

func (r Resolver) Article() gqlgen_clean_architecture.ArticleResolver {
	return &articleResolver{}
}

func (r Resolver) Query() gqlgen_clean_architecture.QueryResolver {
	return NewQueryResolver(r.articleUseCase)
}
