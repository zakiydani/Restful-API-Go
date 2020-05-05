package main

import (
	"net/http"

	"strconv"

	"github.com/zakiydani/http-service/model"

	"github.com/labstack/echo"
)

func main() {
	// 1.init data
	store := model.NewArticleStoreInMemory()

	e := echo.New()

	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap
			return c.JSON(http.StatusOK, articles)
		})

	e.GET("/articles/:id", func(c echo.Context) error {	
		articles := store.ArticleMap
			return c.JSON(http.StatusOK, articles)
		})

	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")
		article, _ := model.CreateArticle(title, body)
		store.Save(article)
			return c.JSON(http.StatusOK, article)
	})

	e.PUT("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		title := c.FormValue("title")
		body := c.FormValue("body")
		store.Update(id, title, body)
			return c.JSON(http.StatusOK, store.ArticleMap)	
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		store.Remove(id)
			return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.Logger.Fatal(e.Start(":8080"))

	// e.POST("articles", func(c echo.Context) error {
	// 	user := c.FormValue("user")
	// 	age := c.FormValue("age")
	// 	return c.String(http.StatusOK, user+","+age)
	// })

	// e.GET("articles", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk mendapatkan data list")
	// })

	// e.POST("articles", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk create article")
	// })

	// e.PUT("articles/:id", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk update article berdasarkan id")
	// })

	// e.DELETE("articles/:id", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk menghapus article berdasarkan id")
	// })

	// e.GET("articles/:id", func(c echo.Context) error{
	// 	return c.String(http.StatusOK, "untuk mendapatkan article berdasarkan id")
	// })

	// e.Logger.Fatal(e.Start(":8080"))
}