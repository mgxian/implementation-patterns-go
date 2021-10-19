package service

import "article/domain/model"

type ArticleServiceStub struct {
	CreateArticleFn func(title, description, body string, authorId int) (model.Article, error)
}

func NewArticleServiceStub() ArticleServiceStub {
	return ArticleServiceStub{}
}

func (s ArticleServiceStub) CreateArticle(title string, description string, body string, authorId int) (model.Article, error) {
	return s.CreateArticleFn(title, description, body, authorId)
}
