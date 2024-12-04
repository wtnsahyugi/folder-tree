package main

import (
	"folder-tree/config"
	"folder-tree/internal/rest"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	err = rest.StartServer(*cfg)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
