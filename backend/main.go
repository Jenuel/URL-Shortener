package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"time"

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

type urlRequest struct {
	url string `json:"url"`
}

func shrinkUrl(c echo.Context) error {
	var req urlRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	code, err := generateCode(req.url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating code"})
	}

	_, err = db.Exec("INSERT INTO links (original_url, short_code, click_count) VALUES (?, ?, ?)", req.url, code, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving URL to database"})
	}

	return c.String(http.StatusOK, "Shrinking was successful")
}

func isKeyExisting(key string) (bool, error) {
	var shortCode string
	query := `SELECT short_code FROM links WHERE short_code = $1`
	err := db.QueryRow(query, key).Scan(&shortCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func generateCode(url string) (string, error) {
	length := 5
	for {
		hash := sha256.New()
		hash.Write([]byte(url + time.Now().String()))
		hashBytes := hash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)

		if length > len(hashString) {
			length = len(hashString)
		}
		shortKey := hashString[:length]

		exists, err := isKeyExisting(shortKey)
		if err != nil {
			return "", err
		}

		if !exists {
			return shortKey, nil
		}
	}
}

func getAllLinks(c echo.Context) error {
	var links []urlData

	query := "SELECT id, original_url, short_code, click_count FROM links"
	rows, err := db.Query(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching links")
	}
	defer rows.Close()

	for rows.Next() {
		var link urlData
		if err := rows.Scan(&link.id, &link.original_url, &link.short_code, &link.click_count); err != nil {
			return c.String(http.StatusInternalServerError, "Error scanning row")
		}
		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return c.String(http.StatusInternalServerError, "Error processing rows")
	}

	return c.JSON(http.StatusOK, links)
}

func redirectUrl(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
