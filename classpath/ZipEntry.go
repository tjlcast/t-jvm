package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath		string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath:absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(self.absPath)
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == className {
			rc, err := file.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
