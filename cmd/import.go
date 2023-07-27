package cmd

import (
	"errors"
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var importCmd = &cobra.Command{
	Use:   "import {file_name}",
	Short: "Import data from excel",
	Long: `Import excel data into your database. 

Example:
  EMM-MoneyBox import /tmp/test.xlsx`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("must provide file name")
		}

		// 打開並讀取目標文件
		file := readExcelFile(args[0])
		if file == nil {
			return errors.New("can not read data from file")
		}
		// 記得執行完成關閉文件
		defer func() {
			if err := file.Close(); err != nil {
				util.Logger.Error(err.Error())
			}
		}()

		// 獲取工作表列表，遍歷讀取數據
		var sheetNameList = file.GetSheetList()
		for _, currentSheetName := range sheetNameList {
			// report sheet 非數據表，不計入
			if currentSheetName == defaultSheetName {
				continue
			}

			rows, err := file.Rows(currentSheetName)
			if err != nil {
				util.Logger.Error(err.Error())
			}
			var cashFlowMap = readSheetData(rows)
			util.Logger.Infof("%s's flows read", currentSheetName)
			// fixme: 保存 cashFlowList 時，要考慮事務細粒度，考慮增加 batchInsert()
			for date, cashFlowList := range cashFlowMap {
				saveIntoDB(date, cashFlowList)
				util.Logger.Infof("%s of %s's flows imported", util.FormatDateToString(date), currentSheetName)
			}
		}
		return nil
	},
}

func readExcelFile(fileName string) *excelize.File {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		util.Logger.Error(err.Error())
	}
	return file
}

/**
 * 讀取工作表的數據，以 date 爲 key 整理 cashFlows
 */
func readSheetData(sheetRowCursor *excelize.Rows) map[time.Time][]entity.CashFlowEntity {

	cashFlowMap := make(map[time.Time][]entity.CashFlowEntity)

	// 第一行爲標題行，校驗格式是否正確
	sheetRowCursor.Next()
	rowColumnList, err := sheetRowCursor.Columns()
	if err != nil {
		util.Logger.Error(err.Error())
	}
	if !verifySheetTitle(rowColumnList) {
		return cashFlowMap
	}

	// 遍歷每一行的數據，組裝 CashFlow
	for sheetRowCursor.Next() {
		rowColumnList, err = sheetRowCursor.Columns()
		if err != nil {
			util.Logger.Error(err.Error())
		}

		// 依序組裝每一行數據，形成 title-value Map
		var columnCellMap = map[string]string{}
		for index, colCell := range rowColumnList {
			columnCellMap[defaultRowTitle[index]] = colCell
		}
		var cashFlow = entity.CashFlowEntity{}.Build(columnCellMap)
		var cashFlowDate = time.Now()
		if columnCellMap["Date"] != "" {
			cashFlowDate = util.FormatDateFromString(columnCellMap["Date"])
		}
		cashFlowMap[cashFlowDate] = append(cashFlowMap[cashFlowDate], cashFlow)
	}

	if err = sheetRowCursor.Close(); err != nil {
		util.Logger.Error(err.Error())
	}

	return cashFlowMap
}

func verifySheetTitle(titleColumnList []string) bool {
	for index, colCell := range titleColumnList {
		if colCell != defaultRowTitle[index] {
			util.Logger.Warn("sheet title un-expected, parse failed.")
			return false
		}
	}
	return true
}

func saveIntoDB(cashFlowDate time.Time, cashFlowList []entity.CashFlowEntity) {
	for _, cashFlow := range cashFlowList {
		if cashFlow.Id != primitive.NilObjectID {
			var existedCashFlow = cashFlowMapper.GetCashFlowByObjectId(cashFlow.Id)
			if !existedCashFlow.IsEmpty() {
				util.Logger.Debugw("cash_flow existed, ignored import.",
					"objectId", cashFlow.Id.Hex())
				continue
			}
		}
		cashFlowMapper.InsertCashFlowByEntity(cashFlow, cashFlowDate)
		util.Logger.Debug("cash_flow inserted: " + cashFlow.ToString())
	}
}

func init() {
	rootCmd.AddCommand(importCmd)
}
