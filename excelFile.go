package consolecompaniesinnapp

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func ReadData() {

	var confData ConfigData

	readFileData(confData)

	xl, err := xlsx.OpenFile(confData.FileName)
	if err != nil {
		fmt.Println(err)
	}

	sheet := xl.Sheet[confData.SheetName]

	defer xl.Save(confData.FullFilePath + "\\" + confData.FileName)

	for j := confData.StartCell; j < confData.EndCell; j++ {
		cell := sheet.Cell(j, 0)
		tempVal := cell.Value

		inn, fullName, err := getInfo(tempVal)
		if err != nil {
			cell = sheet.Cell(j, 1)
			cell.Value = inn
			cell = sheet.Cell(j, 2)
			cell.Value = fullName
			continue
		}

		cell = sheet.Cell(j, 1)
		cell.Value = inn
		cell = sheet.Cell(j, 2)
		cell.Value = fullName
	}
}
