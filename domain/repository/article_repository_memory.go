package repository

import (
	"article/domain/capability"
	"article/domain/model"
)

type ArticleRepositoryMemory struct {
	clock    capability.Clock
	articles map[string]model.Article
}

func (m *ArticleRepositoryMemory) Save(article model.Article) (model.Article, error) {
	now := m.clock.Now()
	article.SetCreatedAt(now)
	article.SetUpdatedAt(now)
	m.articles[article.Slug()] = article
	return article, nil
}

func NewArticlesRepositoryMemory(clock capability.Clock) *ArticleRepositoryMemory {
	return &ArticleRepositoryMemory{
		clock:    clock,
		articles: make(map[string]model.Article, 0),
	}
}
