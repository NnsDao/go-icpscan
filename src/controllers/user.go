package controllers

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"icpscan/src/models"
	"icpscan/src/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var users []models.Users
// var user models.Users

// func UserIndex(context *gin.Context) {

// 	Db.Find(&users)
// 	context.JSON(200, gin.H{
// 		"success": true,
// 		"data":    users,
// 	})
// }

// func UserShow(context *gin.Context) {
// 	Db.First(&user, context.Param("id"))
// 	context.JSON(200, gin.H{
// 		"success": true,
// 		"data":    user,
// 	})
// }

// func UserCreate(context *gin.Context) {
// 	user := models.Users{Name: context.PostForm("name")}
// 	Db.Create(&user)
// 	context.JSON(200, gin.H{
// 		"success": true,
// 		"data":    user,
// 	})
// }

// func UserUpdate(context *gin.Context) {
// 	Db.First(&user, context.Param("id"))
// 	user.Name = context.PostForm("name")
// 	Db.Save(&user)
// 	context.JSON(200, gin.H{
// 		"success": true,
// 		"data":    user,
// 	})
// }

// func UserDelete(context *gin.Context) {
// 	Db.First(&user, context.Param("id"))
// 	Db.Delete(&user)
// 	context.JSON(200, gin.H{
// 		"success": true,
// 	})
// }

// @Summary 用户登录、注册
// @Description 用户登录、注册
// @Param  principal_id query string true "principal_id"
// @Success 200 {object} response.JSONResult{data=response.LoginRes}
// @Router /api/user/login [get]
func Login(c *gin.Context) {

	// 判断是否为ajax请求
	// requestType := c.GetHeader("X-Requested-With")
	// if requestType != "XMLHttpRequest" {
	// 	return
	// }

	principalId := c.Query("principal_id")
	if principalId == "" || strings.Count(principalId, "") == 58 {
		c.JSON(500, gin.H{
			"success": false,
			"data":    "",
			"message": "principalId参数不存在",
		})
		return
	}
	var user *models.Users
	var res response.LoginRes
	if err := Db.Table("users").Select("name").Where("principal_id = ?", principalId).First(&user).Error; err == gorm.ErrRecordNotFound {
		var Aprivate [32]byte
		//产生随机数
		if _, err := io.ReadFull(rand.Reader, Aprivate[:]); err != nil {
			os.Exit(0)
		}
		name := fmt.Sprintf("%x", Aprivate)
		res = response.LoginRes{
			Name: name,
		}

		Db.Table("users").Create(&models.Users{
			Name:        name,
			PrincipalId: principalId,
		})

		c.JSON(200, gin.H{
			"success": true,
			"data":    res,
			"message": "",
		})
		return
	}

	res = response.LoginRes{
		Name: user.Name,
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    res,
		"message": "",
	})
}

func WalletRelation(c *gin.Context) {
	res := response.WalletRelationRes{}
	err := c.Bind(&res)
	if err != nil {
		c.JSON(401, gin.H{
			"success": false,
			"data":    "",
			"message": "arg err",
		})
		return
	}

	fmt.Printf("%+v", res)

	if err := Db.Where("principal = ?", res.Principal).First(&models.WalletRelation{}).Error; err == gorm.ErrRecordNotFound {

		fmt.Println(err)

		cmdStr := "./transition_account " + res.Principal
		cmd := exec.Command("bash", "-c", cmdStr)

		out, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(401, gin.H{
				"success": false,
				"data":    "",
				"message": "arg err",
			})
			return
		}

		addr := fmt.Sprintf("%s", out)

		fmt.Println(addr)
		fmt.Println(res.Address)

		if addr != res.Address {
			c.JSON(401, gin.H{
				"success": false,
				"data":    "",
				"message": "arg err",
			})
			return
		}

		walletRelation := models.WalletRelation{
			Principal: res.Principal,
			Address:   res.Address,
		}

		Db.Create(&walletRelation)
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "",
	})

}
