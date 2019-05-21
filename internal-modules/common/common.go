package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeqown/infrastructure/types/codes"
)

const (
	// CreateTimeDesc ...
	CreateTimeDesc = "-createtime"
	// CreateTimeAsc ...
	CreateTimeAsc = "createtime"
	// UpdateTimeDesc ...
	UpdateTimeDesc = "-updatetime"
	// UpdateTimeAsc ...
	UpdateTimeAsc = "updatetime"
)

// ResponseError ...
func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code": codes.CodeSystemErr,
		"err":  err.Error(),
	})
}

// Response ...
func Response(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}
