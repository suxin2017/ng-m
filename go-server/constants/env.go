package constants

import (
	"errors"
	"os"
	"path"
	"runtime"
	"strconv"
)

func GetProjectRoot() string {
	_, cwd, _, _ := runtime.Caller(1)
	return path.Join(cwd, "../../../")
}
func GetRootDir() string {
	return path.Join(GetProjectRoot(), "go-server")
}

func GetServerDir() string {
	return GetRootDir()
}

func GetLogDir() string {
	return path.Join(GetServerDir(), "logs")
}

func GetScriptDir() string {
	return path.Join(GetProjectRoot(), "script")
}

func GetGinLogPath() string {
	return path.Join(GetLogDir(), "gin.log")
}

func GetNginxConfigDir() string {
	return path.Join(GetProjectRoot(), "ng-config")
}

func GetNginxDomainConfigDir() string {
	return path.Join(GetNginxConfigDir(), "domain")
}

func GetNginxUploadDir() string {
	return path.Join(GetNginxConfigDir(), "upload")
}

func GetNginxDomainResourceDir(domainId uint) string {
	return path.Join(GetNginxDomainConfigDir(), strconv.FormatUint(uint64(domainId), 10))
}

func InitEnvDir() {
	InitNotExistDir(GetLogDir())
	InitNotExistDir(GetNginxDomainConfigDir())
	InitNotExistDir(GetNginxUploadDir())
}

func InitNotExistDir(dirPath string) {
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		os.MkdirAll(dirPath, 0750)
	}
}

func IsDev() bool {
	return os.Getenv("ENV") == "development"
}
