package routes

import (
	"Rest_Api_Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(context *gin.Context) {
	//extract data
	var newUser models.User
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Unable to parse data!"})
		return
	}
	//call save method to save it
	//newUser.ID = 1
	err = newUser.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user!"})
		return
	}
	//return status
	context.JSON(http.StatusCreated, gin.H{"message": "New user created!"})
}

func UserLogin(context *gin.Context) {
	var user models.User //user struct for extracting data

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Unable to parse data!"})
		return
	}

	err = user.ValidateLogin()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"msg:": "failed to authenticate"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"loginStatus:": "success"})
}
