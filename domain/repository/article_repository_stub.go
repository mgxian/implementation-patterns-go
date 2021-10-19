package repository

import (
	"article/domain/model"
	"time"
)

type ArticleRepositoryStub struct{}

func NewArticleRepositoryStub() ArticleRepositoryStub {
	return ArticleRepositoryStub{}
}

func (r *ArticleRepositoryStub) Save(article model.Article) (model.Article, error) {
	now := time.Now()
	article.SetCreatedAt(now)
	article.SetUpdatedAt(now)
	return article, nil
}
