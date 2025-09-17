package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"notepad/config"
	"notepad/model"
	"notepad/router"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var Static embed.FS

func init() {
	config.SetConfig()
}

func main() {
	go model.StartCleaner()
	r := gin.Default()
	static, _ := fs.Sub(Static, "static")
	r.StaticFS("/static", http.FS(static))
	templ := template.Must(template.New("").ParseFS(Static, "static/*.html"))
	r.SetHTMLTemplate(templ)
	router.Jump(r)

	router.NotePad(r)
	router.CleanData(r)
	r.Run("0.0.0.0:" + config.Port)
}
