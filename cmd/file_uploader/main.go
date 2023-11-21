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
	retryCtrl := make(chan string, 10)

	go func() {
		for fileName := range retryCtrl {
			wg.Add(1)
			uploadCtrl <- struct{}{}
			go controller.UploadFile(uploadCtrl, retryCtrl, &wg, fileName)
		}
	}()

	for _, file := range files {
		wg.Add(1)
		uploadCtrl <- struct{}{}
		go controller.UploadFile(uploadCtrl, retryCtrl, &wg, file.Name())
	}
	wg.Wait()
}
