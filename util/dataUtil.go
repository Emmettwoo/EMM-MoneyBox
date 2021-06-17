package util

import (
	"encoding/json"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

type CashFlow struct {
	Title  string `xml:"name,attr"`
	Amount float32
	Desc   string
}

type DayCashFlow struct {
	Date          string
	CashFlowArray []CashFlow
}

type UserConfig struct {
	UserName   string
	DayLimit   float32
	MonthLimit float32
}

type DataModel struct {
	UserConfig       UserConfig
	DayCashFlowArray []DayCashFlow
}

func Init() {
	userConfig := UserConfig{
		UserName:   "Emmett Woo",
		DayLimit:   50,
		MonthLimit: 2000,
	}

	dayCashFlow := DayCashFlow{
		Date: "2021-06-17",
		CashFlowArray: []CashFlow{
			{
				Title:  "Initial First",
				Amount: 0,
				Desc:   "Initial User Data With Zero CashFlow",
			},
			{
				Title:  "Initial Second",
				Amount: 0,
				Desc:   "Initial User Data With Zero CashFlow",
			},
		},
	}

	initData := DataModel{
		UserConfig:       userConfig,
		DayCashFlowArray: []DayCashFlow{dayCashFlow},
	}

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

	// 缺点是每一次都得把完整数据读取出来，数据多了以后严重影响读取速度
	var dataModel DataModel
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&dataModel)
	if err != nil {
		fmt.Println("错误: 用户数据读取失败: ", err.Error())
	} else {
		// 遍历用户每日现金流数据并直接输出
		fmt.Printf("\n*** %s ***\n", dataModel.UserConfig.UserName)
		for _, dayCashFlow := range dataModel.DayCashFlowArray {
			fmt.Printf("\n@ %s\n", dayCashFlow.Date)
			for _, cashFlow := range dayCashFlow.CashFlowArray {
				fmt.Println(cashFlow.Title, ": ", cashFlow.Amount)
			}
		}
	}
	fmt.Println("\n*** The End ***")
}

func getDataFilePath() string {
	dataFileName := ".emm-moneybox.json"
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("错误: 获取用户目录失败: %v\n", err)
	}
	return home + "/" + dataFileName
}
