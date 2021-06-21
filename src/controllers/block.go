package controllers

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/MatheusMeloAntiquera/api-go/src/models"
	"github.com/MatheusMeloAntiquera/api-go/src/response"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"net/http"
	"strconv"
)

var detail models.Detail


type Identifier struct {
	NetworkIdentifier Network `json:"network_identifier"`
	BlockIdentifier Block `json:"block_identifier"`
}

type Network struct {
	BlockChain string `json:"blockchain"`
	Network string `json:"network"`
}

type Block struct {
	Index int32 `json:"index"`
}


func (obj *Identifier) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, obj); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


func BlockIndex(context *gin.Context) {


	//备注 结束循环 TODO 定时
	//for i := 192897; i < 193308; i++ {
	//	当前的参数

	// 查找当前的blockheight

	var blockId []models.Block
	Db.Table("blocks").Select("mblockheight").Order("mblockheight desc").Limit(1).Scan(&blockId)

	fmt.Println("pppp", blockId[0].Mblockheight)

	ih, err := strconv.ParseInt(blockId[0].Mblockheight, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	j := int32(ih)+1
	fmt.Printf("j value is %d, type is %T", j, j)
	//return
	identify :=  Identifier {}

	identify.NetworkIdentifier.BlockChain = "Internet Computer"
	identify.NetworkIdentifier.Network = "00000000000000020101"
	identify.BlockIdentifier.Index = j


	jsonStu, err := json.Marshal(identify)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	//fmt.Println(string(jsonStu))


	req, err := http.NewRequest("POST", "https://rosetta-api.internetcomputer.org/block", bytes.NewBuffer(jsonStu))
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
		return
	}
	// 保存数据

	// 当前区块高度
	var bindex = js.Get("block").Get("block_identifier").Get("index").MustInt()
	// 上一个区块
	var preindex = js.Get("block").Get("parent_block_identifier").Get("index").MustInt()

	// 当前hash
	var cindex = js.Get("block").Get("parent_block_identifier").Get("hash").MustString()

	// 转账凭证hash
	var tidentify = js.Get("block").Get("transactions").GetIndex(0).Get("transaction_identifier").Get("hash").MustString()

	// 转账时间
	var bttime = js.Get("block").Get("timestamp").MustInt()

	// 二层数据

	// 转账区块高度
	var bheight = js.Get("block").Get("transactions").GetIndex(0).Get("metadata").Get("block_height").MustInt()

	// 转账备注
	var mmemo = js.Get("block").Get("transactions").GetIndex(0).Get("metadata").Get("memo").MustUint64()

	// 转账区块时间 utc
	var utime = js.Get("block").Get("transactions").GetIndex(0).Get("metadata").Get("timestamp").MustInt()


	//fmt.Println(strconv.Itoa(bindex))
	//fmt.Println(preindex)
	//fmt.Println(cindex)
	//fmt.Println(bttime)
	//fmt.Println("转账hash",tidentify)
	//
	//fmt.Println()
	//
	//fmt.Println("metada1:",bheight)
	//fmt.Println("metada2:",mmemo)
	//fmt.Println("metada3:",utime)
	// 存储多层信息

	// 转账数据
	var tData = js.Get("block").Get("transactions").GetIndex(0).Get("operations").MustArray()

	//fmt.Println(tData)


	for _, row := range tData {
		if each_map, ok := row.(map[string]interface{}); ok {

			// 用户地址
			var oaddress string
			if start_ts, ok := each_map["account"]; ok {


				if plll, ok := start_ts.(map[string]interface{}); ok {
					oaddress = plll["address"].(string)


				}

			}
			var oindex string
			if start_ts, ok := each_map["operation_identifier"]; ok {


				if plll, ok := start_ts.(map[string]interface{}); ok {
					oindex = string(plll["index"].(json.Number))
				}

			}


			var ovalue string
			var odecimals string
			var osymbol string
			if start_ts, ok := each_map["amount"]; ok {

				// 用户地址
				if plll, ok := start_ts.(map[string]interface{}); ok {

					// 转账的金额
					ovalue = string(plll["value"].(string))

					if op,ok := plll["currency"].(map[string]interface{}); ok{
						odecimals = string(op["decimals"].(json.Number))
						osymbol = string(op["symbol"].(string))
					}
				}
			}


			var otype string
			var ostatus string

			otype = each_map["type"].(string)
			ostatus = each_map["status"].(string)

			//fmt.Println("当前用户地址",oaddress)
			//fmt.Println(ovalue)
			//fmt.Println(otype)
			//fmt.Println(odecimals)
			//fmt.Println(osymbol)
			//fmt.Println(ostatus)
			//fmt.Println("当前索引",oindex)

			detail := models.Detail{
				Tranidentifier: tidentify,
				Oindex: oindex,
				Otype:  otype ,
				Ostatus: ostatus,
				Oaccountaddress: oaddress,
				Oamountvalue: ovalue,
				Oamountcurrencydecimals: odecimals,
				Oamountcurrencysymbol: osymbol,
				Mblockheight: strconv.Itoa(bheight),
				Mmemo: strconv.FormatUint(mmemo, 10),
				Mtimestamp: strconv.Itoa(utime),

			}

			Db.Create(&detail)
			context.JSON(200, gin.H{
				"success": true,
				"data":    detail,
			})


		}
	}



	// 存储外层信息
	block := models.Block{
		Blockidentifier: strconv.Itoa(bindex),
		Parentblock:  strconv.Itoa(preindex) ,
		Transactiohash: cindex,
		Mblockheight: strconv.Itoa(bheight),
		Mmemo: strconv.FormatUint(mmemo, 10),
		Tranidentifier: tidentify,
		Mtimestamp: strconv.Itoa(utime),
		Blocktimestamp: strconv.Itoa(bttime),
	}

	Db.Create(&block)
	context.JSON(200, gin.H{
		"success": true,
		"data":    block,
	})


	//Db.Find(&blocks)
	context.JSON(200, gin.H{
		"success": true,
		"data":    js,
	})
	//}

}
// TODO 排行 TOP100
func BlockShow(context *gin.Context) {

	//SELECT MAX(mtimestamp) as mtime, oaccountaddress, sum(oamountvalue) as total, count(*) as times  FROM `details` WHERE otype  <> 'FEE' GROUP BY `oaccountaddress` ORDER BY total desc LIMIT 5
	var blockShow []models.BlockShow
	Db.Table("details").Select("MAX(mtimestamp) as mtime, oaccountaddress, sum(oamountvalue) as total, count(*) as times ").Where("otype  <> ? ",
		"FEE").Group("oaccountaddress").Order("total desc").Limit(100).Scan(&blockShow)

	context.JSON(200, gin.H{
		"success": true,
		"data":    blockShow,
	})
}



