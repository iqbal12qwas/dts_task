package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShowAll(ctx *gin.Context) {
	data := gin.H{
		"title":  "Index Page",
		"Title": "Test",
		"Description": "Test",
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
