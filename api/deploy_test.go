package api_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/geprog/static-web/api"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("domain", "mypage")
	writer.WriteField("user", "admin123")
	part, _ := writer.CreateFormFile("archive", "sample.tar.gz")

	jsonFile, err := os.Open(path.Join("test", "sample.tar.gz"))
	defer jsonFile.Close()
	assert.NoError(t, err)

	tarReader := io.TeeReader(jsonFile, part)

	_, err = io.Copy(part, tarReader)
	assert.NoError(t, err)
	writer.Close() // <<< important part

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/deploy", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, api.Deploy(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
