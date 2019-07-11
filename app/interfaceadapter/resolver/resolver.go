package resolver

import (
	"context"
	"strconv"

	gqlgen_clean_architecture "github.com/ymmooot/gqlgen-clean-architecture"
	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
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
	return &articleResolver{r}
}
func (r *Resolver) Query() gqlgen_clean_architecture.QueryResolver {
	return &queryResolver{r}
}

type articleResolver struct{ *Resolver }

func (r *articleResolver) ID(ctx context.Context, obj *data.Article) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Article(ctx context.Context, id *string) (*data.Article, error) {
	res, err := r.articleRepository.Find(*id)
	return res, err
}
