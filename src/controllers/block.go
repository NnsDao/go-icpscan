package controllers

import (
	"bytes"
	"fmt"
	"github.com/MatheusMeloAntiquera/api-go/src/models"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

var blocks []models.Block
var block models.Block


var details []models.Detail
var detail models.Detail


type AutoGenerated struct {
	Operations []struct {
		Account struct {
			Address string `json:"address"`
		} `json:"account"`
		Amount struct {
			Currency struct {
				Decimals int `json:"decimals"`
				Symbol string `json:"symbol"`
			} `json:"currency"`
			Value string `json:"value"`
		} `json:"amount"`
		OperationIdentifier struct {
			Index int `json:"index"`
		} `json:"operation_identifier"`
		Status string `json:"status"`
		Type string `json:"type"`
	} `json:"operations"`
}

func BlockIndex(context *gin.Context) {



	jsonStr :=[]byte(`{
			"network_identifier": {
				"blockchain": "Internet Computer",
				"network": "00000000000000020101"
			},
			"block_identifier": {
				"index": 86135
			}
		}`)


	req, err := http.NewRequest("POST", "https://rosetta-api.internetcomputer.org/block", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("authority", "rosetta-api.internetcomputer.org")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("sec-ch-ua", "' Not A;Brand';v='99', 'Chromium';v='90', 'Google Chrome';v='90'")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("origin", "https://dashboard.internetcomputer.org")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://dashboard.internetcomputer.org/")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")


	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	js, err2 := simplejson.NewJson(str)

	if err2 != nil {
		fmt.Println("抓取fail" ,err2 )
	}
	// 保存数据

	// 当前区块高度
	var bindex = js.Get("block").Get("block_identifier").Get("index").MustInt()
	// 上一个区块
	var preindex = js.Get("block").Get("parent_block_identifier").Get("index").MustInt()

	// 当前hash
	var cindex = js.Get("block").Get("parent_block_identifier").Get("hash").MustString()

	// 转账凭证hash
	var tidentify = js.Get("block").Get("transactions").GetIndex(2).Get("transaction_identifier").Get("hash").MustString()

	// 转账时间
	var bttime = js.Get("block").Get("timestamp").MustInt()

	// 二层数据

	// 转账区块高度
	var bheight = js.Get("block").Get("transactions").GetIndex(0).Get("metadata").Get("block_height").MustInt()

	// 转账备注
	var mmemo = js.Get("block").Get("transactions").GetIndex(0).Get("metadata").Get("memo").MustUint64()

	// 转账区块时间 utc
	var utime = js.Get("block").Get("transactions").GetIndex(0).Get("metadata").Get("timestamp").MustInt()


	fmt.Println(strconv.Itoa(bindex))
	fmt.Println(preindex)
	fmt.Println(cindex)
	fmt.Println(bttime)
	fmt.Println(tidentify)

	fmt.Println()

	fmt.Println("metada1:",bheight)
	fmt.Println("metada2:",mmemo)
	fmt.Println("metada3:",utime)

	// 存储多层信息


	// 转账数据
	var tData = js.Get("block").Get("transactions").GetIndex(0).Get("operations").MustArray()

	fmt.Println(tData)

	// 存储外层信息
	//block := models.Block{
	//	Blockidentifier: strconv.Itoa(bindex),
	//	Parentblock:  strconv.Itoa(preindex) ,
	//	Transactiohash: cindex,
	//	Mblockheight: strconv.Itoa(bheight),
	//	Mmemo: strconv.FormatUint(mmemo, 10),
	//	Tranidentifier: tidentify,
	//	Mtimestamp: strconv.Itoa(utime),
	//	Blocktimestamp: strconv.Itoa(bttime),
	//
	//}
	//
	//db.Create(&block)
	//context.JSON(200, gin.H{
	//	"success": true,
	//	"data":    block,
	//})


	//db.Find(&blocks)
	context.JSON(200, gin.H{
		"success": true,
		"data":    js,
	})
}

func BlockShow(context *gin.Context) {
	db.First(&user, context.Param("id"))
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

// 爬取数据

func BlockCreate(context *gin.Context) {

	user := models.User{Name: context.PostForm("name")}
	db.Create(&user)
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func BlockUpdate(context *gin.Context) {
	db.First(&user, context.Param("id"))
	user.Name = context.PostForm("name")
	db.Save(&user)
	context.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func BlockDelete(context *gin.Context) {
	db.First(&user, context.Param("id"))
	db.Delete(&user)
	context.JSON(200, gin.H{
		"success": true,
	})
}
