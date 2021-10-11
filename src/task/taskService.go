package task

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"io/ioutil"
	"net/http"

	"icpscan/src/controllers"
	"icpscan/src/models"

	"github.com/bitly/go-simplejson"
	"github.com/robfig/cron/v3"
)

type TaskService interface {
	PullBlockDetail()
	Run()
}

type taskService struct {
	cron *cron.Cron
}

func NewTaskService(cronCli *cron.Cron) (TaskService, error) {
	return &taskService{
		cron: cronCli,
	}, nil
}

type Identifier struct {
	NetworkIdentifier Network `json:"network_identifier"`
	BlockIdentifier   Block   `json:"block_identifier"`
}

type Network struct {
	BlockChain string `json:"blockchain"`
	Network    string `json:"network"`
}

type Block struct {
	Index int64 `json:"index"`
}

func (obj *Identifier) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, obj); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *taskService) PullBlockDetail() {
	var blockId []models.Block
	var ih int64
	err := controllers.Db.Table("blocks").Select("mblockheight").Order("mblockheight desc").Limit(1).Scan(&blockId).Error
	fmt.Printf("err is %+v\n", err)
	fmt.Printf("blockId is %+v\n", blockId)

	if len(blockId) == 0 {
		ih = -1
	} else {
		fmt.Println("pppp", blockId[0].Mblockheight)
		ih, err = strconv.ParseInt(blockId[0].Mblockheight, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
	}

	// if err == gorm.ErrRecordNotFound {

	// } else if err != nil {
	// 	fmt.Printf("err is %+v", err)
	// 	return
	// } else {

	// }

	j := ih + 1
	fmt.Printf("j value is %d, type is %T", j, j)
	//return
	// redis锁
	currentBlockHeightKey := strconv.FormatInt(j, 10)
	ctx := context.Background()
	pong, err := controllers.RedisDb.Ping(ctx).Result()
	fmt.Println(pong, err)
	value := controllers.RedisDb.Get(ctx, currentBlockHeightKey)
	fmt.Printf("redis value is %+v\n", value)
	if value.Val() != "" {
		return
	}
	controllers.RedisDb.Set(ctx, currentBlockHeightKey, 1, 5)
	value1 := controllers.RedisDb.Get(ctx, currentBlockHeightKey)
	fmt.Printf("redis value is Get %+v\n", value1)
	identify := Identifier{}

	identify.NetworkIdentifier.BlockChain = "Internet Computer"
	identify.NetworkIdentifier.Network = "00000000000000020101"
	identify.BlockIdentifier.Index = j

	jsonStu, err := json.Marshal(identify)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))

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
	fmt.Printf("resp is %+v\n", resp)
	fmt.Printf("err is %+v\n", err)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("res is %+v\n", str)

	js, err2 := simplejson.NewJson(str)

	if err2 != nil {
		fmt.Println("抓取fail", err2)
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
	if len(tData) == 0 {
		fmt.Println("data is empty")
		return
	}

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

					if op, ok := plll["currency"].(map[string]interface{}); ok {
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
				Tranidentifier:          tidentify,
				Oindex:                  oindex,
				Otype:                   otype,
				Ostatus:                 ostatus,
				Oaccountaddress:         oaddress,
				Oamountvalue:            ovalue,
				Oamountcurrencydecimals: odecimals,
				Oamountcurrencysymbol:   osymbol,
				Mblockheight:            strconv.Itoa(bheight),
				Mmemo:                   strconv.FormatUint(mmemo, 10),
				Mtimestamp:              strconv.Itoa(utime),
			}

			controllers.Db.Create(&detail)
			fmt.Printf("detail is %+v", detail)

		}
	}

	// 存储外层信息
	block := models.Block{
		Blockidentifier: strconv.Itoa(bindex),
		Parentblock:     strconv.Itoa(preindex),
		Transactiohash:  cindex,
		Mblockheight:    strconv.Itoa(bheight),
		Mmemo:           strconv.FormatUint(mmemo, 10),
		Tranidentifier:  tidentify,
		Mtimestamp:      strconv.Itoa(utime),
		Blocktimestamp:  strconv.Itoa(bttime),
	}
	fmt.Printf("block is %+v", block)

	controllers.Db.Create(&block)
	i := controllers.RedisDb.Del(ctx, currentBlockHeightKey)
	fmt.Printf("redis delete is %+v\n", i)

}

