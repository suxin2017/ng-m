package controller

import (
	"path"
	"testing"

	"suxin2017.com/server/constants"
)

func TestGetFileTreeFromDir(t *testing.T) {

	GetFileTreeFromDir(constants.GetNginxDomainResourceDir(9))
	GetFileTreeFromDir(path.Join(constants.GetNginxDomainResourceDir(9), "nginx"))
}
