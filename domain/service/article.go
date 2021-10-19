package service

import (
	"article/domain/model"
	"article/domain/repository"
	"fmt"
	"strings"
	"time"
)

type ArticleServiceInterface interface {
	CreateArticle(title string, description string, body string, authorId int) (model.Article, error)
}

type ArticleDTO struct {
	Slug        string    `json:"slug,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Body        string    `json:"body,omitempty"`
	AuthorId    int       `json:"author_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func NewArticleDTOFromArticle(article model.Article) ArticleDTO {
	return ArticleDTO{
		Slug:        article.Slug(),
		Title:       article.Title(),
		Description: article.Description(),
		Body:        article.Body(),
		AuthorId:    article.AuthorId(),
		CreatedAt:   article.CreatedAt(),
		UpdatedAt:   article.UpdatedAt(),
	}
}

type ArticleExistedError struct {
	slug string
}

func NewArticleExistsError() ArticleExistedError {
	return ArticleExistedError{}
}

func (e ArticleExistedError) WithSlug(slug string) ArticleExistedError {
	return ArticleExistedError{slug: slug}
}

func (e ArticleExistedError) Error() string {
	return fmt.Sprintf("the article with slug %s already exists", e.slug)
}

type ArticleService struct {
	articleRepository repository.ArticleRepository
}

func (s ArticleService) CreateArticle(title string, description string, body string, authorId int) (model.Article, error) {
	slug := strings.ToLower(strings.Replace(title, " ", "-", -1))
	if s.articleRepository.ExistsBySlug(slug) {
		return model.Article{}, NewArticleExistsError().WithSlug(slug)
	}
	article := model.NewArticle(slug, title, description, body, authorId)
	result, err := s.articleRepository.Save(article)
	return result, err
}

func NewArticleService(articleRepository repository.ArticleRepository) *ArticleService {
	return &ArticleService{articleRepository: articleRepository}
}
