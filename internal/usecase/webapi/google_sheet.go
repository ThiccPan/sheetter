package webapi

import (
	"fmt"
	"log"

	"github.com/thiccpan/sheetter/internal/entity"
	"google.golang.org/api/sheets/v4"
)

type SheetApi struct {
	srv     *sheets.Service
	sheetId string
}

func NewSheetApi(srv *sheets.Service, sheetId string) *SheetApi {
	return &SheetApi{
		srv: srv,
		sheetId: sheetId,
	}
}

func (sa *SheetApi) WriteToRow(user entity.User) error {

	writeRange := "A" + fmt.Sprint(user.Row) + ":B"
	writeData := [][]interface{}{
		{
			user.Name,
		},
		{
			user.Email,
		},
	}

	writeValue := sheets.ValueRange{
		MajorDimension: "COLUMNS",
		Values:         writeData,
	}

	writeReq := sa.srv.
		Spreadsheets.
		Values.
		Update(
			sa.sheetId,
			writeRange,
			&writeValue,
		).
		ValueInputOption("RAW")

	writeRes, err := writeReq.Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(writeRes.UpdatedData)
	return nil
}

func (sa *SheetApi) ReadFromSheet() ([]entity.User, error) {
	// Read from sheet with id below:
	spreadsheetId := sa.sheetId
	readRange := "A1:B"
	resp, err := sa.srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	data := []entity.User{}
	for _, row := range resp.Values {
		if len(row) == 0 {
			break
		}
		
		data = append(data, entity.User{
			Name: row[0].(string),
			Email: row[1].(string),
		})
	}

	return data, nil
}
