package cmd

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var defaultSheetName = "report"
var defaultRowTitle = []string{"Id", "Date", "Category", "Amount", "Desc"}

var exportCmd = &cobra.Command{
	Use:   "export {from_date} {to_date}",
	Short: "Export data to excel",
	Long: `Export specifics data into a excel sheet. 

Example:
  EMM-MoneyBox export 20230101 20231231`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return errors.New("must provide from_date and to_date")
		}

		file := createExcelFile()
		exportData(file, args[0], args[1])
		saveExcelFile(file)
		return nil
	},
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

func saveExcelFile(file *excelize.File) {

	// 根據指定路徑保存活頁簿
	writeExcelRow(file, defaultSheetName, "A2", "Ended Time")
	writeExcelRow(file, defaultSheetName, "B2", time.Now())
	if err := file.SaveAs("export.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func exportData(file *excelize.File, fromDate string, toDate string) {

	var cashFlowRowIndex = 1

	var queryDateCurrent = util.FormatDateFromString(fromDate)
	var queryDateEnded = util.FormatDateFromString(toDate)

	var currentYearAndMonth = "nil"

	for queryDateEnded.After(queryDateCurrent) {
		dayFlow := dayFlowMapper.GetDayFlowByDate(queryDateCurrent)
		if dayFlow.IsEmpty() {
			fmt.Printf("%s's flow is empty.\n", util.FormatDateToString(queryDateCurrent))
			queryDateCurrent = queryDateCurrent.AddDate(0, 0, 1)
			continue
		}

		var queryDateCurrentInString = util.FormatDateToString(queryDateCurrent)
		fmt.Printf("%s's flow is exporting.\n", queryDateCurrentInString)
		cashFlowArray := cashFlowMapper.GetCashFlowsByObjectIdArray(dayFlow.CashFlows)

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
		}

		for _, cashFlow := range cashFlowArray {
			cashFlowRowIndex++
			var cashFlowIndexInString = strconv.Itoa(cashFlowRowIndex)
			writeExcelRow(file, currentYearAndMonth, "A"+cashFlowIndexInString, cashFlow.Id.Hex())
			writeExcelRow(file, currentYearAndMonth, "B"+cashFlowIndexInString, queryDateCurrentInString)
			writeExcelRow(file, currentYearAndMonth, "C"+cashFlowIndexInString, cashFlow.Category)
			writeExcelRow(file, currentYearAndMonth, "D"+cashFlowIndexInString, cashFlow.Amount)
			writeExcelRow(file, currentYearAndMonth, "E"+cashFlowIndexInString, cashFlow.Desc)
		}

		queryDateCurrent = queryDateCurrent.AddDate(0, 0, 1)
	}
}

func writeExcelRow(file *excelize.File, sheetName string, cellPosition string, cellValue interface{}) {
	if err := file.SetCellValue(sheetName, cellPosition, cellValue); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
