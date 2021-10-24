package main

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

var arraySliceTest = []string{
	"AA",
	"BB",
}

func TestPrintGlobalArray(t *testing.T) {
	fmt.Println(arraySliceTest)
}

type arrayExample struct {
	name string
}

func TestStrukturDataMap(t *testing.T) {
	arrayExData := &[]arrayExample{
		{
			name: "name 1",
		},
		{
			name: "name 2",
		},
	}

	var data = struct {
		item      string
		branchid  int64
		itemArray *[]arrayExample
	}{
		item:      "zein",
		branchid:  1,
		itemArray: arrayExData,
	}

	fmt.Printf("%v", data)
}

type Data struct {
	item      string
	branchID  int64
	itemArray []arrayExample
}

// func TestMapStructureData(t *testing.T) {
// 	arrayExData := []arrayExample{
// 		{
// 			name: "name 1",
// 		},
// 		{
// 			name: "name 2",
// 		},
// 	}

// 	data := Data{
// 		item:      "item 1",
// 		branchID:  2,
// 		itemArray: arrayExData,
// 	}

// 	var m map[int]Data

// }

func TestArrayAppend(t *testing.T) {
	var branchValue []string
	if true {
		branchValue = append(branchValue, "test")
	}

	if true {
		branchValue = append(branchValue, "test2")

	}

	fmt.Print(branchValue)
}

func TestGenerateHeaderXLSX(t *testing.T) {
	xlxs, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		t.Error(err)
	}
	sheet1Name := "Sheet One"

	xlxs.SetSheetName(xlxs.GetSheetName(1), sheet1Name)

	// Setup Default Header
	for k, v := range ExcelMappingCell {
		xlxs.SetCellValue(sheet1Name, k, v)
	}
	// Delete BookUpdate.xlsx first
	err = xlxs.SaveAs("BookUpdate.xlsx")
	if err != nil {
		t.Error(err)
	}
}
