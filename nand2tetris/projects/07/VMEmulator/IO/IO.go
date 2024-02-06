package IO

import (
	"VMEmulator/Code"
	"VMEmulator/Parser"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var path string

func ReadFile() string {
	fmt.Print("Please enter the file path: ")
	fmt.Scanf("%s", &path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "None"
	}
	defer file.Close()

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
	return fileName[:len(fileName)-3]
}

func WriteFile() {
	path = path[:len(path)-3] + ".asm"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, ins := range Code.Ans {
		_, err := file.WriteString(ins + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
