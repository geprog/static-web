package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Teardown(c echo.Context) error {
	user := c.Request().Header.Get("Authorization")
	user = strings.Replace(user, "Bearer ", "", 1)

	if user == "" {
		return fmt.Errorf("list: user is empty")
	}

	domain := c.FormValue("domain")

	err := TeardownPage(user, domain)
	if err != nil {
		return err
	}

	data := struct {
		Ok bool `json:"ok"`
	}{
		Ok: true,
	}

	return c.JSON(http.StatusOK, data)
}
