package service_test

import (
	"article/domain/model"
	"article/domain/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"article/domain/service"
)

var _ = Describe("Article", func() {
	It("should create an article", func() {
		articleRepository := repository.NewArticleRepositoryStub()
		articleService := service.NewArticleService(&articleRepository)
		article := model.NewArticle("fake-article", "Fake Article", "Description", "Something", 1)
		got, err := articleService.CreateArticle("Fake Article", "Description", "Something", 1)

		Expect(err).NotTo(HaveOccurred())
		Expect(got.Slug()).To(Equal(article.Slug()))
		Expect(got.Title()).To(Equal(article.Title()))
		Expect(got.Description()).To(Equal(article.Description()))
		Expect(got.Body()).To(Equal(article.Body()))
		Expect(got.CreatedAt()).NotTo(BeZero())
		Expect(got.UpdatedAt()).NotTo(BeZero())
	})
})
