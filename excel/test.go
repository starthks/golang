package main

import (
	"fmt"

	exl "github.com/xuri/excelize/v2"
)

func Check(e error) {
	if e != nil {
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		panic(e)
	}
}
func Close(file *exl.File) {
	if file != nil {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	f := exl.NewFile()
	defer Close(f)
	for idx, row := range [][]interface{}{
		{nil, "Apple", "Orange", "Pear"}, {"Small", 2, 3, 3},
		{"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
	} {
		cell, err := exl.CoordinatesToCellName(1, idx+1)
		Check(err)
		f.SetSheetRow("sheet1", cell, &row)
	}

	if err := f.AddChart("sheet1", "f1", &exl.Chart{
		Type: exl.Col3DClustered,
		Series: []exl.ChartSeries{
			{
				Name:       "sheet1!$a$2",
				Categories: "sheet1!$b$1:$d$1",
				Values:     "sheet1!$b$2:$d$2",
			},
			{
				Name:       "Sheet1!$A$3",
				Categories: "Sheet1!$B$1:$D$1",
				Values:     "Sheet1!$B$3:$D$3",
			},
			{
				Name:       "Sheet1!$A$4",
				Categories: "Sheet1!$B$1:$D$1",
				Values:     "Sheet1!$B$4:$D$4",
			},
		},
		Title: []exl.RichTextRun{
			{
				Text: "Fruit 3D Clustered",
			},
		},
        Legend: exl.ChartLegend{
            ShowLegendKey: false,
        },
	}); err != nil {
		Check(err)
	}
	if err := f.SaveAs("book1.xlsx"); err != nil {
		Check(err)
	}
}
