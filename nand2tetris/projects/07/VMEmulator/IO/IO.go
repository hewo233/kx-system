package IO

import (
	"VMEmulator/Code"
	"VMEmulator/Parser"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Init() {
	Parser.Init()
	Code.Init()
}

func DealFile(path string) {
	Init()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Print(line + "\n")
		Parser.Parse(line)
	}
	//fmt.Println("READ OVER")

	fileName := filepath.Base(path)
	useName := fileName[:len(fileName)-3]

	Code.Pass(useName)
	WriteFile(path)
}

func WriteFile(path string) {
	path = path[:len(path)-3] + ".asm"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	for _, ins := range Code.Ans {
		_, err := file.WriteString(ins + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ReadFile() {
	var path string
	fmt.Print("Please enter the file path: ")
	fmt.Scanf("%s", &path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".vm" {
				fullPath := filepath.Join(path, file.Name())
				DealFile(fullPath)
			}
		}
		Splice(path)
	} else {
		DealFile(path)
	}
}
