package util

import (
	"github.com/gin-gonic/gin"
)

/**
 * Created by muhammad.khadafi on 27/05/2022
 */

func APIResponse(c *gin.Context, message string, code int, status string, data interface{}) {
	c.JSON(code,
		gin.H{
			"message": message,
			"code":    code,
			"status":  status,
			"data":    data,
		})
}
