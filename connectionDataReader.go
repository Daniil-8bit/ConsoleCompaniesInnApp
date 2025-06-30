package consolecompaniesinnapp

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigData struct {
	FileName     string `json:"fileName"`
	FullFilePath string `json:"fullFilePath"`
	StartCell    int    `json:"startCell"`
	EndCell      int    `json:"endCell"`
	SheetName    string `json:"sheetName"`
}

func readFileData(cd ConfigData) {

	fileData, err := os.ReadFile("ConnectionStrings.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(fileData, &cd)

	if err != nil {
		fmt.Println(err)
		return
	}
}
