package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/geprog/static-web/cli/api"
	"github.com/urfave/cli/v2"
)

var TeardownCommand = &cli.Command{

	Name:   "teardown",
	Usage:  "teardown a page",
	Action: teardown,
}

func teardown(cCtx *cli.Context) error {
	form := url.Values{}
	form.Add("domain", cCtx.Args().First())

	req, err := http.NewRequestWithContext(cCtx.Context, http.MethodPost, api.GetUrl()+"/api/teardown", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+api.GetToken())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(body))

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error")
	}

	return nil
}
