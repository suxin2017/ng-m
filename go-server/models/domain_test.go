package models_test

import (
	"log"
	"testing"

	. "suxin2017.com/server/models"
)

func TestGetNginxConfig(t *testing.T) {
	location := Location{
		Root:      "",
		Path:      "/",
		MatchRule: "=",
		Upstream: Upstream{
			Name: "upda",
			Servers: []Server{
				{
					ServerHost: "123.123.123.123",
					ServerPort: "778",
				},
				{
					ServerHost: "123.123.123.123",
					ServerPort: "778",
				},
			},
		},
	}

	log.Println(GetNginxConfigFromLocationAndDomain(location, Domain{
		Domain: "www.abc.com",
	}))

}
