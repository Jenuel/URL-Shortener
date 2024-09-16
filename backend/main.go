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
	dsn := "root@tcp(127.0.0.1:3306)/shortener"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	e := echo.New() // Creates a new instance
	// Define routes
	e.POST("/shrink", shrinkUrl)
	e.GET("/links", getAllLinks)
	e.GET("/:id", redirectUrl)

	e.Logger.Fatal(e.Start(":1323"))
}

// Exported field names
type URLData struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original"`
	ShortCode   string `json:"short"`
	ClickCount  int    `json:"click"`
}

type URLRequest struct {
	URL string `json:"url"`
}

func shrinkUrl(c echo.Context) error {
	var req URLRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	code, err := generateCode(req.URL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating code"})
	}

	_, err = db.Exec("INSERT INTO links (original_url, short_code, click_count) VALUES (?, ?, ?)", req.URL, code, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving URL to database"})
	}

	return c.String(http.StatusOK, "Shrinking was successful")
}

func isKeyExisting(key string) (bool, error) {
	var shortCode string
	query := `SELECT short_code FROM links WHERE short_code = ?`
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
	var links []URLData

	query := "SELECT id, original_url, short_code, click_count FROM links"
	rows, err := db.Query(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching links")
	}
	defer rows.Close()

	for rows.Next() {
		var link URLData
		if err := rows.Scan(&link.ID, &link.OriginalURL, &link.ShortCode, &link.ClickCount); err != nil {
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
	code := c.Param("id")
	query := "SELECT original_url FROM links WHERE short_code = ?"

	var originalURL string

	err := db.QueryRow(query, code).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Short code not found.")
		}
		return c.String(http.StatusInternalServerError, "Server error occurred.")
	}

	return c.Redirect(http.StatusFound, originalURL)
}
