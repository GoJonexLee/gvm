package classpath

import (
	"io/ioutil"
	"errors"
	"path/filepath"
	"archive/zip"
)

type ZipEntry struct {
	absPath string
	zipRC *zip.ReadCloser
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}

func (ze *ZipEntry) openJar() error {
	r, err := zip.OpenReader(ze.absPath)
	if err == nil {
		ze.zipRC = r
	}
	return err
}

func (ze *ZipEntry) findClass(className string) *zip.File {
	for _, f := range ze.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if ze.zipRC == nil {
		if err := ze.openJar(); err != nil {
			return nil, nil, err
		}
	}

	classFile := ze.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, ze, err
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}