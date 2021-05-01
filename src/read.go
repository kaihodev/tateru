package tateru

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func ReadConfig(loc string) *[]byte {
	targetPath := filepath.Dir(loc)
	fileName := filepath.Base(loc)
	if targetPath == "" { targetPath, _ = os.Getwd() }
	if fileName == "." { fileName = ".taterurc" }
	basePath := "/"
	for {
		rel, _ := filepath.Rel(basePath, targetPath)
		if rel == "." { break }
		p := path.Join(targetPath, fileName)
		if !exists(p) {
			targetPath = path.Join(targetPath, "..")
			continue
		}
		bytes, err := ioutil.ReadFile(p)
		if err != nil { log.Fatalln(err) }
		return &bytes
	}
	return nil
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