func (t *taskService) pullAllBlock(blockHeight int64) {
	//return
	// redis锁
	currentBlockHeightKey := strconv.FormatInt(blockHeight, 10)
	ctx := context.Background()
	pong, err := controllers.RedisDb.Ping(ctx).Result()
	fmt.Println(pong, err)
	value := controllers.RedisDb.Get(ctx, currentBlockHeightKey)
	if value.Val() != "" {
		return
	}
	controllers.RedisDb.Set(ctx, currentBlockHeightKey, 1, 5)
	identify := Identifier{}

	identify.NetworkIdentifier.BlockChain = "Internet Computer"
	identify.NetworkIdentifier.Network = "00000000000000020101"
	identify.BlockIdentifier.Index = blockHeight

	jsonStu, err := json.Marshal(identify)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))

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
		fmt.Printf("err is %+v\n", err)
		panic(err)
	}

	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	js, err2 := simplejson.NewJson(str)

	if err2 != nil {
		fmt.Println("抓取fail", err2)
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

	// 转账数据
	var tData = js.Get("block").Get("transactions").GetIndex(0).Get("operations").MustArray()

	if len(tData) == 0 {
		fmt.Println("data is empty")
		return
	}

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

					if op, ok := plll["currency"].(map[string]interface{}); ok {
						odecimals = string(op["decimals"].(json.Number))
						osymbol = string(op["symbol"].(string))
					}
				}
			}

			var otype string
			var ostatus string

			otype = each_map["type"].(string)
			ostatus = each_map["status"].(string)

			detail := models.Detail{
				Tranidentifier:          tidentify,
				Oindex:                  oindex,
				Otype:                   otype,
				Ostatus:                 ostatus,
				Oaccountaddress:         oaddress,
				Oamountvalue:            ovalue,
				Oamountcurrencydecimals: odecimals,
				Oamountcurrencysymbol:   osymbol,
				Mblockheight:            strconv.Itoa(bheight),
				Mmemo:                   strconv.FormatUint(mmemo, 10),
				Mtimestamp:              strconv.Itoa(utime),
			}

			controllers.Db.Create(&detail)
			fmt.Printf("detail is %+v", detail)

		}
	}

	// 存储外层信息
	block := models.Block{
		Blockidentifier: strconv.Itoa(bindex),
		Parentblock:     strconv.Itoa(preindex),
		Transactiohash:  cindex,
		Mblockheight:    strconv.Itoa(bheight),
		Mmemo:           strconv.FormatUint(mmemo, 10),
		Tranidentifier:  tidentify,
		Mtimestamp:      strconv.Itoa(utime),
		Blocktimestamp:  strconv.Itoa(bttime),
	}
	fmt.Printf("block is %+v", block)

	controllers.Db.Create(&block)
	i := controllers.RedisDb.Del(ctx, currentBlockHeightKey)
	fmt.Printf("redis delete is %+v\n", i)
}

func (t *taskService) whilePull() {
	var max int64 = 9999999
	var blockId []models.Block
	var ih int64
	err := controllers.Db.Table("blocks").Select("mblockheight").Order("mblockheight desc").Limit(1).Scan(&blockId).Error
	if err != nil {
		fmt.Println(err)
	}
	ih, err = strconv.ParseInt(blockId[0].Mblockheight, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	needPull := ih + 1
	for ; needPull < max; needPull++ {
		t.pullAllBlock(needPull)
	}
}

func (t *taskService) Run() {
	// id, err := t.cron.AddFunc("@every 3s", t.PullBlockDetail)
	id, err := t.cron.AddFunc("32 06 11 10 *", t.whilePull)
	if err != nil {
		log.Printf("PullBlockDetail err %v", err)
	}

	log.Printf("PullBlockDetail run task id is %d", id)

	t.cron.Start()
}
