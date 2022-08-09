package utils

import (
	"gin-api/helper"
	// "strconv"
	// "fmt"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Search  string `json:"search"`
}

//GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c *gin.Context) (helper.Pagination) {
	// Initializing default
	var input helper.Pagination
	limit := 10
	page := 1
	sort := "id asc"
	search := ""
	c.ShouldBindJSON(&input)
	if input.Limit != 0 {
		return helper.Pagination{
			Limit: input.Limit,
			Page:  input.Page,
			Sort:  input.Sort,
			Search:  input.Search,
		}
	}
	
	return helper.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
		Search: search,
	}

}