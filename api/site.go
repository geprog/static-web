package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/geprog/static-web/lib"
)

const (
	pagesPath = "sites"
)

func GetPageFolder(domain string) string {
	return path.Join(pagesPath, domain)
}

func GetPageMeta(domain string) (*lib.PageMeta, error) {
	jsonFile, err := os.Open(path.Join(GetPageFolder(domain), ".meta", "config.json"))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var pageMetaData lib.PageMeta
	err = json.Unmarshal(byteValue, &pageMetaData)
	if err != nil {
		return nil, err
	}

	return &pageMetaData, nil
}

func WritePageMeta(pageMeta *lib.PageMeta) error {
	if err := os.MkdirAll(path.Join(GetPageFolder(pageMeta.Domain), ".meta"), 0755); err != nil {
		return fmt.Errorf("WritePageMeta: Mkdir() failed: %w", err)
	}

	jsonFile, err := os.Create(path.Join(GetPageFolder(pageMeta.Domain), ".meta", "config.json"))
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(pageMeta)
	if err != nil {
		return err
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	return nil
}

func GetPages(user string) ([]*lib.PageMeta, error) {
	var pages []*lib.PageMeta

	filepath.Walk(pagesPath, func(file string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			if pageMeta, err := GetPageMeta(fi.Name()); err == nil && pageMeta.Owner == user {
				pages = append(pages, pageMeta)
			}
		}
		return nil
	})

	return pages, nil
}

func RegisterPage(user string, domain string) (*lib.PageMeta, error) {
	pageMeta, err := GetPageMeta(domain)

	// other error
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("RegisterPage: GetPageMeta() failed: %w", err)
	}

	// page does not exist yet
	if err != nil && os.IsNotExist(err) {
		pageMeta = &lib.PageMeta{
			Owner:      user,
			Domain:     domain,
			LastUpdate: time.Now().Unix(),
		}
	}

	if pageMeta.Owner != user {
		return nil, fmt.Errorf("RegisterPage: page already exists and is owned by another user")
	}

	if err := WritePageMeta(pageMeta); err != nil {
		return nil, fmt.Errorf("RegisterPage: WritePageMeta() failed: %w", err)
	}

	return pageMeta, nil
}

func TeardownPage(user, domain string) error {
	pageMeta, err := GetPageMeta(domain)

	// other error
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("TeardownPage: GetPageMeta() failed: %w", err)
	}

	// page does not exist
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("TeardownPage: page does not exist")
	}

	if pageMeta.Owner != user {
		return fmt.Errorf("TeardownPage: page is owned by another user")
	}

	if err := os.RemoveAll(GetPageFolder(domain)); err != nil {
		return fmt.Errorf("TeardownPage: RemoveAll() failed: %w", err)
	}

	return nil
}
