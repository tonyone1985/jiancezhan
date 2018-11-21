// webbase project main.go
package main

import (
	"fmt"

	//"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//	router.Static("/css", "views/css")
	//	router.Static("/images", "views/images")
	//	router.Static("/js", "views/js")
	//	router.Static("/editor", "views/editor")
	//	router.Static("/src", "views/src")

	router.Static("/", "views")

	//	router.GET("/.html", func(c *gin.Context) {
	//		c.HTML(http.StatusOK, "", gin.H{
	//			"title": "Posts",
	//		})
	//	})
	/*
		router.LoadHTMLGlob("views/pages/*")

		router.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Posts",
			})
		})
		router.GET("/form", func(c *gin.Context) {
			c.HTML(http.StatusOK, "form.html", gin.H{
				"title": "Users",
			})
		})
		router.GET("/main.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "main.html", gin.H{
				"title": "Users",
			})
		})
		router.GET("/main.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "main.html", gin.H{
				"title": "Users",
			})
		})
	*/
	router.Run(":38080")

	fmt.Println("Hello World!")
}
