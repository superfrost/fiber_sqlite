package helpers

import "os"

func CreateWorkDirs() {
	os.MkdirAll("./public/img/mini/result", 0755)
	os.MkdirAll("./public/img/result", 0755)
}
