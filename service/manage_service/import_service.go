package manage_service

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

var sheetRowNumberLabel = "row_num"
var requiredRowFieldList = []string{"BelongsDate", "FlowType", "Amount"}
var importFailedRowNumberList []int
var importIgnoredRowNumberList []int
var importSucceedRowNumberList []int

func ImportService(filePath string) error {

	// 打開並讀取目標文件
	file := readExcelFile(filePath)
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

		importSucceedRowNumberList = []int{}
		importIgnoredRowNumberList = []int{}
		importFailedRowNumberList = []int{}

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
			util.Logger.Debugf("%s of %s's flows imported", util.FormatDateToStringWithoutDash(date), currentSheetName)
		}
		util.Logger.Infow("sheet has been imported",
			"sheet_name", currentSheetName,
			"succeed_row", importSucceedRowNumberList,
			"ignored_row", importIgnoredRowNumberList,
			"failed_row", importFailedRowNumberList)
	}
	return nil
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
		// check category info and get the correct id
		var newCategoryId = handleCategoryInfo(
			cashFlowMapByColumn["CategoryId"], cashFlowMapByColumn["CategoryName"])
		if newCategoryId == "" {
			fmt.Println("failed: row " + strconv.Itoa(currentRowNumber) + ": category not satisfied")
			importFailedRowNumberList = append(importFailedRowNumberList, currentRowNumber)
			continue
		}
		cashFlowMapByColumn["CategoryId"] = newCategoryId

		// 必填欄位校驗
		if !isRequiredFieldSatisfied(currentRowNumber, cashFlowMapByColumn) {
			fmt.Println("failed: row " + strconv.Itoa(currentRowNumber) + ": required field not satisfied")
			importFailedRowNumberList = append(importFailedRowNumberList, currentRowNumber)
			continue
		}

		var cashFlowDate = util.FormatDateFromStringWithoutDash(cashFlowMapByColumn["BelongsDate"])
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

func handleCategoryInfo(categoryId string, categoryName string) string {

	// use category id to fetch first
	if categoryId != "" {
		var categoryEntity = mapper.CategoryCommonMapper.GetCategoryByObjectId(categoryId)
		if !categoryEntity.IsEmpty() {
			return categoryEntity.Id.Hex()
		}
		util.Logger.Warnw("category not existed", "category_id", categoryId)
	}

	// if category name is empty, fail it.
	if categoryName == "" {
		return ""
	}

	// use category name to fetch correct id
	categoryEntity := mapper.CategoryCommonMapper.GetCategoryByName(categoryName)
	if !categoryEntity.IsEmpty() {
		return categoryEntity.Id.Hex()
	}
	util.Logger.Warnw("category not existed", "category_name", categoryName)

	// create new category for this flow
	var plainId = mapper.CategoryCommonMapper.InsertCategoryByEntity(entity.CategoryEntity{
		Name:   categoryName,
		Remark: "create by import",
	})
	return plainId
}

func saveIntoDB(cashFlowMapByColumnList []map[string]string) {
	for _, cashFlowMapByColumn := range cashFlowMapByColumnList {
		var cashFlowEntity = entity.CashFlowEntity{}.Build(cashFlowMapByColumn)
		if cashFlowEntity.Id != primitive.NilObjectID {
			var existedCashFlow = mapper.CashFlowCommonMapper.GetCashFlowByObjectId(cashFlowEntity.Id.Hex())
			if !existedCashFlow.IsEmpty() {
				util.Logger.Warnw("cash_flow existed, ignored import.",
					sheetRowNumberLabel, cashFlowMapByColumn[sheetRowNumberLabel],
					"objectId", cashFlowEntity.Id.Hex())
				fmt.Println("ignored: row " + cashFlowMapByColumn[sheetRowNumberLabel] + ": cash_flow existed")
				importIgnoredRowNumberList = append(importIgnoredRowNumberList,
					util.ToInteger(cashFlowMapByColumn[sheetRowNumberLabel]))
				continue
			}
		}
		var newPlainId = mapper.CashFlowCommonMapper.InsertCashFlowByEntity(cashFlowEntity)
		cashFlowEntity.Id = util.Convert2ObjectId(newPlainId)
		util.Logger.Debug("cash_flow inserted: " + cashFlowEntity.ToString())
		fmt.Println("succeed: row " + cashFlowMapByColumn[sheetRowNumberLabel] + ": cash_flow saved")
		importSucceedRowNumberList = append(importSucceedRowNumberList,
			util.ToInteger(cashFlowMapByColumn[sheetRowNumberLabel]))
	}
}
