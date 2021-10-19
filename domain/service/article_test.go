package service_test

import (
	"article/domain/model"
	"article/domain/repository/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"article/domain/service"
)

var _ = Describe("Article", func() {
	var (
		mockCtrl          *gomock.Controller
		articleRepository *mock.MockArticleRepository
		articleService    *service.ArticleService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		articleRepository = mock.NewMockArticleRepository(mockCtrl)
		articleService = service.NewArticleService(articleRepository)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should create an article", func() {
		article := model.NewArticle("fake-article", "Fake Article", "Description", "Something", 1)
		articleRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(article model.Article) (model.Article, error) {
			now := time.Now()
			article.SetCreatedAt(now)
			article.SetUpdatedAt(now)
			return article, nil
		})

		got, err := articleService.CreateArticle("Fake Article", "Description", "Something", 1)

		Expect(err).NotTo(HaveOccurred())
		Expect(got.Slug()).To(Equal(article.Slug()))
		Expect(got.Title()).To(Equal(article.Title()))
		Expect(got.Description()).To(Equal(article.Description()))
		Expect(got.Body()).To(Equal(article.Body()))
		Expect(got.CreatedAt()).NotTo(BeZero())
		Expect(got.UpdatedAt()).NotTo(BeZero())
	})

	It("can convert from article to ArticleDTO", func() {
		article := model.NewArticle("fake-article", "Fake Article", "Description", "Something", 1)
		article.SetCreatedAt(time.Now())
		article.SetUpdatedAt(time.Now())
		articleDTO := service.NewArticleDTOFromArticle(article)

		Expect(articleDTO.Slug).To(Equal(article.Slug()))
		Expect(articleDTO.Title).To(Equal(article.Title()))
		Expect(articleDTO.Description).To(Equal(article.Description()))
		Expect(articleDTO.Body).To(Equal(article.Body()))
		Expect(articleDTO.CreatedAt).To(Equal(article.CreatedAt()))
		Expect(articleDTO.UpdatedAt).To(Equal(article.UpdatedAt()))
	})
})
