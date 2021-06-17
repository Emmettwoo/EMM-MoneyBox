package util

import (
	"encoding/json"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

type Room struct {
	Name   string `xml:"name,attr"`
	Desc   string
	Member []string
}

func Init() {
	initData := []Room{{"308", "Five Men Family", []string{"Emmettwoo", "Gaspiller", "CZ_CS", "Safina", "CDFeng"}}, {"309", "Single Dog", []string{"LuLuLu"}}}
	// 创建文件
	filePtr, err := os.OpenFile(getDataFilePath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("错误: 用户数据不存在: ", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(initData)
	if err != nil {
		fmt.Println("错误: 用户数据写入失败: ", err.Error())
		return
	}
}

func Read() {
	filePtr, err := os.Open(getDataFilePath())
	if err != nil {
		fmt.Println("错误: 用户数据不存在: ", err.Error())
		return
	}
	defer filePtr.Close()
	var dataArray []Room
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&dataArray)
	if err != nil {
		fmt.Println("错误: 用户数据读取失败: ", err.Error())
	} else {
		for _, data := range dataArray {
			fmt.Println(data.Name)
			for _, member := range data.Member {
				fmt.Println(member)
			}
			fmt.Println()
		}
		// fmt.Println(dataArray)
	}
}

func getDataFilePath() string {
	dataFileName := ".emm-moneybox.json"
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("错误: 获取用户目录失败: %v\n", err)
	}
	return home + "/" + dataFileName
}
