package api_test

import (
	"article"
	"article/domain/service"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ = Describe("Article", func() {
	var (
		address = "127.0.0.1:6666"
	)

	BeforeSuite(func() {
		go article.RunAt(address)
	})

	It("should create an article", func() {
		bodyJson := `
		{
			  "title": "Fake Title",
			  "description": "Description",
			  "body": "Something",
			  "author_id": 1 
		}
		`
		var err error
		var resp *http.Response
		var article service.ArticleDTO
		Eventually(func() error {
			url := fmt.Sprintf("http://%s%s", address, "/articles")
			resp, err = http.Post(url, echo.MIMEApplicationJSON, strings.NewReader(bodyJson))
			return err
		}).ShouldNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusCreated))
		content, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		err = json.Unmarshal(content, &article)
		Expect(err).NotTo(HaveOccurred())
		Expect(article.Slug).To(Equal("fake-title"))
		Expect(article.Description).To(Equal("Description"))
		Expect(article.Body).To(Equal("Something"))
		Expect(article.AuthorId).To(Equal(1))
		Expect(article.CreatedAt).NotTo(BeZero())
		Expect(article.UpdatedAt).NotTo(BeZero())
	})

	It("should return status code 409 when the article already exists", func() {
		bodyJson := `
		{
			  "title": "Fake Title",
			  "description": "Description",
			  "body": "Something",
			  "author_id": 1 
		}
		`
		var err error
		var resp *http.Response
		Eventually(func() error {
			url := fmt.Sprintf("http://%s%s", address, "/articles")
			resp, err = http.Post(url, echo.MIMEApplicationJSON, strings.NewReader(bodyJson))
			return err
		}).ShouldNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusConflict))
		content, err := ioutil.ReadAll(resp.Body)
		want := `
		{
			"message": "the article with slug fake-title already exists"
		}
		`
		Expect(content).To(MatchJSON(want))
	})
})
