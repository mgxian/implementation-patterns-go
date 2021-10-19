package repository

import "article/domain/model"

type ArticleRepository interface {
	Save(article model.Article) (model.Article, error)
	ExistsBySlug(slug string) bool
}
