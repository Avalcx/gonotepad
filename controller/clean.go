package controller

import (
	"fmt"
	"notepad/config"
	"notepad/model"
	"os"
)

func Clean(textName string) {
	switch config.TextType {
	case "1":
		cleanFile(textName)
	case "2":
		cleanMap(textName)
	}
}

func cleanFile(textName string) {
	filePath := config.FileDir + textName
	fmt.Println("remove file")
	err := os.Remove(filePath)
	if err == nil {
		return
	}
}

func cleanMap(textName string) {
	delete(model.TextData, textName)
}
