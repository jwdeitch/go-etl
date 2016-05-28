package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"net/http"
	"encoding/json"
	. "github.com/inturn/go-helpers"
)

type Row struct {
	Cells      []string
	LineNumber int
}

var Rows []Row

func main() {
	fmt.Println("running: http://localhost:9090")

	http.HandleFunc("/process", process)
	err := http.ListenAndServe(":9090", nil)
	Check(err)
}

func process(w http.ResponseWriter, r *http.Request) {
	openXlsx("/Users/jordan1/Desktop/Untitled.xlsx")
	json, err := json.Marshal(Rows)
	Check(err)
	fmt.Fprintf(w, string(json))
}

func openXlsx(path string) {
	excelFileName := path
	xlFile, err := xlsx.OpenFile(excelFileName)
	Check(err)

	for _, sheet := range xlFile.Sheets {
		for rowNumber, row := range sheet.Rows {
			if rowNumber == 10 {
				break
			}

			var Cells []string
			for _, cell := range row.Cells {
				Cells = append(Cells,cell.Value)
			}

			Rows = append(Rows, Row{Cells,rowNumber})
		}
	}

}
