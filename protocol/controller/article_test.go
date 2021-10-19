package controller_test

import (
	"article/domain/model"
	"article/domain/service"
	"article/domain/service/mock"
	"article/protocol/controller"
	clock "article/provider/clock/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var _ = Describe("Article", func() {
	var (
		mockCtrl          *gomock.Controller
		articleService    *mock.MockArticleServiceInterface
		articleController controller.ArticleController
		mockClock         *clock.MockClock
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		articleService = mock.NewMockArticleServiceInterface(mockCtrl)
		articleController = controller.NewArticleController(articleService)
		mockClock = clock.NewMockClock(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should create an article", func() {
		bodyJson := `
		{
			"title": "Fake Article",
			"description": "Description",
			"body": "Something",
			"author_id": 1
		}
		`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(bodyJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		now, err := time.Parse("2006-01-02 15:04:05", "2021-10-18 18:42:18")
		Expect(err).NotTo(HaveOccurred())
		mockClock.EXPECT().Now().AnyTimes().Return(now)

		articleService.EXPECT().
			CreateArticle(
				gomock.Eq("Fake Article"),
				gomock.Eq("Description"),
				gomock.Eq("Something"),
				gomock.Eq(1)).
			DoAndReturn(func(_, _, _ string, _ int) (model.Article, error) {
				article := model.NewArticle(
					"fake-article",
					"Fake Article",
					"Description",
					"Something",
					1)
				article.SetCreatedAt(mockClock.Now())
				article.SetUpdatedAt(mockClock.Now())
				return article, nil
			})

		_ = articleController.CreateArticle(c)

		want := `
		{
			 "slug": "fake-article",
			 "title": "Fake Article",
			 "description": "Description",
			 "body": "Something",
			 "author_id": 1,
			 "created_at": "2021-10-18T18:42:18Z",
			 "updated_at": "2021-10-18T18:42:18Z"
		}
		`

		Expect(rec.Code).To(Equal(http.StatusCreated))
		Expect(rec.Body.String()).To(MatchJSON(want))
	})

	It("should return 409 status code when article already exists", func() {
		bodyJson := `
		{
			"title": "Fake Article",
			"description": "Description",
			"body": "Something",
			"author_id": 1
		}
		`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(bodyJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		articleService.EXPECT().
			CreateArticle(gomock.Eq("Fake Article"), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(model.Article{}, service.NewArticleExistsError().WithSlug("fake-article"))

		_ = articleController.CreateArticle(c)

		want := `
		{
			 "message": "the article with slug fake-article already exists"
		}
		`

		Expect(rec.Code).To(Equal(http.StatusConflict))
		Expect(rec.Body.String()).To(MatchJSON(want))
	})
})
