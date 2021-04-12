package pageutils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//QueryPaginationParameter 解析query 分页参数， 返回 page, pageSize
func QueryPaginationParameter(c *gin.Context) (int, int) {
	page := 0
	pageSize := 10
	pageParam := c.DefaultQuery("page", "0")
	p, err := strconv.Atoi(pageParam)
	if err == nil {
		page = p
	}

	pageSizeParam := c.DefaultQuery("pageSize", "10")
	ps, err := strconv.Atoi(pageSizeParam)
	if err == nil {
		pageSize = ps
	}
	return page, pageSize
}

//PostFormPaginationParameter 解析postForm 分页参数， 返回 page, pageSize
func PostFormPaginationParameter(c *gin.Context) (int, int) {
	page := 0
	pageSize := 10
	pageParam := c.DefaultPostForm("page", "0")
	p, err := strconv.Atoi(pageParam)
	if err == nil {
		page = p
	}

	pageSizeParam := c.DefaultPostForm("pageSize", "10")
	ps, err := strconv.Atoi(pageSizeParam)
	if err == nil {
		pageSize = ps
	}
	return page, pageSize
}
