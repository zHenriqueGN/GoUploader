package main

import (
	"fmt"
	"os"

	"github.com/zHenriqueGN/GoUploader/internal/config"
)

func main() {
	numFiles := config.EnvVars.NumFiles
	for i := 0; i < numFiles; i++ {
		file, err := os.Create(fmt.Sprintf("./tmp/file-%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString("loren ipsum dolor sit amet set consetetur sadipscing at vero eos et accusam aliquyam erat")
	}
}
