package helpers

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

const paginationLimit int = 8

func GeneratePaginationFromRequest(c *gin.Context, count int) models.Pagination {
	// Initializing default
	//	var mode string

	var page int
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

	}
	pageCount := int(math.Ceil(float64(count) / float64(paginationLimit)))
	if pageCount == 0 {
		pageCount = 1
	}
	var prevPage, nextPage int
	if page > 1 {
		prevPage = page - 1
	}
	if page < pageCount {
		nextPage = page + 1
	}
	return models.Pagination{
		Limit:    paginationLimit,
		Page:     page,
		Previous: prevPage,
		Next:     nextPage,
	}

}
