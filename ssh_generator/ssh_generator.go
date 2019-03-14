package sshgenerator

import (
	"html/template"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Server data used to generate ssh config
type Server struct {
	Port         int
	User         string
	Host         string
	Hostname     string
	ProxyCommand string
	Identityfile string
}

const configTemplate = `
Host {{.Host}}
    {{with .Hostname -}}Hostname {{.}}{{end}}
    {{with .User -}}User {{.}}{{end}}
    {{with .Port -}}Port {{.}}{{end}}
    {{with .Identityfile -}}Identityfile {{.}}{{end}}
    {{with .ProxyCommand -}}ProxyCommand {{.}}{{end}}
`

// RenderTemplate generates an ssh config from server type slicey boi
func RenderTemplate(servers []Server) {
	for i := range servers {
		tmpl, err := template.New("config").Parse(configTemplate)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(os.Stdout, servers[i])
		if err != nil {
			panic(err)
		}
	}
}
