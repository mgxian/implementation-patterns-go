package model_test

import (
	"article/domain/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Article", func() {
	It("should create an article", func() {
		slug := "fake-article"
		title := "Fake Article"
		description := "Description"
		body := "Something"
		authorId := 1

		got := model.NewArticle(slug, title, description, body, authorId)

		Expect(got.Slug()).To(Equal(slug))
	})

	It("can set createdAt", func() {
		got := model.Article{}
		now := time.Now()
		got.SetCreatedAt(now)

		Expect(got.CreatedAt()).To(Equal(now))
	})

	It("can set updatedAt", func() {
		got := model.Article{}
		now := time.Now()
		got.SetUpdatedAt(now)

		Expect(got.UpdatedAt()).To(Equal(now))
	})
})
