package consolecompaniesinnapp

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func readData() {

	fileName := ""
	filename += ".xlsx"

	xl, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fullPathFile := ""
	sheetName := ""
	linesAmount := 1

	sheet := xl.Sheet[sheetName]

	defer xl.Save(fullPathFile)

	for j := 0; j < linesAmount; j++ {
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
