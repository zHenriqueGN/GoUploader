package main

import (
	"os"
	"sync"

	"github.com/zHenriqueGN/GoUploader/internal/config"
	"github.com/zHenriqueGN/GoUploader/internal/controller"
)

func init() {
	config.LoadConfigs()
}

func main() {
	files, err := os.ReadDir("./tmp")
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	uploadCtrl := make(chan struct{}, 100)
	for _, file := range files {
		wg.Add(1)
		uploadCtrl <- struct{}{}
		go controller.UploadFile(uploadCtrl, &wg, file.Name())
	}
	wg.Wait()
}
