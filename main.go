package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	gsheetConf "github.com/thiccpan/sheetter/config"
	"github.com/thiccpan/sheetter/internal/handler"
	"github.com/thiccpan/sheetter/internal/usecase"
	"github.com/thiccpan/sheetter/internal/usecase/webapi"
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

	userUsecase := usecase.NewUserUsecase(*webapi.NewSheetApi(srv, gsheetConf.SHEET_ID))

	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "online")
	})

	handler.NewUserRoute(e, userUsecase)
	
	e.Logger.Fatal(e.Start(":8000"))
	
	// scanner := bufio.NewScanner(os.Stdin)
}