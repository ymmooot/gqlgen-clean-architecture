package gateway

import (
	"github.com/ymmooot/gqlgen-clean-architecture/app/entity/data"
)

type articleGateway struct {
	sql SqlHandler
}

func NewArticleGateway(sql SqlHandler) *articleGateway {
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
		err = rows.Scan(&article.ID, &article.Title)
	}

	return article, nil
}
