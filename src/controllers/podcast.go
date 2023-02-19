package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetPodcastInfoRes struct {
	Canister string `json:"canister" form:"canister" binding:"required"`
	Id       string `json:"id" form:"id" binding:"required"`
}

type GetPodcastInfoResp struct {
	Count int64 `json:"count"`
}

type RecordPlayRes struct {
	Canister string `json:"canister" form:"canister" binding:"required"`
	Id       string `json:"id" form:"id" binding:"required"`
}

func GetPodcastInfo(gx *gin.Context) {
	res := GetPodcastInfoRes{}
	err := gx.ShouldBindQuery(&res)
	if err != nil {
		fmt.Printf("err is %v", res)
		gx.JSON(500, gin.H{
			"success": false,
			"message": "params error",
		})
		return
	}

	var count int64

	err = Db.Table("podcast_record_play").Where("canister", res.Canister).Where("podcast_id", res.Id).Count(&count).Error

	if err != nil {
		fmt.Printf("err is %v", res)
		gx.JSON(500, gin.H{
			"success": false,
			"message": "service error",
		})
		return
	}

	gx.JSON(200, gin.H{
		"success": true,
		"message": "",
		"data": GetPodcastInfoResp{
			Count: count,
		},
	})
	return
}

func RecordPlay(gx *gin.Context) {
	res := RecordPlayRes{}
	err := gx.ShouldBindJSON(&res)
	if err != nil {
		fmt.Printf("err is %v", res)
		gx.JSON(500, gin.H{
			"success": false,
			"message": "params error",
		})
		return
	}

	fmt.Printf("res is %v", res)

	err = Db.Table("podcast_record_play").Create(map[string]interface{}{
		"canister":   res.Canister,
		"podcast_id": res.Id,
	}).Error

	if err != nil {
		fmt.Printf("err is %v", res)
		gx.JSON(500, gin.H{
			"success": false,
			"message": "service error",
		})
		return
	}

	gx.JSON(200, gin.H{
		"success": true,
		"message": "",
		"data":    "",
	})
	return
}
