package handler

import (
	"net/http"

	"github.com/arifseft/go-gin-crud/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