// test
func BlockShowPpp(context *gin.Context) {

	var result []response.BlockNewRes

	Db.Table("details").Select("blocks.transactiohash,blocks.parentblock,blocks.mmemo,details.oaccountaddress,blocks.mblockheight, details.tranidentifier,details.oamountvalue,blocks.blocktimestamp,blocks.id").Joins("left join blocks on  blocks.mblockheight = details.mblockheight ").Limit(10).Scan(&result)

	context.JSON(200, gin.H{
		"success": true,
		"data":    result,
	})
}
//

// TODO 最新区块数据

func BlockNew(context *gin.Context) {

	var Res []response.BlockNewRes
	var block []models.Block
	var detail []models.NewDetail

	Db.Table("blocks").Select("mblockheight, mmemo, transactiohash, blocktimestamp").Order("mblockheight desc").Limit(10).Scan(&block)
	mapBlock := funk.ToMap(block, "Mblockheight").(map[string]models.Block)
	mblockheight := funk.Get(block, "Mblockheight")
	fmt.Printf("blocks is %+v\n", block)
	fmt.Printf("mblockheight is %+v\n", mblockheight)

	Db.Table("details").Select("mblockheight, ostatus, tranidentifier, oaccountaddress, SUM(oamountvalue) as osum").Where("mblockheight In ? AND otype <> ? AND oamountvalue > ?",
		mblockheight, "FEE", 0).Group("oaccountaddress, mblockheight").Order("mblockheight desc").Scan(&detail)

	// 组装返回数据
	for _,v := range detail {
		Res = append(Res, response.BlockNewRes{
			Oaccountaddress: v.Oaccountaddress,
			Tranidentifier: v.Tranidentifier,
			Blocktimestamp: mapBlock[v.Mblockheight].Blocktimestamp,
			Transactiohash: mapBlock[v.Mblockheight].Transactiohash,
			Ostatus: v.Ostatus,
			Osum: v.Osum,
			Mmemo: mapBlock[v.Mblockheight].Mmemo,
			Mblockheight: v.Mblockheight,

		})
	}

	context.JSON(200, gin.H{
		"success": true,
		"data":    Res,
	})

}

