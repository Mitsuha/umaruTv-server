package gin

import "github.com/gin-gonic/gin"

type commonResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, code int, obj interface{}) {
	response := commonResponse{
		Code:   code,
		Status: code < 300,
	}

	if _, ok := obj.(string); ok {
		response.Message = obj.(string)
	} else {
		response.Data = obj
	}

	c.JSON(code, response)
}
