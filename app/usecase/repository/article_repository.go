package repository

import "github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"

type ArticleRepository interface {
	Find(articleID string) (*data.Article, error)
}
