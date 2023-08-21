package model

// 存储文本信息
var TextData = make(map[string]string, 20)

// 写入后将对应的textName 在这个map中置为true
var CleanList = make(map[string]bool, 20)
