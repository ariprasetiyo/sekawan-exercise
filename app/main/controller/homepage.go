package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/indexx.html",
		gin.H{
			"title":      "Geeksbeginner",
			"textRender": "Hallo ari",
		},
	)
}

func About(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/about.html",
		gin.H{},
	)
}

func Contact(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/contact.html",
		gin.H{
			"title": "Contact",
		},
	)
}
