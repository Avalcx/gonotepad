package controller

import (
	"errors"
	"notepad/config"
	"os"
)

func Write(textName string, text string, textType string) error {
	// readType 为1时 使用文件方式 为2时 使用内存方式
	switch textType {
	case "1":
		return writeFile(textName, text)
	case "2":
		return writeMap(textName, text)
	default:
		return errors.New("error")
	}
}

func writeFile(textName string, text string) error {
	err := os.WriteFile(config.FileDir+textName, []byte(text), 0666)
	if err != nil {
		return errors.New("write file error")
	} else {
		return nil
	}
}

func writeMap(textName string, text string) error {
	textMap[textName] = text
	return nil
}
