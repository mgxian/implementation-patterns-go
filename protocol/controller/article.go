package controller

import (
	"article/domain/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type ArticleController struct {
	articleService service.ArticleServiceInterface
}

func (ac ArticleController) CreateArticle(c echo.Context) error {
	var request ArticleCreateRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	article, err := ac.articleService.CreateArticle(request.Title, request.Description, request.Body, request.AuthorId)
	if err == nil {
		return c.JSON(http.StatusCreated, service.NewArticleDTOFromArticle(article))
	}

	if _, ok := err.(service.ArticleExistedError); ok {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.String(http.StatusInternalServerError, err.Error())
}

func NewArticleController(articleService service.ArticleServiceInterface) ArticleController {
	return ArticleController{articleService: articleService}
}
