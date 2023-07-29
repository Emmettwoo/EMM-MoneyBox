package cmd

import (
	"errors"
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

var sheetRowNumberLabel = "row_num"
var requiredRowFieldList = []string{"Category", "Date", "Type", "Amount"}
var importFailedRowNumberList []int
var importIgnoredRowNumberList []int
var importSucceedRowNumberList []int

var importCmd = &cobra.Command{
	Use:   "import {file_path}",
	Short: "import data from excel",
	Long: `Import excel data into your database. 

Example:
  EMM-MoneyBox import /tmp/test.xlsx`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("must provide file path")
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
				util.Logger.Errorw("read sheet rows failed", "error", err)
			}
			util.Logger.Infof("processing sheet %s", currentSheetName)
			var cashFlowMapByDate = readSheetData(rows)
			// fixme: 保存 cashFlowList 時，要考慮事務細粒度，考慮增加 batchInsert()
			for date, cashFlowMapByColumnList := range cashFlowMapByDate {
				saveIntoDB(cashFlowMapByColumnList)
				util.Logger.Debugf("%s of %s's flows imported", util.FormatDateToString(date), currentSheetName)
			}
			util.Logger.Infow("sheet has been imported",
				"sheet_name", currentSheetName,
				"succeed_row", importSucceedRowNumberList,
				"ignored_row", importIgnoredRowNumberList,
				"failed_row", importFailedRowNumberList)
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
func readSheetData(sheetRowCursor *excelize.Rows) map[time.Time][]map[string]string {

	cashFlowMapByDate := make(map[time.Time][]map[string]string)

	// 第一行爲標題行，校驗格式是否正確
	sheetRowCursor.Next()
	var currentRowNumber = 1
	rowColumnList, err := sheetRowCursor.Columns()
	if err != nil {
		util.Logger.Error(err.Error())
	}
	if !isSheetTitleVerified(rowColumnList) {
		return cashFlowMapByDate
	}

	// 遍歷每一行的數據，組裝 CashFlow
	for sheetRowCursor.Next() {
		// 更新當前行號
		currentRowNumber++

		rowColumnList, err = sheetRowCursor.Columns()
		if err != nil {
			util.Logger.Error(err.Error())
		}

		// 依序組裝每一行數據，形成 title-value Map
		var cashFlowMapByColumn = map[string]string{}
		for index, colCell := range rowColumnList {
			cashFlowMapByColumn[defaultRowTitle[index]] = colCell
		}
		cashFlowMapByColumn[sheetRowNumberLabel] = strconv.Itoa(currentRowNumber)

		// 必填欄位校驗
		if !isRequiredFieldSatisfied(currentRowNumber, cashFlowMapByColumn) {
			importFailedRowNumberList = append(importFailedRowNumberList, currentRowNumber)
			continue
		}

		var cashFlowDate = util.FormatDateFromString(cashFlowMapByColumn["Date"])
		cashFlowMapByDate[cashFlowDate] = append(cashFlowMapByDate[cashFlowDate], cashFlowMapByColumn)
	}

	if err = sheetRowCursor.Close(); err != nil {
		util.Logger.Error(err.Error())
	}

	return cashFlowMapByDate
}

func isSheetTitleVerified(titleColumnList []string) bool {
	for index, colCell := range titleColumnList {
		if colCell != defaultRowTitle[index] {
			util.Logger.Warn("sheet title un-expected, parse failed.")
			return false
		}
	}
	return true
}

func isRequiredFieldSatisfied(currentRowNumber int, columnCellMap map[string]string) bool {
	for _, requiredRowField := range requiredRowFieldList {
		if columnCellMap[requiredRowField] == "" {
			util.Logger.Errorw("field could not be empty, import failed",
				sheetRowNumberLabel, currentRowNumber, "field", requiredRowField)
			return false
		}
	}
	return true
}

func saveIntoDB(cashFlowMapByColumnList []map[string]string) {
	for _, cashFlowMapByColumn := range cashFlowMapByColumnList {
		var cashFlowEntity = entity.CashFlowEntity{}.Build(cashFlowMapByColumn)
		if cashFlowEntity.Id != primitive.NilObjectID {
			var existedCashFlow = cashFlowMapper.GetCashFlowByObjectId(cashFlowEntity.Id)
			if !existedCashFlow.IsEmpty() {
				util.Logger.Warnw("cash_flow existed, ignored import.",
					sheetRowNumberLabel, cashFlowMapByColumn[sheetRowNumberLabel],
					"objectId", cashFlowEntity.Id.Hex())
				importIgnoredRowNumberList = append(importIgnoredRowNumberList,
					util.ToInteger(cashFlowMapByColumn[sheetRowNumberLabel]))
				continue
			}
		}
		cashFlowMapper.InsertCashFlowByEntity(cashFlowEntity)
		util.Logger.Debug("cash_flow inserted: " + cashFlowEntity.ToString())
		importSucceedRowNumberList = append(importSucceedRowNumberList,
			util.ToInteger(cashFlowMapByColumn[sheetRowNumberLabel]))
	}
}

func init() {
	rootCmd.AddCommand(importCmd)
}
