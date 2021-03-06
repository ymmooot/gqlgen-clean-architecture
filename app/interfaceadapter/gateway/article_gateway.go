package gateway

import (
	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase/repository"
)

type articleGateway struct {
	sql SQLHandler
}

func NewArticleGateway(sql SQLHandler) repository.ArticleRepository {
	return &articleGateway{sql: sql}
}

func (g *articleGateway) Find(articleID string) (*data.Article, error) {
	rows, err := g.sql.Query("SELECT * FROM articles WHERE id = ?", articleID)
	if err != nil {
		return nil, err
	}

	var article *data.Article

	for rows.Next() {
		article = &data.Article{}
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.PublishedAt,
			&article.CreatedAt, &article.UpdatedAt, &article.DeletedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return article, nil
}
