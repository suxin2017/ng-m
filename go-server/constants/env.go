package constants

import (
	"errors"
	"os"
	"path"
	"runtime"
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
	return path.Join(GetNginxConfigDir(), "demoain")
}

func InitEnvDir() {
	InitNotExistDir(GetLogDir())
	InitNotExistDir(GetNginxDomainConfigDir())
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
