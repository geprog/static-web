package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/urfave/cli/v2"

	"github.com/geprog/static-web/cli/api"
	"github.com/geprog/static-web/lib"
)

var DeployCommand = &cli.Command{
	Name:      "deploy",
	Usage:     "deploy a folder as static page",
	ArgsUsage: "[path/to/.woodpecker.yml]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "domain",
			Aliases: []string{"d"},
			Usage:   "domain to deploy to",
		},
	},
	Action: deploy,
}

func deploy(cCtx *cli.Context) error {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("domain", cCtx.String("domain"))
	part, _ := writer.CreateFormFile("archive", "sample.tar.gz")

	err := lib.Compress("./page", part)
	if err != nil {
		return err
	}

	writer.Close() // <<< important part

	req, err := http.NewRequestWithContext(cCtx.Context, http.MethodPost, api.GetUrl()+"/api/deploy", body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+api.GetToken())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Printf(string(bodyRes))

	return nil
}
