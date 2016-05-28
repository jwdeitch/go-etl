package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"time"
	"strconv"
	"io"
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
	http.HandleFunc("/recieve", receive)
	err := http.ListenAndServe(":9090", nil)
	Check(err)
}

func receive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseMultipartForm(10 << 20) // 10 Megabytes

	file, handler, err := r.FormFile("spreadsheet")
	Check(err)

	defer file.Close()
	time := strconv.Itoa(int(time.Now().Unix()))
	filePath := "./uploads/"+handler.Filename + "_" + time
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	Check(err)

	defer f.Close()
	io.Copy(f, file)

	openXlsx(filePath)
	json, err := json.Marshal(Rows)
	Check(err)
	fmt.Fprintf(w, string(json))

}

func process(w http.ResponseWriter, r *http.Request) {
	openXlsx("/Users/jordan1/Desktop/Untitled.xlsx")
	json, err := json.Marshal(Rows)
	Check(err)
	fmt.Fprintf(w, string(json))
}

func openXlsx(path string) {
	xlFile, err := xlsx.OpenFile(path)
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
