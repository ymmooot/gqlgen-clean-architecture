package usecase

import (
	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase/repository"
)

type ArticleUseCase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUseCase(r repository.ArticleRepository) ArticleUseCase {
	return ArticleUseCase{r}
}

func (a ArticleUseCase) Find(articleID string) (*data.Article, error) {
	return a.articleRepository.Find(articleID)
}
