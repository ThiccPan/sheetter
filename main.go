package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	gsheetConf "github.com/thiccpan/sheetter/config"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)	

func main() {
	ctx := context.Background()
	credentialByte, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentialByte, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := gsheetConf.GetClient(config)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "online")
	})

	e.GET("/sheet", func(c echo.Context) error {
		ReadFromSheet(srv)
		return nil
	})
	
	e.Logger.Fatal(e.Start(":8000"))
	
	// scanner := bufio.NewScanner(os.Stdin)
}

func WriteToSheet(scanner *bufio.Scanner, srv *sheets.Service) error {
	fmt.Println("input row:")
	scanner.Scan()
	row := scanner.Text()
	log.Println(row)

	fmt.Println("input name:")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("input email:")
	scanner.Scan()
	email := scanner.Text()

	writeRange := "A" + string(row[:]) + ":B"
	writeData := [][]interface{}{
		{
			name,
		},
		{
			email,
		},
	}

	writeValue := sheets.ValueRange{
		MajorDimension: "COLUMNS",
		Values: writeData,
	}

	writeReq := srv.
		Spreadsheets.
		Values.
		Update(
			gsheetConf.SHEET_ID,
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

func ReadFromSheet(srv *sheets.Service) error {
	// Read from sheet with id below:
	spreadsheetId := gsheetConf.SHEET_ID
	readRange := "A1:C"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			for _, v := range row {
				fmt.Printf("%s ", v)
			}
			fmt.Println()
		}
	}

	return nil
}