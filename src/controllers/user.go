package controllers

import (
	"github.com/MatheusMeloAntiquera/api-go/src/models"
	"github.com/gin-gonic/gin"
)

var users []models.User
var user models.User

func UserIndex(context *gin.Context) {

	Db.Find(&users)
	context.JSON(200, gin.H{
		"success": true,
		"data":    users,
	})
}

func UserShow(context *gin.Context) {
	Db.First(&user, context.Param("id"))
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func UserCreate(context *gin.Context) {
	user := models.User{Name: context.PostForm("name")}
	Db.Create(&user)
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func UserUpdate(context *gin.Context) {
	Db.First(&user, context.Param("id"))
	user.Name = context.PostForm("name")
	Db.Save(&user)
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func UserDelete(context *gin.Context) {
	Db.First(&user, context.Param("id"))
	Db.Delete(&user)
	context.JSON(200, gin.H{
		"success": true,
	})
}
