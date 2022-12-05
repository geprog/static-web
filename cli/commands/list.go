package commands

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/geprog/static-web/cli/api"
	"github.com/geprog/static-web/lib"
	"github.com/urfave/cli/v2"
)

var ListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list all pages",
	Action:  list,
}

func list(cCtx *cli.Context) error {
	req, err := http.NewRequestWithContext(cCtx.Context, http.MethodGet, api.GetUrl()+"/api/list", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+api.GetToken())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data struct {
		Ok    bool           `json:"ok"`
		Pages []lib.PageMeta `json:"pages"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	tmpl, err := template.New("_").Parse(tmplInfo)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, data)
}

var tmplInfo = `Pages:
{{range .Pages}}
---
domain: "{{.Domain}}"
owner:  {{.Owner}}
last update: {{.LastUpdate}}
{{end}}`
