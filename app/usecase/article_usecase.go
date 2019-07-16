package usecase

import (
	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase/repository"
)

type ArticleUseCase interface {
	Find(articleID string) (*data.Article, error)
}

type ArticleUseCaseImpl struct {
	articleRepository *repository.ArticleRepository
}

func NewArticleUseCaseImpl(r *repository.ArticleRepository) ArticleUseCase {
	return &ArticleUseCaseImpl{r}
}

func (a *ArticleUseCaseImpl) Find(articleID string) (*data.Article, error) {
	return (*a.articleRepository).Find(articleID)
}
