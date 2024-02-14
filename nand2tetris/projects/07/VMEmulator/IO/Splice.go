package IO

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func AppendFileContent(src, dst string) error {
	destFile, err := os.OpenFile(dst, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer destFile.Close()

	FirstFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer FirstFile.Close()

	_, err = io.Copy(destFile, FirstFile)
	if err != nil {
		return err
	}
	return nil
}

func Splice(path string) {
	fileName := filepath.Base(path)
	destName := path + "/" + fileName + ".asm"
	println(destName)

	err := AppendFileContent("other/Begin", destName)
	if err != nil {
		fmt.Println(err)
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".asm" {
			fullPath := filepath.Join(path, file.Name())
			if fullPath == destName {
				continue
			}
			err := AppendFileContent(fullPath, destName)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
