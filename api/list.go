package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/geprog/static-web/lib"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	user := c.Request().Header.Get("Authorization")
	user = strings.Replace(user, "Bearer ", "", 1)

	if user == "" {
		return fmt.Errorf("list: user is empty")
	}

	pages, err := GetPages(user)
	if err != nil {
		return err
	}

	data := struct {
		Ok    bool            `json:"ok"`
		Pages []*lib.PageMeta `json:"pages"`
	}{
		Ok:    true,
		Pages: pages,
	}

	return c.JSON(http.StatusOK, data)
}
