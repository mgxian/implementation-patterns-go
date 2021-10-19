package repository_test

import (
	"article/domain/model"
	"article/domain/repository"
	"article/provider/clock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Article", func() {
	It("should save an article", func() {
		clockStub := clock.NewStub()
		articleRepository := repository.NewArticlesRepositoryMemory(clockStub)
		article := model.NewArticle("fake-article", "Fake Article", "Description", "Something", 1)
		article.SetCreatedAt(clockStub.Now())
		got, err := articleRepository.Save(article)
		Expect(err).NotTo(HaveOccurred())
		Expect(got.Slug()).To(Equal(article.Slug()))
		Expect(got.Title()).To(Equal(article.Title()))
		Expect(got.Description()).To(Equal(article.Description()))
		Expect(got.Body()).To(Equal(article.Body()))
		Expect(got.CreatedAt()).To(Equal(clockStub.Now()))
		Expect(got.UpdatedAt()).To(Equal(clockStub.Now()))
	})
})
