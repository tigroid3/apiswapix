package image

import (
	"fmt"
	"os"
	"path/filepath"
)

//дальнейшая реализация на сервере фото https://github.com/imgproxy/imgproxy

func StaticDir() string {
	return os.Getenv("PATH_PROJECT") + "static/"
}

func PreparePathByDirAndFilename(directoryName, filename string) string {
	subDir1 := filename[0:2]
	subDir2 := filename[2:4]

	dirPath := filepath.Join(StaticDir(), directoryName, subDir1, subDir2)
	err := os.MkdirAll(dirPath, 775)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	return filepath.Join(dirPath, filename)
}

func GetUrl(directoryName, fileName string) string {
	return ""
}
