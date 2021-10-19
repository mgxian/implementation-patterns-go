package article

import (
	"article/domain/repository"
	"article/domain/service"
	"article/protocol/controller"
	"article/provider/clock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunAt(address string) {
	e := echo.New()
	e.Use(middleware.Recover())

	systemClock := clock.NewSystem()
	articleRepository := repository.NewArticlesRepositoryMemory(systemClock)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	e.POST("/articles", articleController.CreateArticle)

	e.Logger.Fatal(e.Start(address))
}

func RunWithLogAt(address string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	systemClock := clock.NewSystem()
	articleRepository := repository.NewArticlesRepositoryMemory(systemClock)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	e.POST("/articles", articleController.CreateArticle)

	e.Logger.Fatal(e.Start(address))
}
