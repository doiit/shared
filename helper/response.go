package helper

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit   int64   `json:"limit"`
	Offset  int64   `json:"offset"`
	Prev    float64 `json:"prev"`
	Next    float64 `json:"next"`
	Current float64 `json:"current"`
	Total   int     `json:"total"`
}

func GetPagination(c *gin.Context) Pagination {
	// Get the page limit from the environment variable or use the default value of 10
	PAGE_LIMIT := GetEnv("PAGE_LIMIT", "10")
	limit, err := strconv.ParseInt(PAGE_LIMIT, 10, 64)

	if err != nil {
		limit = 10
	}

	offset := int64(0)

	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	if o := c.Query("offset"); o != "" {
		fmt.Sscanf(o, "%d", &offset)
	}

	return Pagination{
		Limit:   limit,
		Offset:  offset,
		Prev:    math.Max(0, float64(offset-limit)),
		Next:    math.Max(0, float64(offset+limit)),
		Current: math.Floor(float64(offset)/float64(limit)) + 1,
	}
}

func SendResponse(c *gin.Context, statusCode int, errorCode string, message string, success bool, data any) {
	defer c.Abort()
	c.JSON(statusCode, gin.H{
		"error":   errorCode,
		"message": message,
		"success": success,
		"data":    data,
	})
}

func SendSuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	SendResponse(c, statusCode, "", message, true, data)
}

func SendErrorResponse(c *gin.Context, statusCode int, errorCode, message string) {
	SendResponse(c, statusCode, errorCode, message, false, nil)
}
