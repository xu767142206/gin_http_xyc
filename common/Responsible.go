package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJson(httpCode, code int, message string, data interface{}) (int, gin.H) {

	resMap := gin.H{
		"code":    code,
		"message": message,
		"data":    data}
	return httpCode, resMap
}

func Success(message string, data interface{}) (int, gin.H) {

	resMap := gin.H{
		"code":    200,
		"message": message,
		"data":    data}
	return http.StatusOK, resMap
}

func ClientError(message string, data interface{}) (int, gin.H) {

	resMap := gin.H{
		"code":    400,
		"message": message,
		"data":    data}
	return http.StatusOK, resMap
	// return http.StatusBadRequest, resMap //400状态吗
}

func ServiceError(message string, data interface{}) (int, gin.H) {

	resMap := gin.H{
		"code":    500,
		"message": message,
		"data":    data}
	return http.StatusOK, resMap
	// return http.StatusInternalServerError, resMap //500状态吗
}
