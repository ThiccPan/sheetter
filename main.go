package main

import (
	"context"
	"fmt"
	"log"
	"os"

	googleOauthConf "github.com/thiccpan/sheetter/config"
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
	
	client := googleOauthConf.GetClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Read from sheet with id below:
	spreadsheetId := "18rlq5xu4YmxOmk3Xb-Nj6NbBygTeSKukq9EomNynvWI"
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

	writeRange := "A4:B"
	writeData := [][]interface{}{
		{
			"user3",
		},
		{
			"user3@gmail.com",
		},
	}

	writeValue := sheets.ValueRange{
		MajorDimension: "COLUMNS",
		Values: writeData,
	}

	writeReq := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &writeValue).ValueInputOption("RAW")
	writeRes, err := writeReq.Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(writeRes)
}
