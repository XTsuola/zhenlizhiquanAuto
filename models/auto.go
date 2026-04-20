package models

type SkinDiyBase struct {
	CardId string `json:"cardId" gorm:"column:cardId"`
	Name   string `json:"name"`
	Skill  string `json:"skill"`
	Effect string `json:"effect"`
	Reason string `json:"reason"`
	Remark string `json:"remark"`
}
type SkinDiyAddData struct {
	Data []SkinDiyBase `json:"data"`
}

type CardDiyBase struct {
	Name     string `json:"name"`
	Zhenyin  int    `json:"zhenyin"`
	Cost     int    `json:"cost"`
	Quality  int    `json:"quality"`
	CardType int    `json:"cardType" gorm:"column:cardType"`
	Att      int    `json:"att"`
	Life     int    `json:"life"`
	Effect   string `json:"effect"`
	Img      string `json:"img"`
	Info     string `json:"info"`
	Remark   string `json:"remark"`
}
type CardDiyAddData struct {
	Data []CardDiyBase `json:"data"`
}

type FrequencyBase struct {
	Name     string `json:"name"`
	Qu       int    `json:"qu"`
	HeroId   int    `json:"heroId" gorm:"column:heroId"`
	HeroLife int    `json:"heroLife" gorm:"column:heroLife"`
	Cards    string `json:"cards"`
	Time     string `json:"time"`
}
type FrequencyAddAData struct {
	Data []FrequencyBase `json:"data"`
}

type CardData struct {
	Attack int    `json:"attack"`
	Life   int    `json:"life"`
	Effect string `json:"effect"`
}

type CardBase struct {
	Name    string     `json:"name"`
	Zhenyin int        `json:"zhenyin"`
	Quality int        `json:"quality"`
	Cost    int        `json:"cost"`
	Type    int        `json:"type"`
	Img     string     `json:"img"`
	Grade   string     `json:"grade"`
	Tag     string     `json:"tag"`
	Data    []CardData `json:"data"`
}

type CardAddData struct {
	Data []CardBase `json:"data"`
}

type ShenqiData struct {
	Effect string `json:"effect"`
}

type ShenqiBase struct {
	Name    string       `json:"name"`
	Zhenyin int          `json:"zhenyin"`
	Quality int          `json:"quality"`
	Type    int          `json:"type"`
	Img     string       `json:"img"`
	Data    []ShenqiData `json:"data"`
}

type ShenqiAddData struct {
	Data []ShenqiBase `json:"data"`
}

type HeroData struct {
	Effect string `json:"effect"`
}

type HeroBase struct {
	Name      string     `json:"name"`
	Quality   int        `json:"quality"`
	Zhu       int        `json:"zhu"`
	Fu        int        `json:"fu"`
	SkillName string     `json:"skillName" gorm:"column:skillName"`
	Img       string     `json:"img"`
	Data      []HeroData `json:"data"`
}

type HeroddData struct {
	Data []HeroBase `json:"data"`
}

type ShardBase struct {
	Quality   int    `json:"quality"`
	LevelData string `json:"levelData" gorm:"column:levelData"`
	SkillData string `json:"skillData" gorm:"column:skillData"`
}

type ShardAddData struct {
	Data []ShardBase `json:"data"`
}

type SkinBase struct {
	CardId  int      `json:"cardId" gorm:"column:cardId"`
	Name    string   `json:"name"`
	Zhenyin int      `json:"zhenyin"`
	Cost    int      `json:"cost"`
	Skill   string   `json:"skill"`
	Img     string   `json:"img"`
	Shuxing string   `json:"shuxing"`
	Origin  string   `json:"origin"`
	Remark  string   `json:"remark"`
	Effect  []string `json:"effect"`
}

type SkinAddData struct {
	Data []SkinBase `json:"data"`
}

type AnswerBase struct {
	QuestionId int    `json:"questionId" gorm:"column:questionId"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Time       string `json:"time"`
}

type AnswerAddData struct {
	Data []AnswerBase `json:"data"`
}

type XuanshouInfo struct {
	Name string `json:"name"`
	Kedu int    `json:"kedu" gorm:"column:kedu"`
	Hero []int  `json:"hero" gorm:"column:hero"`
}

type ShijiesaiBase struct {
	No          int `json:"no"`
	AInfo       XuanshouInfo
	BInfo       XuanshouInfo
	ShengfuList []int `json:"shengfuList" gorm:"column:shengfuList"`
}

type ShijiesaiAddData struct {
	Data []ShijiesaiBase `json:"data"`
}
