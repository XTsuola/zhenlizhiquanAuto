package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	my "go_project/config"
	"net/http"
	"os"
	"time"
)

var BaseUrl = "http://127.0.0.1:8002"
var Timeout = time.Second * 20

func LoadJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func AutoAddTableData[T any](table string, url string, data T) {
	err := my.DB.Exec("TRUNCATE TABLE " + table).Error // users 是你的表名
	if err != nil {
		fmt.Println(err)
		return
	}
	resErr := POST[T](url, data)
	if resErr != nil {
		fmt.Println(resErr)
		return
	}
}

func POST[T any](url string, addData T) error {
	client := http.Client{
		Timeout: Timeout,
	}
	data, _ := json.Marshal(addData)
	resp, resErr := client.Post(
		BaseUrl+url,
		"application/json; charset=utf-8",
		bytes.NewBuffer(data),
	)
	defer resp.Body.Close()
	return resErr
}
