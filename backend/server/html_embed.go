package server

import (
	"embed"
	_ "embed"
)

//go:embed html_files/*
var htmlFS embed.FS

func GetHtmlFile(filename string) ([]byte, error) {
	return htmlFS.ReadFile(filename)
}
