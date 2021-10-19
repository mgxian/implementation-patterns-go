package model

import (
	"time"
)

type Article struct {
	slug        string
	title       string
	description string
	body        string
	authorId    int
	createdAt   time.Time
	updatedAt   time.Time
}

func (a Article) Slug() string {
	return a.slug
}

func (a Article) CreatedAt() time.Time {
	return a.createdAt
}

func (a Article) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Article) SetUpdatedAt(now time.Time) {
	a.updatedAt = now
}

func (a *Article) SetCreatedAt(now time.Time) {
	a.createdAt = now
}

func (a Article) Title() string {
	return a.title
}

func (a Article) Description() string {
	return a.description
}

func (a Article) Body() string {
	return a.body
}

func (a Article) AuthorId() int {
	return a.authorId
}

func NewArticle(slug string, title string, description string, body string, authorId int) Article {
	return Article{
		slug:        slug,
		title:       title,
		description: description,
		body:        body,
		authorId:    authorId,
	}
}
