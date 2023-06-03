package main

import (
	"fmt"
	"os"
)

func main() {
	var file, showPosition, err = loadFile()

	if err != nil {
		fmt.Println("Não foi possível ENCONTRAR o arquivo!")
		os.Exit(1)
	}

	floor := 0
	buff := make([]byte, 1)
	pos := 0
	for no, err := file.Read(buff); err == nil; no, err = file.Read(buff) {
		for _, char := range string(buff[0:no]) {
			pos++
			switch char {
			case '(':
				floor++
			case ')':
				floor--
			}

			if len(showPosition) > 0 {
				if floor == -1 {
					fmt.Printf("POS: %d\n", pos)
					os.Exit(0)
				}
			}
		}
	}

	fmt.Printf("Floor %d", floor)
}

func loadFile() (*os.File, string, error) {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	var showPosition string
	if len(os.Args) > 2 {
		showPosition = os.Args[2]
	}

	var file, err = os.OpenFile(path, os.O_RDWR, os.ModeType)
	if err != nil {
		return nil, "", err
	}

	return file, showPosition, nil
}
