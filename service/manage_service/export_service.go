package manage_service

import (
	"errors"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

var defaultSheetName = "report"
var defaultRowTitle = []string{"Id", "CategoryId", "CategoryName", "BelongsDate", "FlowType", "Amount", "Description"}

func ExportService(fromDateInString string, toDateInString string, filePath string) error {

	if filePath == "" {
		filePath = "./export.xlsx"
	}
	var fromDate = util.FormatDateFromStringWithoutDash(fromDateInString)
	var toDate = util.FormatDateFromStringWithoutDash(toDateInString)
	if err := isExportRequiredFiledSatisfied(fromDate, toDate, filePath); err != nil {
		return err
	}

	file := createExcelFile()
	exportData(file, fromDateInString, toDateInString)
	saveExcelFile(file, filePath)
	return nil
}

func isExportRequiredFiledSatisfied(fromDate time.Time, toDate time.Time, filePath string) error {

	if util.IsDateTimeEmpty(fromDate) {
		return errors.New("from_date could not be empty")
	}
	if util.IsDateTimeEmpty(toDate) {
		return errors.New("to_date could not be empty")
	}
	if fromDate.After(toDate) {
		return errors.New("from_date should before to_date")
	}
	if !strings.HasSuffix(filePath, ".xlsx") {
		return errors.New("file_path should be end with '.xlsx'")
	}

	return nil
}

func createExcelFile() *excelize.File {

	var file = excelize.NewFile()
	// 創建一個工作表
	index, _ := file.NewSheet(defaultSheetName)
	// 設置活頁簿的默認工作表
	file.SetActiveSheet(index)
	// 設置存儲格的值
	writeExcelRow(file, defaultSheetName, "A1", "Start Time")
	writeExcelRow(file, defaultSheetName, "B1", time.Now())
	// 刪除默認的 Sheet1 表
	_ = file.DeleteSheet("Sheet1")

	return file
}

func saveExcelFile(file *excelize.File, filePath string) {

	// 根據指定路徑保存活頁簿
	writeExcelRow(file, defaultSheetName, "A2", "Ended Time")
	writeExcelRow(file, defaultSheetName, "B2", time.Now())
	// fixme: support specific output path
	if err := file.SaveAs(filePath); err != nil {
		util.Logger.Errorln(err)
	}
}

func exportData(file *excelize.File, fromDate, toDate string) {

	var cashFlowRowIndex = 1

	var queryDateCurrent = util.FormatDateFromStringWithoutDash(fromDate)
	// add one day for include the last day's data
	var queryDateEnded = util.FormatDateFromStringWithoutDash(toDate).AddDate(0, 0, 1)

	var currentYearAndMonth = "nil"

	for queryDateEnded.After(queryDateCurrent) {
		cashFlowArray := mapper.CashFlowCommonMapper.GetCashFlowsByBelongsDate(queryDateCurrent)
		if len(cashFlowArray) == 0 {
			util.Logger.Debugf("%s's flow is empty.\n", util.FormatDateToStringWithoutDash(queryDateCurrent))
			queryDateCurrent = queryDateCurrent.AddDate(0, 0, 1)
			continue
		}

		var queryDateCurrentInString = util.FormatDateToStringWithoutDash(queryDateCurrent)
		util.Logger.Debugf("%s's flow is exporting.\n", queryDateCurrentInString)

		// 年份有變化，則初始化新 Sheet
		var newYearAndMonth = queryDateCurrentInString[0:6]
		if newYearAndMonth != currentYearAndMonth {
			currentYearAndMonth = newYearAndMonth

			_, _ = file.NewSheet(currentYearAndMonth)

			// 這裏存在一個問題，若年月回溯，Index 已失效，好在是由程式控制递增。
			cashFlowRowIndex = 1
			writeExcelRow(file, currentYearAndMonth, "A1", defaultRowTitle[0])
			writeExcelRow(file, currentYearAndMonth, "B1", defaultRowTitle[1])
			writeExcelRow(file, currentYearAndMonth, "C1", defaultRowTitle[2])
			writeExcelRow(file, currentYearAndMonth, "D1", defaultRowTitle[3])
			writeExcelRow(file, currentYearAndMonth, "E1", defaultRowTitle[4])
			writeExcelRow(file, currentYearAndMonth, "F1", defaultRowTitle[5])
			writeExcelRow(file, currentYearAndMonth, "G1", defaultRowTitle[6])
		}

		for _, cashFlow := range cashFlowArray {
			cashFlowRowIndex++
			var cashFlowIndexInString = strconv.Itoa(cashFlowRowIndex)
			// refer to hardcode defaultRowTitle, bad idea.
			writeExcelRow(file, currentYearAndMonth, "A"+cashFlowIndexInString, cashFlow.Id.Hex())
			writeExcelRow(file, currentYearAndMonth, "B"+cashFlowIndexInString, cashFlow.CategoryId.Hex())
			writeExcelRow(file, currentYearAndMonth, "C"+cashFlowIndexInString,
				mapper.CategoryCommonMapper.GetCategoryByObjectId(cashFlow.CategoryId.Hex()).Name)
			writeExcelRow(file, currentYearAndMonth, "D"+cashFlowIndexInString, queryDateCurrentInString)
			writeExcelRow(file, currentYearAndMonth, "E"+cashFlowIndexInString, cashFlow.FlowType)
			writeExcelRow(file, currentYearAndMonth, "F"+cashFlowIndexInString, cashFlow.Amount)
			writeExcelRow(file, currentYearAndMonth, "G"+cashFlowIndexInString, cashFlow.Description)
		}

		queryDateCurrent = queryDateCurrent.AddDate(0, 0, 1)
	}
}

func writeExcelRow(file *excelize.File, sheetName, cellPosition string, cellValue interface{}) {
	if err := file.SetCellValue(sheetName, cellPosition, cellValue); err != nil {
		util.Logger.Errorln(err)
	}
}