// TODO 区块高度、价格、罐的注册数量、消息等数据

//func BlockUpdate(context *gin.Context) {
//	Db.First(&user, context.Param("id"))
//	user.Name = context.PostForm("name")
//	Db.Save(&user)
//	context.JSON(200, gin.H{
//		"success": true,
//		"data":    user,
//	})
//}

func BlockDetail(c *gin.Context) {

	bid := c.Query("id") //

	fmt.Printf("id: %s;", bid)

	Db.Where("Mblockheight = ?", bid).Find(&detail)
	c.JSON(200, gin.H{
		"success": true,
		"data":    detail,
	})
}

// 查询交易、账户信息
func SearchDetail(c *gin.Context) {
	recorde_addr := c.DefaultQuery("recorde_addr", "")
	if (recorde_addr == "") {
		c.JSON(500, gin.H{
			"success": false,
			"data":    "",
			"message": "参数错误",
		})
		return
	}

	// todo 知道怎么区分account、TransactionHash是进行更改，先无脑查询
	var detail []models.Detail
	var recordCount int64
	var res response.SearchDetailRes
	Db.Table("details").Where("tranidentifier = ?", recorde_addr).Or("oaccountaddress = ?", recorde_addr).Count(&recordCount);

	// 如果等于3证明是TransactionHash,则需要进行交易详情返回
	if (recordCount == 3) {
		res.Type = 1
		Db.Table("details").Where("tranidentifier = ?", recorde_addr).Or("oaccountaddress = ?", recorde_addr).Scan(&detail);
		for _, v := range detail {
			switch v.Oindex {
			case "0":
				res.From = v.Oaccountaddress
				res.Status = v.Ostatus
			case "1":
				res.To = v.Oaccountaddress
				res.Amount = v.Oamountvalue
				res.BlockHeight = v.Mblockheight
				res.Tranidentifier = v.Tranidentifier
				res.Symbol = v.Oamountcurrencysymbol
				res.Timestamp = v.Mtimestamp
				res.Decimals = v.Oamountcurrencydecimals
				res.Memo = v.Mmemo
			case "2":
				res.Fee = v.Oamountvalue
			}
		}
	} else {
		// 此逻辑为账号信息
		var accountDetail models.AccountDetail
		Db.Table("details").Select("SUM(oamountvalue) AS balance").Where("oaccountaddress = ?", recorde_addr).Scan(&accountDetail)
		res.Type = 2
		res.Account = recorde_addr
		res.Balance = accountDetail.Balance
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    res,
	})
}

// 账户交易详情
func GetAccountDealDetail(c *gin.Context) {
	account := c.DefaultQuery("account", "")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	if (account == "") {
		c.JSON(500, gin.H{
			"success": false,
			"data":    "",
			"message": "参数错误",
		})
		return
	}

	pageInt,_ := strconv.Atoi(page)
	pageSizeInt,_ := strconv.Atoi(pageSize)

	if (pageSizeInt > 100) {
		c.JSON(500, gin.H{
			"success": false,
			"data":    "",
			"message": "分页大小最大为100",
		})
		return
	}

	var detail []models.Detail
	Db.Table("details").Where("oaccountaddress = ?", account).Offset(pageInt).Limit(pageSizeInt).Scan(&detail)

	var item response.AccountDealDetail
	var res []response.AccountDealDetail
	for _, v := range detail {
		switch v.Oindex {
		case "0":
			item.From = v.Oaccountaddress
		case "1":
			item.To = v.Oaccountaddress
			item.Amount = v.Oamountvalue
			item.Timestamp = v.Mtimestamp
			item.Tranidentifier = v.Tranidentifier
		case "2":
			item.Fee = v.Oamountvalue
		}
		res = append(res, item)
		item = response.AccountDealDetail{}
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    res,
	})


}

