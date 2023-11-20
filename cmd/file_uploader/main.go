package main

import (
	"os"

	"github.com/zHenriqueGN/GoUploader/internal/uploader"
)

func main() {
	files, err := os.ReadDir("./tmp")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		uploader.UploadFile(file.Name())
	}
}
