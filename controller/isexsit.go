package controller

import (
	"notepad/config"
	"os"
)

func IsExsit(textName string, textType string) error {
	switch textType {
	case "1":
		return fileIsExsit(textName)
	case "2":
		return mapIsExsit(textName)
	}
	return nil
}

func fileIsExsit(textName string) error {
	filePath := config.FileDir + textName
	_, err := os.Stat(filePath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := Write(textName, "", config.TextType)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func mapIsExsit(textName string) error {
	return nil
}
