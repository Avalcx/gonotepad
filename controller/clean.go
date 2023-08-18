package controller

import (
	"notepad/config"
	"os"
	"time"
)

func Clean(textName string, textType string, sleepTime int) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	switch textType {
	case "1":
		cleanFile(textName)
		return
	case "2":
		cleanMap(textName)
		return
	}
	return
}

func cleanFile(textName string) {
	os.Remove(config.FileDir + textName)
}

func cleanMap(textName string) {
	delete(textMap, textName)
}
