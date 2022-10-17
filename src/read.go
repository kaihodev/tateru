package tateru

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ReadConfig(loc string) *[]byte {
	targetPath := filepath.Dir(loc)
	fileName := filepath.Base(loc)
	if targetPath == "" { targetPath, _ = os.Getwd() }
	if fileName == "." { fileName = ".taterurc" }
	basePath := "/"
	for {
		targetPath, _ = filepath.Abs(targetPath)
		rel, _ := filepath.Rel(basePath, targetPath)
		if rel == "." { break }
		p := filepath.Join(targetPath, fileName)
		if !exists(p) {
			targetPath = filepath.Join(targetPath, "..")
			continue
		}
		bytes, err := ioutil.ReadFile(p)
		if err != nil { log.Fatalln(err) }
		return annihilateCRLFUnsafe(&bytes)
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

func annihilateCRLFUnsafe(bs *[]byte) *[]byte {
	b, i, offset := *bs, 0, 0
	for L := len(b); i != L; i++ {
		c := b[i]
		if c == '\r' {
			offset++
			continue
		}
		b[i - offset] = b[i]
	}
	*bs = b[:i - offset]
	return bs
}
