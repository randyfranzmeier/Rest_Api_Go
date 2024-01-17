package middleware

import (
	"Rest_Api_Go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "not authorized!!!"})
		return
	}
	//abort with status json immediately ends the request
	err, userID := utils.ValidToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg:": "not authorized!"})
		return
	}

	context.Set("userID", userID)
	context.Next() //this enables the proceeding lines of code to be handled
}
