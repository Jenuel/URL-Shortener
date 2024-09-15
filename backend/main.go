package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

func main() {
	dsn := "admin:admin@tcp(127.0.0.1:3306)/shortener"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	e := echo.New() //Creates a new Instance
	e.Logger.Fatal(e.Start(":1323"))

	//routes
	e.POST("/shrink", shrinkUrl)
	e.GET("/links", getAllLinks)
	e.GET("/:id", redirectUrl)
}

type urlData struct {
	id           int    `json:"id"`
	original_url string `json:"original"`
	short_code   string `json:"short"`
	click_count  int    `json:"click"`
}

func shrinkUrl(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getAllLinks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func redirectUrl(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
