package controller

import (
	"errors"
	"notepad/config"
	"notepad/model"
	"os"
)

// var textMap = make(map[string]string, 50)

func Read(textName string, textType string) (string, error) {

	// readType 为1时 使用文件方式 为2时 使用内存方式
	switch textType {
	case "1":
		return readFile(textName)
	case "2":
		return readMap(textName)
	default:
		return "error", errors.New("error")
	}
}

func readFile(textName string) (string, error) {
	filePath := config.FileDir + textName
	text, err := os.ReadFile(filePath)
	return string(text), err
}

func readMap(textName string) (string, error) {
	text, _ := model.Get(textName)
	return text, nil
}
