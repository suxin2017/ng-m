package utils

import (
	"os"
)

func Test() {
	println("hello")
}

func getFileOrCreate(dst string) (*os.File, error) {

	return os.Create(dst)

}
