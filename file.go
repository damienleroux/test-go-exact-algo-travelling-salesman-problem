package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadFile(filePath string) []Sale {
	saleJsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Print("[LOAD_SALES_ERROR]", err)
	}
	defer saleJsonFile.Close()
	byteValue, _ := ioutil.ReadAll(saleJsonFile)
	var sales []Sale
	json.Unmarshal(byteValue, &sales)

	return sales
}
