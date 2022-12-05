package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/geprog/static-web/lib"
	"github.com/labstack/echo/v4"
)

func Deploy(c echo.Context) error {
	// Read form fields
	domain := c.FormValue("domain")
	user := c.Request().Header.Get("Authorization")
	user = strings.Replace(user, "Bearer ", "", 1)

	log.Println("deploy: domain", domain)

	if domain == "" {
		domain = lib.GetRandomName(0)
	}

	if user == "" {
		return fmt.Errorf("upload: user is empty")
	}

	// Register (new) page and check if user is allowed to deploy
	page, err := RegisterPage(user, domain)
	if err != nil {
		return err
	}

	// Source
	file, err := c.FormFile("archive")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Extract
	folder := GetPageFolder(domain)
	if err := lib.Decompress(src, folder); err != nil {
		return err
	}

	page.LastUpdate = time.Now().Unix()
	err = WritePageMeta(page)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"page": page,
		"file": file.Filename,
		"ok":   true,
	})
}
