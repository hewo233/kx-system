package IO

import (
	"Assembler/Code"
	"Assembler/Parser"
	"bufio"
	"fmt"
	"os"
)

var path string

func ReadFile() {
	fmt.Print("Please enter the file path: ")
	fmt.Scanf("%s", &path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
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
}

func WriteFile() {
	path = path[:len(path)-4] + ".hack"
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
