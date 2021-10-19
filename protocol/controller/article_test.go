package controller_test

import (
	"article/domain/model"
	"article/domain/service"
	"article/protocol/controller"
	"article/provider/clock"
	"errors"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var _ = Describe("Article", func() {
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

		articleService := service.NewArticleServiceStub()
		clock := clock.NewStub()
		now, _ := time.Parse("2006-01-02 15:04:05", "2021-10-18 18:42:18")
		clock.SetNow(now)
		articleService.CreateArticleFn = func(title, description, body string, authorId int) (model.Article, error) {
			if title == "Fake Article" && description == "Description" && body == "Something" && authorId == 1 {
				article := model.NewArticle("fake-article", "Fake Article", "Description", "Something", 1)
				article.SetCreatedAt(clock.Now())
				article.SetUpdatedAt(clock.Now())
				return article, nil
			}
			return model.Article{}, errors.New("create article error")
		}
		articleController := controller.NewArticleController(articleService)

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
})
