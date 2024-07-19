package main

import (
	"embed"
)

//go:embed public/*
var static embed.FS

func main() {

}
