package models

import (
	"fmt"
	"path"
	"strings"
	"text/template"

	"gorm.io/gorm"
	"suxin2017.com/server/constants"
)

type Domain struct {
	CommonModel
	Domain        string     `validate:"required" json:"domain,omitempty"`
	Desc          string     `json:"desc,omitempty"`
	CreatedUserID uint       `json:"createdUserId,omitempty"`
	CreatedUser   User       `json:"createdUser,omitempty"`
	UpdatedUserId uint       `json:"updatedUserId,omitempty"`
	UpdatedUser   User       `json:"updatedUser,omitempty"`
	Users         []User     `gorm:"many2many:user_domain;" json:"users,omitempty"`
	Locations     []Location `json:"locations,omitempty"`
}

type Location struct {
	CommonModel
	DomainId  uint   `json:"domainId,omitempty"`
	MatchRule string `json:"matchRule,omitempty"`
	Type      int    `json:"type,omitempty"`
	Root      string `json:"root,omitempty"`
	Path      string `json:"path,omitempty"`
	// Upstream   Upstream `json:"upstream,omitempty"`
	// UpstreamId uint     `json:"upstreamId,omitempty"`
	ServerHost string `json:"serverHost,omitempty"`
	ServerPort string `json:"serverPort,omitempty"`
}

func (l *Location) BeforeCreate(*gorm.DB) error {
	l.Type = 1
	return nil
}

type Upstream struct {
	CommonModel
	Name    string `json:"name,omitempty"`
	Servers []Server
}

type Server struct {
	ServerHost string `json:"serverHost,omitempty"`
	ServerPort string `json:"serverPort,omitempty"`
	UpstreamID uint
}

func GetNginxConfigFromLocationAndDomain(locations []Location, domain Domain) string {
	var sb strings.Builder
	sb.WriteString(`server {
`)

	sb.WriteString(fmt.Sprintf("	access_log  logs/nginx/%v.access.log  main;\n", domain.ID))
	sb.WriteString(fmt.Sprintf("	server_name %v;\n", domain.Domain))
	sb.WriteString("	listen 80;\n")

	locatonStr := `
	location {{.MatchRule}} {{.Path}} {
{{if eq .Type 1}}		root  {{.Root}};{{end}}
{{if eq .Type 2}}		proxy_pass http://{{.ServerHost}}:{{.ServerPort}};{{end}}
	}
`
	for _, location := range locations {

		if location.Root != "" {
			location.Root = path.Join(constants.GetNginxDomainResourceDir(domain.ID), location.Root)
		}
		if location.Root != "" || location.ServerHost != "" {

			t := template.Must(template.New("location").Parse(locatonStr))

			t.Execute(&sb, location)
		}

	}

	sb.WriteString(`
}`)
	return sb.String()
}
