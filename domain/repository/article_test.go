package repository_test

import (
	"article/domain/model"
	"article/domain/repository"
	"article/provider/clock/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Article", func() {
	var (
		mockCtrl          *gomock.Controller
		clock             *mock.MockClock
		articleRepository *repository.ArticleRepositoryMemory
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		clock = mock.NewMockClock(mockCtrl)
		articleRepository = repository.NewArticlesRepositoryMemory(clock)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should save an article", func() {
		now := time.Now()
		clock.EXPECT().Now().Return(now)
		article := model.NewArticle("fake-article", "Fake Article", "Description", "Something", 1)

		got, err := articleRepository.Save(article)

		Expect(err).NotTo(HaveOccurred())
		Expect(got.Slug()).To(Equal("fake-article"))
		Expect(got.Title()).To(Equal("Fake Article"))
		Expect(got.Description()).To(Equal("Description"))
		Expect(got.Body()).To(Equal("Something"))
		Expect(got.CreatedAt()).To(Equal(now))
		Expect(got.UpdatedAt()).To(Equal(now))
	})
})
