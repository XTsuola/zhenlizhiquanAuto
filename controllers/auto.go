package controllers

import (
	"fmt"
	"go_project/models"
	"go_project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新增皮肤diy
func skinDiyAdd(url string) {
	var data []models.SkinDiyBase
	if loadErr := utils.LoadJSON("data/skinDiy.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}

	var addData models.SkinDiyAddData
	addData.Data = data
	utils.AutoAddTableData[models.SkinDiyAddData]("skin_diy", url, addData)
}

// 新增卡牌diy
func cardDiyAdd(url string) {
	var data []models.CardDiyBase
	if loadErr := utils.LoadJSON("data/cardDiy.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.CardDiyAddData
	addData.Data = data
	utils.AutoAddTableData[models.CardDiyAddData]("card_diy", url, addData)
}

// 新增卡组diy
func frequencyAdd(url string) {
	var data []models.FrequencyBase
	if loadErr := utils.LoadJSON("data/frequency.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.FrequencyAddAData
	addData.Data = data
	utils.AutoAddTableData[models.FrequencyAddAData]("frequency", url, addData)
}

// 新增种族
func cardAdd(url string) {
	var data []models.CardBase
	if loadErr := utils.LoadJSON("data/card.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.CardAddData
	addData.Data = data
	utils.AutoAddTableData[models.CardAddData]("card", url, addData)
}

// 新增神器
func shenqiAdd(url string) {
	var data []models.ShenqiBase
	if loadErr := utils.LoadJSON("data/shenqi.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.ShenqiAddData
	addData.Data = data
	utils.AutoAddTableData[models.ShenqiAddData]("shenqi", url, addData)
}

// 英雄
func heroAdd(url string) {
	var data []models.HeroBase
	if loadErr := utils.LoadJSON("data/hero.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.HeroddData
	addData.Data = data
	fmt.Println(addData.Data)
	utils.AutoAddTableData[models.HeroddData]("hero", url, addData)
}

// 英雄碎片
func shardAdd(url string) {
	var data []models.ShardBase
	if loadErr := utils.LoadJSON("data/shard.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.ShardAddData
	addData.Data = data
	fmt.Println(addData.Data)
	utils.AutoAddTableData[models.ShardAddData]("shard", url, addData)
}

// 皮肤
func skinAdd(url string) {
	var data []models.SkinBase
	if loadErr := utils.LoadJSON("data/skin.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.SkinAddData
	addData.Data = data
	fmt.Println(addData.Data)
	utils.AutoAddTableData[models.SkinAddData]("skin", url, addData)
}

// 世界赛
func shijiesaiAdd(url string) {
	var data []models.ShijiesaiBase
	if loadErr := utils.LoadJSON("data/shijiesai.json", &data); loadErr != nil {
		fmt.Println("读取失败：", loadErr)
		return
	}
	var addData models.ShijiesaiAddData
	addData.Data = data
	fmt.Println(addData.Data)
	utils.AutoAddTableData[models.ShijiesaiAddData]("shijiesai", url, addData)
}

// 自动化更新数据
func mysqlAuto(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "自动化已执行",
	})
	go func() {
		//skinDiyAdd("/skinDiy/addAll")
		//cardDiyAdd("/cardDiy/addAll")
		//frequencyAdd("/frequency/cardsAddAll")
		cardAdd("/card/add")
		//shenqiAdd("/shenqi/add")
		//heroAdd("/hero/add")
		//shardAdd("/hero/shardAdd")
		//skinAdd("/skin/add")
		//shijiesaiAdd("/shijiesai/addList")
	}()
}
