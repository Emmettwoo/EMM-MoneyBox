package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/emmettwoo/EMM-MoneyBox/model"

	homedir "github.com/mitchellh/go-homedir"
)

func Init() {
	userConfig := model.UserConfig{
		UserName:   "Emmett Woo",
		DayLimit:   50,
		MonthLimit: 2000,
	}

	dayCashFlow := model.DayCashFlow{
		Date: "2021-06-17",
		CashFlowArray: []model.CashFlow{
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

	initData := model.DataModel{
		UserConfig:       userConfig,
		DayCashFlowArray: []model.DayCashFlow{dayCashFlow},
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
	var dataModel model.DataModel
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&dataModel)
	if err != nil {
		fmt.Println("错误: 用户数据读取失败: ", err.Error())
	} else {
		showDataModel(dataModel)
	}
}

func Write(date string, title string, amount float32, desc string) {
	// TODO: 验证数字是否大于零（暂不考虑收入记录），日期是否符合格式要求。

	// 获取指定date已有的现金流数据（不存在时返回现金流为空的对象）
	dayCashFlow := getDayCashFlow(date)
	showDayCashFlow(dayCashFlow)

	// TODO: 在原现金流数据基础上写入新数据

	// TODO:  将更新后的数据写入JSON文件
}

func getDayCashFlow(date string) model.DayCashFlow {
	filePtr, err := os.Open(getDataFilePath())
	if err != nil {
		fmt.Println("错误: 用户数据不存在: ", err.Error())
		return model.DayCashFlow{}
	}
	defer filePtr.Close()

	// 缺点是每一次都得把完整数据读取出来，数据多了以后严重影响读取速度
	var dataModel model.DataModel
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&dataModel)
	if err != nil {
		fmt.Println("错误: 用户数据读取失败: ", err.Error())
	} else {
		// 获取指定日期的现金流
		for _, dayCashFlow := range dataModel.DayCashFlowArray {
			if dayCashFlow.Date == date {
				return dayCashFlow
			}
		}
	}
	return model.DayCashFlow{Date: date}
}

func showDataModel(dataModel model.DataModel) {
	// 遍历用户每日现金流数据并直接输出
	fmt.Printf("\n*** %s ***\n", dataModel.UserConfig.UserName)
	for _, dayCashFlow := range dataModel.DayCashFlowArray {
		showDayCashFlow(dayCashFlow)
	}
	fmt.Println("\n*** The End ***")
}

func showDayCashFlow(dayCashFlow model.DayCashFlow) {
	fmt.Printf("\n@ %s\n", dayCashFlow.Date)
	for _, cashFlow := range dayCashFlow.CashFlowArray {
		fmt.Println(cashFlow.Title, ": ", cashFlow.Amount)
	}
}

func getDataFilePath() string {
	dataFileName := ".EMM-MoneyBox.json"
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("错误: 获取用户目录失败: %v\n", err)
	}
	return home + "/" + dataFileName
}
