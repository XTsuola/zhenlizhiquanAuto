package controllers

import (
	"log"
	"net/http"

	"go_project/models"
	"go_project/utils"

	"github.com/gin-gonic/gin"
)

type autoJob struct {
	name    string
	enabled bool
	run     func()
}

func loadAndAdd[T any](file, table, url string) {
	var data []T
	if err := utils.LoadJSON(file, &data); err != nil {
		log.Printf("读取失败 %s: %v", file, err)
		return
	}
	utils.AutoAddTableData(table, url, models.AddData[T]{Data: data})
	log.Printf("同步完成: %s -> %s (%d 条)", file, table, len(data))
}

var autoJobs = []autoJob{
	{name: "skinDiy", enabled: false, run: func() {
		loadAndAdd[models.SkinDiyBase]("data/skinDiy.json", "skin_diy", "/skinDiy/addAll")
	}},
	{name: "cardDiy", enabled: false, run: func() {
		loadAndAdd[models.CardDiyBase]("data/cardDiy.json", "card_diy", "/cardDiy/addAll")
	}},
	{name: "frequency", enabled: true, run: func() {
		loadAndAdd[models.FrequencyBase]("data/frequency.json", "frequency", "/frequency/cardsAddAll")
	}},
	{name: "card", enabled: false, run: func() {
		loadAndAdd[models.CardBase]("data/card.json", "card", "/card/add")
	}},
	{name: "shenqi", enabled: false, run: func() {
		loadAndAdd[models.ShenqiBase]("data/shenqi.json", "shenqi", "/shenqi/add")
	}},
	{name: "hero", enabled: false, run: func() {
		loadAndAdd[models.HeroBase]("data/hero.json", "hero", "/hero/add")
	}},
	{name: "shard", enabled: false, run: func() {
		loadAndAdd[models.ShardBase]("data/shard.json", "shard", "/hero/shardAdd")
	}},
	{name: "skin", enabled: false, run: func() {
		loadAndAdd[models.SkinBase]("data/skin.json", "skin", "/skin/add")
	}},
	{name: "question", enabled: false, run: func() {
		loadAndAdd[models.QuestionBase]("data/question.json", "question", "/question/addAll")
	}},
	{name: "answer", enabled: false, run: func() {
		loadAndAdd[models.AnswerBase]("data/answer.json", "answer", "/answer/addAll")
	}},
	{name: "shijiesai", enabled: false, run: func() {
		loadAndAdd[models.ShijiesaiBase]("data/shijiesai.json", "shijiesai", "/shijiesai/addList")
	}},
	{name: "member", enabled: true, run: func() {
		loadAndAdd[models.MemberBase]("data/member.json", "member", "/member/addAll")
	}},
}

// mysqlAuto 按启用任务异步同步数据
func mysqlAuto(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "自动化已执行",
	})
	go func() {
		for _, job := range autoJobs {
			if !job.enabled {
				continue
			}
			log.Printf("开始同步: %s", job.name)
			job.run()
		}
		log.Println("全部同步任务结束")
	}()
}
